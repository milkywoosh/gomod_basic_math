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
	for pow > 2 {
		value *= tmp
		pow--
	}
	return int(value)
}
