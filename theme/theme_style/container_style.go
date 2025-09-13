package theme_style

import (
	"github.com/gofred-io/gofred/breakpoint"
	"github.com/gofred-io/gofred/options/spacing"
	"github.com/gofred-io/gofred/theme"
	"github.com/gofred-io/gofred/theme/color"
)

type ContainerStyleCollection struct {
	Primary   ContainerStyle
	Secondary ContainerStyle
	Tertiary  ContainerStyle
}

type ContainerStyle struct {
	BackgroundColor themeValue[color.Color]
	BorderColor     themeValue[color.Color]
	BorderStyle     themeValue[theme.BorderStyleType]
	BorderWidth     themeValue[spacing.Spacing]
	BorderRadius    themeValue[int]
	BoxShadow       themeValue[string]
	Flex            themeValue[int]
	FlexDirection   themeValue[theme.FlexDirectionType]
	FlexWrap        themeValue[theme.FlexWrapType]
	Height          themeValue[[]breakpoint.BreakpointOptions[int]]
	Margin          themeValue[[]breakpoint.BreakpointOptions[spacing.Spacing]]
	MaxWidth        themeValue[[]breakpoint.BreakpointOptions[int]]
	Padding         themeValue[[]breakpoint.BreakpointOptions[spacing.Spacing]]
	Width           themeValue[[]breakpoint.BreakpointOptions[int]]
	WidthP          themeValue[[]breakpoint.BreakpointOptions[float64]]
	Visible         themeValue[[]breakpoint.BreakpointOptions[bool]]
}
