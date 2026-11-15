package getenv

import (
	"fmt"
	"net"
	"net/netip"
	"net/url"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// comma holds the parameters for slice cases.
var comma = params{separator: ","}

type parseCase[T Parsable] struct {
	name    string
	raw     string
	params  params
	want    T
	wantErr bool
}

// testParse runs parseAs[T] for every case: an error must wrap ErrInvalidValue and come with the zero value.
func testParse[T Parsable](t *testing.T, cases ...parseCase[T]) {
	t.Helper()

	t.Run(fmt.Sprintf("%T", *new(T)), func(t *testing.T) {
		for _, tc := range cases {
			t.Run(tc.name, func(t *testing.T) {
				got, err := parseAs[T](tc.raw, tc.params)
				if tc.wantErr {
					require.ErrorIs(t, err, ErrInvalidValue)
					assert.Zero(t, got)

					return
				}

				require.NoError(t, err)
				assert.Equal(t, tc.want, got)
			})
		}
	})
}

func TestParseString(t *testing.T) {
	testParse(t,
		parseCase[string]{name: "value", raw: "newval", want: "newval"},
	)
	testParse(t,
		parseCase[[]string]{name: "values", raw: "true,newval", params: comma, want: []string{"true", "newval"}},
		parseCase[[]string]{name: "other separator", raw: "a,b", params: params{separator: " "}, want: []string{"a,b"}},
		parseCase[[]string]{name: "elements are not trimmed", raw: " a, b", params: comma, want: []string{" a", " b"}},
		parseCase[[]string]{name: "trailing separator adds an empty element", raw: "a,b,", params: comma, want: []string{"a", "b", ""}},
	)
}

func TestParseInt(t *testing.T) {
	testParse(t,
		parseCase[int]{name: "value", raw: "-128", want: -128},
		parseCase[int]{name: "invalid", raw: "128s7", wantErr: true},
	)
	testParse(t,
		parseCase[int8]{name: "max", raw: "127", want: 127},
		parseCase[int8]{name: "overflow", raw: "128", wantErr: true},
	)
	testParse(t,
		parseCase[int16]{name: "max", raw: "32767", want: 32767},
		parseCase[int16]{name: "overflow", raw: "32768", wantErr: true},
	)
	testParse(t,
		parseCase[int32]{name: "max", raw: "2147483647", want: 2147483647},
		parseCase[int32]{name: "overflow", raw: "2147483648", wantErr: true},
	)
	testParse(t,
		parseCase[int64]{name: "max", raw: "9223372036854775807", want: 9223372036854775807},
		parseCase[int64]{name: "overflow", raw: "9223372036854775808", wantErr: true},
	)
	testParse(t,
		parseCase[[]int]{name: "values", raw: "1,-2", params: comma, want: []int{1, -2}},
		parseCase[[]int]{name: "invalid element", raw: "1,asd", params: comma, wantErr: true},
		parseCase[[]int]{name: "space after separator", raw: "1, 2", params: comma, wantErr: true},
		parseCase[[]int]{name: "trailing separator", raw: "1,2,", params: comma, wantErr: true},
	)
	testParse(t, parseCase[[]int8]{name: "values", raw: "1,2", params: comma, want: []int8{1, 2}})
	testParse(t, parseCase[[]int16]{name: "values", raw: "1,2", params: comma, want: []int16{1, 2}})
	testParse(t, parseCase[[]int32]{name: "values", raw: "1,2", params: comma, want: []int32{1, 2}})
	testParse(t, parseCase[[]int64]{name: "values", raw: "1,2", params: comma, want: []int64{1, 2}})
}

func TestParseUint(t *testing.T) {
	testParse(t,
		parseCase[uint]{name: "value", raw: "1024", want: 1024},
		parseCase[uint]{name: "negative", raw: "-1", wantErr: true},
	)
	testParse(t,
		parseCase[uint8]{name: "max", raw: "255", want: 255},
		parseCase[uint8]{name: "overflow", raw: "256", wantErr: true},
	)
	testParse(t,
		parseCase[uint16]{name: "max", raw: "65535", want: 65535},
		parseCase[uint16]{name: "overflow", raw: "65536", wantErr: true},
	)
	testParse(t,
		parseCase[uint32]{name: "max", raw: "4294967295", want: 4294967295},
		parseCase[uint32]{name: "overflow", raw: "4294967296", wantErr: true},
	)
	testParse(t,
		parseCase[uint64]{name: "max", raw: "18446744073709551615", want: 18446744073709551615},
		parseCase[uint64]{name: "overflow", raw: "18446744073709551616", wantErr: true},
	)
	testParse(t,
		parseCase[uintptr]{name: "value", raw: "123", want: 123},
		parseCase[uintptr]{name: "invalid", raw: "iii99", wantErr: true},
	)
	testParse(t, parseCase[[]uint]{name: "values", raw: "1,2", params: comma, want: []uint{1, 2}})
	testParse(t, parseCase[[]uint8]{name: "values", raw: "1,2", params: comma, want: []uint8{1, 2}})
	testParse(t, parseCase[[]uint16]{name: "values", raw: "1,2", params: comma, want: []uint16{1, 2}})
	testParse(t, parseCase[[]uint32]{name: "values", raw: "1,2", params: comma, want: []uint32{1, 2}})
	testParse(t, parseCase[[]uint64]{name: "values", raw: "1,2", params: comma, want: []uint64{1, 2}})
	testParse(t, parseCase[[]uintptr]{name: "values", raw: "123,456", params: comma, want: []uintptr{123, 456}})
}

func TestParseFloat(t *testing.T) {
	testParse(t,
		parseCase[float32]{name: "value", raw: "1024.123", want: 1024.123},
		parseCase[float32]{name: "overflow", raw: "3.5e38", wantErr: true},
	)
	testParse(t,
		parseCase[float64]{name: "value", raw: "1024.123", want: 1024.123},
		parseCase[float64]{name: "invalid", raw: "1024.123s", wantErr: true},
	)
	testParse(t, parseCase[[]float32]{name: "values", raw: "1.05,2.07", params: comma, want: []float32{1.05, 2.07}})
	testParse(t, parseCase[[]float64]{name: "values", raw: "1.05,2.07", params: comma, want: []float64{1.05, 2.07}})
}

func TestParseComplex(t *testing.T) {
	testParse(t,
		parseCase[complex64]{name: "value", raw: "(1+2i)", want: 1 + 2i},
		parseCase[complex64]{name: "overflow", raw: "3.5e38+1i", wantErr: true},
	)
	testParse(t,
		parseCase[complex128]{name: "value", raw: "1+2i", want: 1 + 2i},
		parseCase[complex128]{name: "invalid", raw: "asd", wantErr: true},
	)
	testParse(t, parseCase[[]complex64]{name: "values", raw: "(1+2i),(3+4i)", params: comma, want: []complex64{1 + 2i, 3 + 4i}})
	testParse(t, parseCase[[]complex128]{name: "values", raw: "(1+2i),(3+4i)", params: comma, want: []complex128{1 + 2i, 3 + 4i}})
}

func TestParseBool(t *testing.T) {
	testParse(t,
		parseCase[bool]{name: "value", raw: "true", want: true},
		parseCase[bool]{name: "invalid", raw: "yes", wantErr: true},
	)
	testParse(t,
		parseCase[[]bool]{name: "values", raw: "true,false", params: comma, want: []bool{true, false}},
		parseCase[[]bool]{name: "invalid element", raw: "true,asd", params: comma, wantErr: true},
	)
}

func TestParseTime(t *testing.T) {
	const layout = "2006/02/01 15:04"

	p := params{separator: ",", layout: layout}
	want := time.Date(2018, 4, 21, 22, 30, 0, 0, time.UTC)

	testParse(t,
		parseCase[time.Time]{name: "value", raw: "2018/21/04 22:30", params: p, want: want},
		parseCase[time.Time]{name: "layout mismatch", raw: "2018-04-21", params: p, wantErr: true},
	)
	testParse(t,
		parseCase[[]time.Time]{name: "values", raw: "2018/21/04 22:30,2018/21/04 22:30", params: p, want: []time.Time{want, want}},
		parseCase[[]time.Time]{name: "invalid element", raw: "2018/21/04 22:30,20222-sdslll", params: p, wantErr: true},
	)
	testParse(t,
		parseCase[time.Duration]{name: "value", raw: "12s", want: 12 * time.Second},
		parseCase[time.Duration]{name: "invalid", raw: "12", wantErr: true},
	)
	testParse(t,
		parseCase[[]time.Duration]{name: "values", raw: "2m,3h", params: comma, want: []time.Duration{2 * time.Minute, 3 * time.Hour}},
		parseCase[[]time.Duration]{name: "invalid element", raw: "2m,3hddd", params: comma, wantErr: true},
	)
}

func TestParseURL(t *testing.T) {
	const raw = "postgres://user:pass@host.com:5432/path?k=v#f"

	want := mustParseURL(t, raw)

	testParse(t,
		parseCase[url.URL]{name: "value", raw: raw, want: want},
		parseCase[url.URL]{name: "invalid escape", raw: raw + "%%2", wantErr: true},
	)
	testParse(t,
		parseCase[[]url.URL]{
			name:   "values",
			raw:    "https://google.com,https://github.com",
			params: comma,
			want:   []url.URL{mustParseURL(t, "https://google.com"), mustParseURL(t, "https://github.com")},
		},
		parseCase[[]url.URL]{name: "invalid element", raw: "https://google.com,htps://%%2github.com", params: comma, wantErr: true},
	)
}

func TestParseURLErrorHidesValue(t *testing.T) {
	for _, raw := range []string{
		"postgres://admin:S3cr3t@db:5432/app%zz",
		"postgres://admin:S3cr3t@db:54x32/app",
	} {
		_, err := parseAs[url.URL](raw, params{})
		require.ErrorIs(t, err, ErrInvalidValue)
		assert.NotContains(t, err.Error(), "S3cr3t")
	}
}

func TestParseNetwork(t *testing.T) {
	testParse(t,
		parseCase[net.IP]{name: "ipv4", raw: "192.168.8.0", want: net.ParseIP("192.168.8.0")},
		parseCase[net.IP]{name: "ipv6", raw: "2001:db8::68", want: net.ParseIP("2001:db8::68")},
		parseCase[net.IP]{name: "invalid", raw: "192.168.8.0ssss", wantErr: true},
	)
	testParse(t,
		parseCase[[]net.IP]{
			name:   "values",
			raw:    "192.168.8.0,2001:cb8::17",
			params: comma,
			want:   []net.IP{net.ParseIP("192.168.8.0"), net.ParseIP("2001:cb8::17")},
		},
		parseCase[[]net.IP]{name: "invalid element", raw: "192.168.8.0,sdsdsd", params: comma, wantErr: true},
	)
	testParse(t,
		parseCase[netip.Addr]{name: "value", raw: "192.168.8.1", want: netip.MustParseAddr("192.168.8.1")},
		parseCase[netip.Addr]{name: "prefix is not an address", raw: "192.168.8.1/24", wantErr: true},
	)
	testParse(t,
		parseCase[[]netip.Addr]{
			name:   "values",
			raw:    "192.168.8.1,2001:db8::68",
			params: comma,
			want:   []netip.Addr{netip.MustParseAddr("192.168.8.1"), netip.MustParseAddr("2001:db8::68")},
		},
		parseCase[[]netip.Addr]{name: "invalid element", raw: "192.168.8.1,sdsdsd", params: comma, wantErr: true},
	)
	testParse(t,
		parseCase[netip.Prefix]{name: "value", raw: "192.168.8.0/24", want: netip.MustParsePrefix("192.168.8.0/24")},
		parseCase[netip.Prefix]{name: "address is not a prefix", raw: "192.168.8.0", wantErr: true},
	)
	testParse(t,
		parseCase[[]netip.Prefix]{
			name:   "values",
			raw:    "192.168.8.0/24,2001:db8::/64",
			params: comma,
			want:   []netip.Prefix{netip.MustParsePrefix("192.168.8.0/24"), netip.MustParsePrefix("2001:db8::/64")},
		},
		parseCase[[]netip.Prefix]{name: "invalid element", raw: "192.168.8.0/24,sdsdsd", params: comma, wantErr: true},
	)
	testParse(t,
		parseCase[net.HardwareAddr]{name: "value", raw: "01:23:45:67:89:ab", want: mustParseMAC(t, "01:23:45:67:89:ab")},
		parseCase[net.HardwareAddr]{name: "invalid", raw: "not-a-mac", wantErr: true},
	)
	testParse(t,
		parseCase[[]net.HardwareAddr]{
			name:   "values",
			raw:    "01:23:45:67:89:ab,02:23:45:67:89:ab",
			params: comma,
			want:   []net.HardwareAddr{mustParseMAC(t, "01:23:45:67:89:ab"), mustParseMAC(t, "02:23:45:67:89:ab")},
		},
		parseCase[[]net.HardwareAddr]{name: "invalid element", raw: "01:23:45:67:89:ab,not-a-mac", params: comma, wantErr: true},
	)
}

func TestParseSliceErrors(t *testing.T) {
	_, err := parseAs[[]string]("a,b", params{})
	require.ErrorIs(t, err, ErrInvalidValue)
	require.ErrorIs(t, err, errEmptySeparator)

	_, err = parseAs[[]int]("1,2,x", comma)
	require.ErrorIs(t, err, ErrInvalidValue)
	require.ErrorIs(t, err, strconv.ErrSyntax)
	assert.ErrorContains(t, err, "element 2")
}

// unsupported is a type that has no parser.
type unsupported struct{}

// The panics guard against programming errors in this package: parseAs cannot reach them
// while every Parsable type has a parser of its own type.
func TestParsePanics(t *testing.T) {
	t.Run("no parser for the type", func(t *testing.T) {
		var dst unsupported

		assert.PanicsWithValue(t, "getenv: no parser for *getenv.unsupported", func() {
			err := parseValue(&dst, "1", comma)
			t.Errorf("returned %v instead of panicking", err)
		})
	})

	t.Run("parser does not match the destination", func(t *testing.T) {
		var dst int

		assert.PanicsWithValue(t, "getenv: parser does not match the destination type", func() {
			err := parseInto(&dst, "1", comma, parseString)
			t.Errorf("returned %v instead of panicking", err)
		})
	})
}

func FuzzParse(f *testing.F) {
	seeds := []string{
		"", "1", "-1", "1,2", "1.5", "true", "2m", "(1+2i)", "2018-04-21T22:30:00Z",
		"192.168.8.1", "192.168.8.0/24", "01:23:45:67:89:ab", "https://google.com",
	}

	for _, s := range seeds {
		f.Add(s)
	}

	f.Fuzz(func(t *testing.T, raw string) {
		p := params{separator: ",", layout: time.RFC3339}

		checkErrContract[int](t, raw, p)
		checkErrContract[[]uint8](t, raw, p)
		checkErrContract[float32](t, raw, p)
		checkErrContract[complex64](t, raw, p)
		checkErrContract[[]bool](t, raw, p)
		checkErrContract[time.Time](t, raw, p)
		checkErrContract[[]time.Duration](t, raw, p)
		checkErrContract[url.URL](t, raw, p)
		checkErrContract[net.IP](t, raw, p)
		checkErrContract[[]netip.Prefix](t, raw, p)
		checkErrContract[net.HardwareAddr](t, raw, p)

		got, err := parseAs[[]string](raw, p)
		require.NoError(t, err)
		assert.Len(t, got, strings.Count(raw, ",")+1)
	})
}

// checkErrContract asserts that a failed parseAs returns the zero value and an error wrapping ErrInvalidValue.
func checkErrContract[T Parsable](t *testing.T, raw string, p params) {
	t.Helper()

	got, err := parseAs[T](raw, p)
	if err != nil {
		require.ErrorIs(t, err, ErrInvalidValue)
		require.Zero(t, got)
	}
}

func BenchmarkParse(b *testing.B) {
	b.ReportAllocs()

	for b.Loop() {
		if _, err := parseAs[[]float64]("1235.67,87.98", comma); err != nil {
			b.Fatal(err)
		}
	}
}

func mustParseURL(tb testing.TB, raw string) url.URL {
	tb.Helper()

	u, err := url.Parse(raw)
	require.NoError(tb, err)

	return *u
}

func mustParseMAC(tb testing.TB, raw string) net.HardwareAddr {
	tb.Helper()

	mac, err := net.ParseMAC(raw)
	require.NoError(tb, err)

	return mac
}
