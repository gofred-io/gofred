package breakpoint

import "github.com/gofred-io/gofred/widget"

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
	BreakPoints      = []BreakPoint{xs, sm, md, lg, xl, xxl}
	breakPointWidths = map[BreakPoint]int{
		xs:  0,
		sm:  480,
		md:  768,
		lg:  1280,
		xl:  1536,
		xxl: 1920,
	}
)

func GetCurrent() BreakPoint {
	clientWidth := widget.Context().ClientWidth()
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
