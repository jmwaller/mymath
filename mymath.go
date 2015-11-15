package mymath

import "math"

// Round -- Returns value rounded to nearest whole number.
func Round(f float64) float64 {
	return math.Floor(f + 0.5)
}

// RoundPlus -- Returns value rounded to nearest places significant digits.
func RoundPlus(f float64, places int) float64 {
	shift := math.Pow(10, float64(places))
	return Round(f*shift) / shift
}

// RoundInPlace -- Replaces value with rounded whole number value.
func RoundInPlace(fPtr *float64) {
	*fPtr = math.Floor(*fPtr + 0.5)
}
