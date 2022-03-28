package gomod_basic_math

func Math_add(data ...int) (result int) {
	result = 0
	for _, value := range data {
		result += value
	}
	return result
}

func Math_power(value float64, pow int) int {
	tmp := value
	for pow > 1 {
		value *= tmp
		pow--
	}
	return int(value)
}
func Math_divide(value int, divider int) int {
	if divider > value {
		return 0
	}
	return value / divider
}

func Math_modulo(value int, mod int) int {
	return value % mod
}
