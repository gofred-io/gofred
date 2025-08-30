package breakpoint

import "github.com/gofred-io/gofred/widget"

type BreakPoint int

const (
	XS BreakPoint = iota
	SM
	MD
	LG
	XL
	XXL
)

var (
	BreakPoints      = []BreakPoint{XS, SM, MD, LG, XL, XXL}
	breakPointWidths = map[BreakPoint]int{
		XS:  0,
		SM:  480,
		MD:  768,
		LG:  1280,
		XL:  1536,
		XXL: 1920,
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

	return XS
}
