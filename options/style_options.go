package options

import (
	"github.com/gofred-io/gofred/application"
	"github.com/gofred-io/gofred/breakpoint"
	"github.com/gofred-io/gofred/options/spacing"
	"github.com/gofred-io/gofred/theme/theme_style"
)

func ButtonStyle(buttonStyle theme_style.ButtonStyle) Option {
	return func(widget application.BaseWidget) {
		if buttonStyle.BackgroundColor != nil {
			color := *buttonStyle.BackgroundColor
			BackgroundColor(color.String())(widget)
		}
		if buttonStyle.BorderColor != nil {
			color := *buttonStyle.BorderColor
			BorderColor(color.String())(widget)
		}
		if buttonStyle.BorderStyle != nil {
			BorderStyle(*buttonStyle.BorderStyle)(widget)
		}
		if buttonStyle.BorderWidth != nil {
			BorderWidth(*buttonStyle.BorderWidth)(widget)
		}
		if buttonStyle.BorderRadius != nil {
			BorderRadius(*buttonStyle.BorderRadius)(widget)
		}
		if buttonStyle.BoxShadow != nil {
			BoxShadow(*buttonStyle.BoxShadow)(widget)
		}
		if buttonStyle.Fill != nil {
			color := *buttonStyle.Fill
			Fill(color.String())(widget)
		}
		if buttonStyle.Flex != nil {
			Flex(*buttonStyle.Flex)(widget)
		}
		if buttonStyle.FlexDirection != nil {
			FlexDirection(*buttonStyle.FlexDirection)(widget)
		}
		if buttonStyle.FlexWrap != nil {
			FlexWrap(*buttonStyle.FlexWrap)(widget)
		}
		if buttonStyle.Height != nil {
			heights := []breakpoint.BreakpointOptions[int](*buttonStyle.Height)
			Height(heights...)(widget)
		}
		if buttonStyle.Margin != nil {
			margins := []breakpoint.BreakpointOptions[spacing.Spacing](*buttonStyle.Margin)
			Margin(margins...)(widget)
		}
		if buttonStyle.MaxWidth != nil {
			maxWidths := []breakpoint.BreakpointOptions[int](*buttonStyle.MaxWidth)
			MaxWidth(maxWidths...)(widget)
		}
		if buttonStyle.Overflow != nil {
			Overflow(*buttonStyle.Overflow)(widget)
		}
		if buttonStyle.Padding != nil {
			padding := []breakpoint.BreakpointOptions[spacing.Spacing](*buttonStyle.Padding)
			Padding(padding...)(widget)
		}
		if buttonStyle.Width != nil {
			widths := []breakpoint.BreakpointOptions[int](*buttonStyle.Width)
			Width(widths...)(widget)
		}
		if buttonStyle.WidthP != nil {
			widthPs := []breakpoint.BreakpointOptions[float64](*buttonStyle.WidthP)
			WidthP(widthPs...)(widget)
		}
		if buttonStyle.Visible != nil {
			visibles := []breakpoint.BreakpointOptions[bool](*buttonStyle.Visible)
			Visible(visibles...)(widget)
		}
	}
}

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
