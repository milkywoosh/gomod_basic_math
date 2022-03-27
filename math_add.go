package gomod_basic_math

func math_add(data ...int) (result int) {
	result = 0
	for _, value := range data {
		result += value
	}
	return result
}
