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
			t.Errorf("Reverse(%f) == %f, want %f", c.in, got, c.want)
		}
	}
}

// RoundPlus -- Returns value rounded to nearest places significant digits.
//func TestRoundPlus(t *testing.T) {
// RoundInPlace -- Replaces value with rounded whole number value.
//func TestRoundInPlace(t *testing.T) {
