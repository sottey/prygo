package main

import (
	"reflect"
	"testing"
)

func TestNormalizePackages(t *testing.T) {
	t.Parallel()

	cases := []struct {
		in   string
		want []string
	}{
		{
			"",
			nil,
		},
		{
			"foo, bar,github.com/sottey/prygo ",
			[]string{
				"bar",
				"foo",
				"github.com/sottey/prygo",
			},
		},
	}

	for i, c := range cases {
		out := normalizePackages(c.in)
		if !reflect.DeepEqual(out, c.want) {
			t.Errorf("%d. normalizePackages(%q) = %+v; not %+v", i, c.in, out, c.want)
		}
	}
}
