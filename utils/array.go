package utils

func Some[T any](data []T, f func(T) bool) bool {
	for _, e := range data {
		if f(e) {
			return true
		}
	}
	return false
}
