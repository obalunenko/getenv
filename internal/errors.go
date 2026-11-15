package internal

import (
	"errors"
)

var (
	// ErrNotSet is an error that is returned when the environment variable is not set.
	ErrNotSet = errors.New("not set")
	// ErrInvalidValue is an error that is returned when the environment variable is not valid.
	ErrInvalidValue = errors.New("invalid value")
	// ErrUnsupportedType is an error that is wrapped by the panic value for a type that cannot be parsed.
	ErrUnsupportedType = errors.New("unsupported type")

	errEmptySeparator = errors.New("empty separator")
	errNotBuiltin     = errors.New("not an EnvParsable type")
)
