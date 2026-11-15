// Package option provides options for parsing environment variables.
//
// Deprecated: the options moved to package getenv. Use getenv.Option,
// getenv.WithSeparator and getenv.WithTimeLayout instead.
package option

import (
	"github.com/obalunenko/getenv"
)

// Option configures how Env and EnvOrDefault parse the value of a variable.
//
// Deprecated: Use getenv.Option instead.
type Option = getenv.Option

// WithSeparator sets the separator of slice values. The default is ",".
//
// Deprecated: Use getenv.WithSeparator instead.
func WithSeparator(separator string) Option {
	return getenv.WithSeparator(separator)
}

// WithTimeLayout sets the layout of time.Time values. The default is time.RFC3339.
//
// Deprecated: Use getenv.WithTimeLayout instead.
func WithTimeLayout(layout string) Option {
	return getenv.WithTimeLayout(layout)
}
