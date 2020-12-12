package math

import "testing"

type testpair struct {
	values []float64
	max    float64
}

var tests = []testpair{
	{[]float64{1, 2}, 2},
	{[]float64{1, 1, 1, 1, 1, 1, 1}, 1},
	{[]float64{-1, 1}, 1},
	{[]float64{}, 0},
}

func TestMax(t *testing.T) {
	for _, pair := range tests {
		v := Max(pair.values)
		if v != pair.max {
			t.Error(
				"For", pair.values,
				"expected", pair.max,
				"got", v,
			)
		}
	}
}
