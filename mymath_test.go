package mymath

import "testing"

// Round -- Returns value rounded to nearest whole number.
func TestRound(t *testing.T) {
	cases := []struct {
		in, want float64
	}{
		{100.5, 101},
		{100.499999999, 100},
		{-0.77777, -1},
		{-0.11111, 0},
		{0, 0},
	}
	for _, c := range cases {
		got := Round(c.in)
		if got != c.want {
			t.Errorf("Round(%f) == %f, want %f", c.in, got, c.want)
		}
	}
}

// RoundPlus -- Returns value rounded to nearest places significant digits.
func TestRoundPlus(t *testing.T) {
	cases := []struct {
		in1  float64
		in2  int
		want float64
	}{
		{100.5, 0, 101},
		{100.499999999, 2, 100.50},
		{-0.77777, 4, -0.7778},
		{-0.11111, 0, 0},
		{0, 1, 0},
	}
	for _, c := range cases {
		got := RoundPlus(c.in1, c.in2)
		if got != c.want {
			t.Errorf("RoundPlus(%f, %d) == %f, want %f", c.in1, c.in2, got, c.want)
		}
	}
}

// RoundInPlace -- Replaces value with rounded whole number value.
//func TestRoundInPlace(t *testing.T) {
