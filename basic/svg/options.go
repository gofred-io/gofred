package svg

import (
	. "github.com/gofred-io/gofred/breakpoint"
	. "github.com/gofred-io/gofred/options"
)

func (s *BSvg) Class(class string) *BSvg {
	Class(class)(s)
	return s
}

func (s *BSvg) Fill(fill string) *BSvg {
	Fill(fill)(s)
	return s
}

func (s *BSvg) Height(opts ...BreakpointOptions[int]) *BSvg {
	Height(opts...)(s)
	return s
}

func (s *BSvg) ID(id string) *BSvg {
	ID(id)(s)
	return s
}

func (s *BSvg) MaxWidth(opts ...BreakpointOptions[int]) *BSvg {
	MaxWidth(opts...)(s)
	return s
}

func (s *BSvg) UserSelect(userSelect UserSelectType) *BSvg {
	UserSelect(userSelect)(s)
	return s
}

func (s *BSvg) Width(opts ...BreakpointOptions[int]) *BSvg {
	Width(opts...)(s)
	return s
}

func (s *BSvg) WidthP(opts ...BreakpointOptions[float64]) *BSvg {
	WidthP(opts...)(s)
	return s
}
