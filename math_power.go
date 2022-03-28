package gomod_basic_math

func Math_power(value float64, pow int) int {
	for i := pow; i > 1; i-- {
		value *= value
	}
	return int(value)
}
