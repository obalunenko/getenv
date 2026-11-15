package internal

import (
	"encoding"
	"net"
	"reflect"
	"strconv"
)

// valueParser converts the raw value to a single value of a type known at run time.
type valueParser func(raw string) (reflect.Value, error)

var (
	stringType          = reflect.TypeFor[string]()
	errorType           = reflect.TypeFor[error]()
	textUnmarshalerType = reflect.TypeFor[encoding.TextUnmarshaler]()
	ipType              = reflect.TypeFor[net.IP]()
	hardwareAddrType    = reflect.TypeFor[net.HardwareAddr]()
)

// kindParsers parse the types whose underlying type is a basic type, by its kind.
var kindParsers = map[reflect.Kind]valueParser{
	reflect.String:     valueOf(parseString),
	reflect.Bool:       valueOf(strconv.ParseBool),
	reflect.Int:        valueOf(parseInt[int]),
	reflect.Int8:       valueOf(parseInt[int8]),
	reflect.Int16:      valueOf(parseInt[int16]),
	reflect.Int32:      valueOf(parseInt[int32]),
	reflect.Int64:      valueOf(parseInt[int64]),
	reflect.Uint:       valueOf(parseUint[uint]),
	reflect.Uint8:      valueOf(parseUint[uint8]),
	reflect.Uint16:     valueOf(parseUint[uint16]),
	reflect.Uint32:     valueOf(parseUint[uint32]),
	reflect.Uint64:     valueOf(parseUint[uint64]),
	reflect.Uintptr:    valueOf(parseUint[uintptr]),
	reflect.Float32:    valueOf(parseFloat[float32]),
	reflect.Float64:    valueOf(parseFloat[float64]),
	reflect.Complex64:  valueOf(parseComplex[complex64]),
	reflect.Complex128: valueOf(parseComplex[complex128]),
}

// customParserOf returns the function that parses raw into the value that dst points to,
// or nil if the type is not supported.
// The type is either a custom type (see valueParserOf), or a slice of custom or EnvParsable types.
func customParserOf(dst any, p Parameters) func(raw string) error {
	rv := reflect.ValueOf(dst).Elem()

	if parse := valueParserOf(rv.Type(), p); parse != nil {
		return func(raw string) error {
			v, err := parse(raw)
			if err != nil {
				return err
			}

			rv.Set(v)

			return nil
		}
	}

	if rv.Kind() != reflect.Slice {
		return nil
	}

	parse := valueParserOf(rv.Type().Elem(), p)
	if parse == nil {
		return nil
	}

	return func(raw string) error {
		elems, err := parseSlice(raw, p.Separator, parse)
		if err != nil {
			return err
		}

		res := reflect.MakeSlice(rv.Type(), len(elems), len(elems))
		for i, e := range elems {
			res.Index(i).Set(e)
		}

		rv.Set(res)

		return nil
	}
}

// valueParserOf returns the parser of a single value of type t, or nil if t is not supported.
// The first one that fits is used:
//   - the parser of a scalar EnvParsable type;
//   - the ParseEnv method of t;
//   - the UnmarshalText method of *t or of t, if it is a pointer type;
//   - the parser of the underlying basic type of t.
func valueParserOf(t reflect.Type, p Parameters) valueParser {
	if parse := builtinValueParserOf(t, p); parse != nil {
		return parse
	}

	if parse := envParserOf(t); parse != nil {
		return parse
	}

	if parse := textParserOf(t); parse != nil {
		return parse
	}

	return kindParserOf(t)
}

// builtinValueParserOf returns the parser of t if t is a scalar EnvParsable type, or nil otherwise.
func builtinValueParserOf(t reflect.Type, p Parameters) valueParser {
	// The slice EnvParsable types are split by the separator, so they cannot be slice elements.
	// net.IP and net.HardwareAddr are the only scalar EnvParsable types of the slice kind.
	if t.Kind() == reflect.Slice && t != ipType && t != hardwareAddrType {
		return nil
	}

	if !isBuiltin(reflect.New(t).Interface()) {
		return nil
	}

	return func(raw string) (reflect.Value, error) {
		ptr := reflect.New(t)

		if err := parseBuiltin(ptr.Interface(), raw, p); err != nil {
			return reflect.Value{}, err
		}

		return ptr.Elem(), nil
	}
}

// envParserOf returns the ParseEnv method of the zero value of t,
// if t has the method ParseEnv(string) (t, error), or nil otherwise.
func envParserOf(t reflect.Type) valueParser {
	m, ok := t.MethodByName("ParseEnv")
	// The method type has the receiver as the first argument.
	if !ok || m.Type != reflect.FuncOf([]reflect.Type{t, stringType}, []reflect.Type{t, errorType}, false) {
		return nil
	}

	parseEnv := reflect.Zero(t).Method(m.Index)

	return func(raw string) (reflect.Value, error) {
		out := parseEnv.Call([]reflect.Value{reflect.ValueOf(raw)})

		// A nil error is a nil interface, so the assertion fails for it.
		if err, ok := out[1].Interface().(error); ok {
			return reflect.Value{}, err
		}

		return out[0], nil
	}
}

// textParserOf returns the parser of t if *t implements encoding.TextUnmarshaler,
// or t is a pointer type that implements it, or nil otherwise.
func textParserOf(t reflect.Type) valueParser {
	switch {
	case reflect.PointerTo(t).Implements(textUnmarshalerType):
		return func(raw string) (reflect.Value, error) {
			ptr := reflect.New(t)

			if err := unmarshalText(ptr, raw); err != nil {
				return reflect.Value{}, err
			}

			return ptr.Elem(), nil
		}
	case t.Kind() == reflect.Pointer && t.Implements(textUnmarshalerType):
		return func(raw string) (reflect.Value, error) {
			ptr := reflect.New(t.Elem())

			if err := unmarshalText(ptr, raw); err != nil {
				return reflect.Value{}, err
			}

			return ptr, nil
		}
	default:
		return nil
	}
}

// unmarshalText calls the UnmarshalText method of ptr, which implements encoding.TextUnmarshaler.
func unmarshalText(ptr reflect.Value, raw string) error {
	u, ok := ptr.Interface().(encoding.TextUnmarshaler)
	if !ok {
		// Unreachable while textParserOf checks that ptr implements encoding.TextUnmarshaler.
		panic("getenv: " + ptr.Type().String() + " does not implement encoding.TextUnmarshaler")
	}

	return u.UnmarshalText([]byte(raw))
}

// kindParserOf returns the parser of t if the underlying type of t is a basic type, or nil otherwise.
func kindParserOf(t reflect.Type) valueParser {
	parse, ok := kindParsers[t.Kind()]
	if !ok {
		return nil
	}

	return func(raw string) (reflect.Value, error) {
		v, err := parse(raw)
		if err != nil {
			return reflect.Value{}, err
		}

		return v.Convert(t), nil
	}
}

// valueOf adapts parse to return a reflect.Value.
func valueOf[E any](parse func(string) (E, error)) valueParser {
	return func(raw string) (reflect.Value, error) {
		v, err := parse(raw)
		if err != nil {
			return reflect.Value{}, err
		}

		return reflect.ValueOf(v), nil
	}
}
