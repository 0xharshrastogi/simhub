package utils

// map is used to map values to their corresponding type, first input is value type, second input is index
type MapFunc[T any, TResult any] func(v T, i int) TResult

func Map[T any, TResult any](s []T, mapCb MapFunc[T, TResult]) []TResult {
	m := make([]TResult, 0, len(s))
	for i, v := range s {
		m = append(m, mapCb(v, i))
	}
	return m
}
