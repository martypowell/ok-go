package main

import (
	"fmt"
	"testing"
)

func TestGetBasicArithmatic(t *testing.T) {
	var tests = []struct {
		a, b, add, sub, multiply int
		divide                   float64
	}{
		{1, 1, 2, 0, 1, 1},
	}

	for _, test := range tests {
		testname := fmt.Sprintf("%d,%d", test.a, test.b)
		t.Run(testname, func(t *testing.T) {
			ansAdd, ansSub, ansMultiply, _ := getBasicArithmetic(test.a, test.b)

			if ansAdd != test.add {
				t.Errorf("got %d, want %d", ansAdd, test.add)
			}
			if ansSub != test.sub {
				t.Errorf("got %d, want %d", ansSub, test.sub)
			}
			if ansMultiply != test.multiply {
				t.Errorf("got %d, want %d", ansMultiply, test.multiply)
			}
			// TODO: Need to figure out how to compare float64 in test
			// if ansDivide != test.divide {
			// 	t.Errorf("got %d, want %d", ansDivide, test.divide)
			// }
		})
	}

}
