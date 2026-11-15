package internal

import (
	"errors"
	"fmt"
	"log/slog"
	"net"
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var errUnknownLevel = errors.New("unknown level")

// level has the ParseEnv method.
type level int

func (level) ParseEnv(raw string) (level, error) {
	switch raw {
	case "debug":
		return 1, nil
	case "warn":
		return 2, nil
	default:
		return 0, fmt.Errorf("%w %q", errUnknownLevel, raw)
	}
}

// point implements encoding.TextUnmarshaler with a pointer receiver.
type point struct {
	x, y int
}

func (p *point) UnmarshalText(text []byte) error {
	_, err := fmt.Sscanf(string(text), "%d:%d", &p.x, &p.y)

	return err
}

// both has the ParseEnv method and implements encoding.TextUnmarshaler.
type both string

func (both) ParseEnv(raw string) (both, error) {
	return both("env:" + raw), nil
}

func (b *both) UnmarshalText(text []byte) error {
	*b = both("text:" + string(text))

	return nil
}

// wrongParseEnv has the ParseEnv method with another signature.
type wrongParseEnv struct{}

func (wrongParseEnv) ParseEnv(raw string) (string, error) {
	return raw, nil
}

// unsupported is a type that has no parser.
type unsupported struct{}

// Types whose underlying type is a basic type, and named slice types.
type (
	port     uint16
	mode     string
	offset   int8
	ratio    float64
	enabled  bool
	wave     complex64
	levels   []level
	ports    []port
	tags     []string
	timeouts []time.Duration
	stamps   []time.Time
	ips      []net.IP
)

func TestParseEnvParser(t *testing.T) {
	testParse(t,
		parseCase[level]{name: "value", raw: "warn", want: 2},
		parseCase[level]{name: "invalid", raw: "loud", wantErr: true},
	)
	testParse(t,
		parseCase[[]level]{name: "values", raw: "debug,warn", params: comma, want: []level{1, 2}},
		parseCase[[]level]{name: "invalid element", raw: "debug,loud", params: comma, wantErr: true},
	)
	testParse(t, parseCase[levels]{name: "named slice", raw: "debug,warn", params: comma, want: levels{1, 2}})

	_, err := Parse[[]level]("debug,loud", comma)
	require.ErrorIs(t, err, errUnknownLevel)
	assert.ErrorContains(t, err, "element 1")
}

func TestParseTextUnmarshaler(t *testing.T) {
	testParse(t,
		parseCase[point]{name: "value", raw: "1:2", want: point{x: 1, y: 2}},
		parseCase[point]{name: "invalid", raw: "1-2", wantErr: true},
	)
	testParse(t,
		parseCase[*point]{name: "pointer type", raw: "1:2", want: &point{x: 1, y: 2}},
		parseCase[*point]{name: "pointer type invalid", raw: "1-2", wantErr: true},
	)
	testParse(t, parseCase[[]point]{name: "values", raw: "1:2,3:4", params: comma, want: []point{{1, 2}, {3, 4}}})
	testParse(t, parseCase[[]*point]{name: "pointers", raw: "1:2,3:4", params: comma, want: []*point{{1, 2}, {3, 4}}})
	testParse(t,
		parseCase[slog.Level]{name: "standard library type", raw: "WARN", want: slog.LevelWarn},
		parseCase[slog.Level]{name: "standard library type invalid", raw: "LOUD", wantErr: true},
	)
}

func TestParseUnderlyingType(t *testing.T) {
	testParse(t,
		parseCase[port]{name: "value", raw: "8080", want: 8080},
		parseCase[port]{name: "overflow", raw: "65536", wantErr: true},
	)
	testParse(t, parseCase[mode]{name: "value", raw: "prod", want: "prod"})
	testParse(t, parseCase[offset]{name: "value", raw: "-128", want: -128})
	testParse(t, parseCase[ratio]{name: "value", raw: "0.5", want: 0.5})
	testParse(t, parseCase[enabled]{name: "value", raw: "true", want: true})
	testParse(t, parseCase[wave]{name: "value", raw: "1+2i", want: 1 + 2i})
	testParse(t,
		parseCase[[]port]{name: "values", raw: "80,443", params: comma, want: []port{80, 443}},
		parseCase[[]port]{name: "invalid element", raw: "80,https", params: comma, wantErr: true},
	)
	testParse(t, parseCase[ports]{name: "named slice", raw: "80,443", params: comma, want: ports{80, 443}})
	testParse(t, parseCase[tags]{name: "named slice of a basic type", raw: "a,b", params: comma, want: tags{"a", "b"}})
}

// The elements of a named slice type are parsed as their EnvParsable type, with the parameters.
func TestParseNamedSliceOfEnvParsable(t *testing.T) {
	testParse(t,
		parseCase[timeouts]{name: "durations", raw: "1s,2m", params: comma, want: timeouts{time.Second, 2 * time.Minute}},
		parseCase[timeouts]{name: "not parsed as int64", raw: "1,2", params: comma, wantErr: true},
	)
	testParse(t, parseCase[stamps]{
		name:   "time layout",
		raw:    "2022-01-20",
		params: Parameters{Separator: ",", Layout: time.DateOnly},
		want:   stamps{time.Date(2022, 1, 20, 0, 0, 0, 0, time.UTC)},
	})
	testParse(t, parseCase[ips]{
		name:   "slice kind element",
		raw:    "192.168.8.1,::1",
		params: comma,
		want:   ips{net.ParseIP("192.168.8.1"), net.ParseIP("::1")},
	})
}

func TestParsePrecedence(t *testing.T) {
	testParse(t, parseCase[both]{name: "ParseEnv over UnmarshalText", raw: "x", want: "env:x"})
	testParse(t, parseCase[slog.Level]{name: "UnmarshalText over the underlying type", raw: "WARN", want: slog.LevelWarn})
}

func TestParseUnsupportedType(t *testing.T) {
	testUnsupported[unsupported](t)
	testUnsupported[chan int](t)
	testUnsupported[*int](t)
	testUnsupported[map[string]string](t)
	testUnsupported[any](t)
	testUnsupported[[][]string](t)
	testUnsupported[[]unsupported](t)
	testUnsupported[wrongParseEnv](t)
}

// testUnsupported checks that MustSupport and Parse panic for T with an error that wraps ErrUnsupportedType.
func testUnsupported[T any](t *testing.T) {
	t.Helper()

	typ := reflect.TypeFor[T]().String()

	t.Run(typ, func(t *testing.T) {
		msg := "getenv: unsupported type: " + typ

		assert.PanicsWithError(t, msg, MustSupport[T])
		assert.PanicsWithError(t, msg, func() {
			_, err := Parse[T]("1", comma)
			t.Errorf("returned %v instead of panicking", err)
		})

		require.ErrorIs(t, recoverError(MustSupport[T]), ErrUnsupportedType)
	})
}

// recoverError returns the error that f panics with.
func recoverError(f func()) (err error) {
	defer func() {
		if e, ok := recover().(error); ok {
			err = e
		}
	}()

	f()

	return nil
}

// This panic guards against programming errors in this package: textParserOf cannot reach it
// while it checks that the type implements encoding.TextUnmarshaler.
func TestUnmarshalTextPanics(t *testing.T) {
	assert.PanicsWithValue(t, "getenv: *int does not implement encoding.TextUnmarshaler", func() {
		err := unmarshalText(reflect.ValueOf(new(int)), "1")
		t.Errorf("returned %v instead of panicking", err)
	})
}

func TestMustSupport(t *testing.T) {
	assert.NotPanics(t, MustSupport[int])
	assert.NotPanics(t, MustSupport[[]time.Time])
	assert.NotPanics(t, MustSupport[level])
	assert.NotPanics(t, MustSupport[[]*point])
	assert.NotPanics(t, MustSupport[ports])
}
