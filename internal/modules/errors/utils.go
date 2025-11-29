package errors

func Must[T comparable](r T, e *Error) T {
	if e != nil {
		panic(e)
	}
	return r
}
