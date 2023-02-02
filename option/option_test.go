package option

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/obalunenko/getenv/internal"
)

func TestOptions(t *testing.T) {
	var p internal.Parameters

	var opt Option

	opt = WithSeparator("|")

	opt.Apply(&p)

	expected := internal.Parameters{
		Separator: "|",
		Layout:    "",
	}

	assert.Equal(t, expected, p)

	opt = WithTimeLayout(time.RFC822)

	opt.Apply(&p)

	expected = internal.Parameters{
		Separator: "|",
		Layout:    time.RFC822,
	}

	assert.Equal(t, expected, p)
}
