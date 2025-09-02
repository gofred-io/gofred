package spacing

type Spacing struct {
	Top    int
	Right  int
	Bottom int
	Left   int
}

func LRTB(left int, right int, top int, bottom int) Spacing {
	return Spacing{Top: top, Right: right, Bottom: bottom, Left: left}
}

func All(padding int) Spacing {
	return Spacing{Top: padding, Right: padding, Bottom: padding, Left: padding}
}

func Axis(horizontal int, vertical int) Spacing {
	return Spacing{Left: horizontal, Right: horizontal, Top: vertical, Bottom: vertical}
}

func Bottom(bottom int) Spacing {
	return Spacing{Bottom: bottom}
}

func Horizontal(horizontal int) Spacing {
	return Spacing{Left: horizontal, Right: horizontal}
}

func Left(left int) Spacing {
	return Spacing{Left: left}
}

func Right(right int) Spacing {
	return Spacing{Right: right}
}

func Top(top int) Spacing {
	return Spacing{Top: top}
}

func Vertical(vertical int) Spacing {
	return Spacing{Top: vertical, Bottom: vertical}
}
