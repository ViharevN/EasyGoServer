package main

import "testing"

func TestAbs(t *testing.T) {
	tests := []struct {
		name string
		val  float64
		want float64
	}{{
		name: "test #1",
		val:  3.1,
		want: 3.1,
	}, {
		name: "test #2",
		val:  -3.1,
		want: 3.1,
	}, {
		name: "test #3",
		val:  -3.2,
		want: 3.2,
	},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if abs := Abs(test.val); abs != test.want {
				t.Errorf("Abs() = %f, want %f", abs, test.want)
			}
		})
	}
}
