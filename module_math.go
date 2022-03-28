package gomod_basic_math

// func Math_add(data ...int) (result int) {
// 	result = 0
// 	for _, value := range data {
// 		result += value
// 	}
// 	return result
// }

func Math_power(value float64, pow int) int {
	for i := pow; i > 1; i-- {
		value *= value
	}
	return int(value)
}
