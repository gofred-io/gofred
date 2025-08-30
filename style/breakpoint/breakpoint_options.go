package breakpoint

type BreakpointValue[T any] struct {
	xs  T
	sm  T
	md  T
	lg  T
	xl  T
	xxl T
}

type BreakpointOptions[T any] func(breakpointValue *BreakpointValue[T])

func XS[T any](value T) BreakpointOptions[T] {
	return func(breakpointValue *BreakpointValue[T]) {
		breakpointValue.xs = value
	}
}

func SM[T any](value T) BreakpointOptions[T] {
	return func(breakpointValue *BreakpointValue[T]) {
		breakpointValue.sm = value
	}
}

func MD[T any](value T) BreakpointOptions[T] {
	return func(breakpointValue *BreakpointValue[T]) {
		breakpointValue.md = value
	}
}

func LG[T any](value T) BreakpointOptions[T] {
	return func(breakpointValue *BreakpointValue[T]) {
		breakpointValue.lg = value
	}
}

func XL[T any](value T) BreakpointOptions[T] {
	return func(breakpointValue *BreakpointValue[T]) {
		breakpointValue.xl = value
	}
}

func XXL[T any](value T) BreakpointOptions[T] {
	return func(breakpointValue *BreakpointValue[T]) {
		breakpointValue.xxl = value
	}
}

func (b *BreakpointValue[T]) Get(breakpoint BreakPoint) T {
	switch breakpoint {
	case xs:
		return b.xs
	case sm:
		return b.sm
	case md:
		return b.md
	case lg:
		return b.lg
	case xl:
		return b.xl
	case xxl:
		return b.xxl
	}

	return b.xs
}
