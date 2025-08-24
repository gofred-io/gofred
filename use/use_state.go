package use

func State[T any](initialValue T) (*T, func(T)) {
	t := new(T)
	*t = initialValue

	setValue := func(value T) {
		*t = value
	}

	return t, setValue
}
