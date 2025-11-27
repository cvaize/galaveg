package utils

func Must[T comparable](r T, e error) T {
	if e != nil {
		panic(e)
	}
	return r
}
