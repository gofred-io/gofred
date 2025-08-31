package breakpoint

import "syscall/js"

type BreakPoint int

const (
	xs BreakPoint = iota
	sm
	md
	lg
	xl
	xxl
)

var (
	breakPointWidths = map[BreakPoint]int{
		xs:  0,
		sm:  480,
		md:  768,
		lg:  1280,
		xl:  1536,
		xxl: 1920,
	}
	BreakPoints = []BreakPoint{xs, sm, md, lg, xl, xxl}
)

func GetCurrent() BreakPoint {
	clientWidth := js.Global().Get("document").Get("documentElement").Get("clientWidth").Int()
	return GetFromWidth(clientWidth)
}

func GetFromWidth(width int) BreakPoint {
	for _, breakPoint := range BreakPoints {
		if width < breakPointWidths[breakPoint] {
			return breakPoint - 1
		}
	}

	return xs
}
