package options

import (
	"github.com/gofred-io/gofred/application"
	"github.com/gofred-io/gofred/breakpoint"
	"github.com/gofred-io/gofred/options/spacing"
	"github.com/gofred-io/gofred/theme/theme_style"
)

func ContainerStyle(containerStyle theme_style.ContainerStyle) Option {
	return func(widget application.BaseWidget) {
		if containerStyle.BackgroundColor != nil {
			color := *containerStyle.BackgroundColor
			BackgroundColor(color.String())(widget)
		}
		if containerStyle.BorderColor != nil {
			color := *containerStyle.BorderColor
			BorderColor(color.String())(widget)
		}
		if containerStyle.BorderStyle != nil {
			BorderStyle(*containerStyle.BorderStyle)(widget)
		}
		if containerStyle.BorderWidth != nil {
			BorderWidth(*containerStyle.BorderWidth)(widget)
		}
		if containerStyle.BorderRadius != nil {
			BorderRadius(*containerStyle.BorderRadius)(widget)
		}
		if containerStyle.BoxShadow != nil {
			BoxShadow(*containerStyle.BoxShadow)(widget)
		}
		if containerStyle.Flex != nil {
			Flex(*containerStyle.Flex)(widget)
		}
		if containerStyle.FlexDirection != nil {
			FlexDirection(*containerStyle.FlexDirection)(widget)
		}
		if containerStyle.FlexWrap != nil {
			FlexWrap(*containerStyle.FlexWrap)(widget)
		}
		if containerStyle.Height != nil {
			heights := []breakpoint.BreakpointOptions[int](*containerStyle.Height)
			Height(heights...)(widget)
		}
		if containerStyle.Margin != nil {
			margins := []breakpoint.BreakpointOptions[spacing.Spacing](*containerStyle.Margin)
			Margin(margins...)(widget)
		}
		if containerStyle.MaxWidth != nil {
			maxWidths := []breakpoint.BreakpointOptions[int](*containerStyle.MaxWidth)
			MaxWidth(maxWidths...)(widget)
		}
		if containerStyle.Padding != nil {
			padding := []breakpoint.BreakpointOptions[spacing.Spacing](*containerStyle.Padding)
			Padding(padding...)(widget)
		}
		if containerStyle.Width != nil {
			widths := []breakpoint.BreakpointOptions[int](*containerStyle.Width)
			Width(widths...)(widget)
		}
		if containerStyle.WidthP != nil {
			widthPs := []breakpoint.BreakpointOptions[float64](*containerStyle.WidthP)
			WidthP(widthPs...)(widget)
		}
		if containerStyle.Visible != nil {
			visibles := []breakpoint.BreakpointOptions[bool](*containerStyle.Visible)
			Visible(visibles...)(widget)
		}
	}
}

func TextStyle(textStyle theme_style.TextStyle) Option {
	return func(widget application.BaseWidget) {
		if textStyle.BackgroundColor != nil {
			color := *textStyle.BackgroundColor
			BackgroundColor(color.String())(widget)
		}
		if textStyle.Color != nil {
			color := *textStyle.Color
			FontColor(color.String())(widget)
		}
		if textStyle.FontSize != nil {
			FontSize(*textStyle.FontSize)(widget)
		}
		if textStyle.FontWeight != nil {
			FontWeight(*textStyle.FontWeight)(widget)
		}
		if textStyle.FontFamily != nil {
			FontFamily(*textStyle.FontFamily)(widget)
		}
		if textStyle.LetterSpacing != nil {
			LetterSpacing(*textStyle.LetterSpacing)(widget)
		}
		if textStyle.LineHeight != nil {
			LineHeight(*textStyle.LineHeight)(widget)
		}
		if textStyle.Overflow != nil {
			Overflow(*textStyle.Overflow)(widget)
		}
		if textStyle.TextAlign != nil {
			TextAlign(*textStyle.TextAlign)(widget)
		}
	}
}
