// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package equal

import (
	"testing"
)

func TestEqual(t *testing.T) {
	for _, test := range []struct {
		x, y interface{}
		want bool
	}{
		{1, 1, true},
		{1, 2, false},                    // different values
		{1, 1.0, false},                  // different types
		{1.000000001, 1.000000001, true}, // different types
		{1.000000001, 1.000000002, true}, // different types
	} {
		if Equal(test.x, test.y) != test.want {
			t.Errorf("Equal(%v, %v) = %t",
				test.x, test.y, !test.want)
		}
	}
}
