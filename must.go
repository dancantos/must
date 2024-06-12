package must

// Must implements a generic Must func. Panics if given an error, otherwise
// passes through the first arg.
func Must[T any](t T, err error) T {
	if err != nil {
		panic(err)
	}
	return t
}

// Ok implements a generic OK func. Panics if given false, otherwise passes
// through the first arg.
func Ok[T any](t T, ok bool) T {
	if !ok {
		panic(ok)
	}
	return t
}
