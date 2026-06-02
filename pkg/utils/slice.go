package utils

func MapSlice[T any, S any](input []T, fn func(T) S) []S {
	res := make([]S, len(input))
	for idx, val := range input {
		res[idx] = fn(val)
	}
	return res
}
