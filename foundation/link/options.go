package link

import (
	"github.com/gofred-io/gofred/basic/anchor"
	"github.com/gofred-io/gofred/options"
)

type Option func(link *link)

func Class(class string) Option {
	return func(link *link) {
		link.opts = append(link.opts, anchor.Class(class))
	}
}

func Flex(flex int) Option {
	return func(link *link) {
		link.opts = append(link.opts, anchor.Flex(flex))
	}
}

func FontColor(color string) Option {
	return func(link *link) {
		link.opts = append(link.opts, anchor.FontColor(color))
	}
}

func FontFamily(fontFamily string) Option {
	return func(link *link) {
		link.opts = append(link.opts, anchor.FontFamily(fontFamily))
	}
}

func FontSize(fontSize int) Option {
	return func(link *link) {
		link.opts = append(link.opts, anchor.FontSize(fontSize))
	}
}

func FontWeight(fontWeight string) Option {
	return func(link *link) {
		link.opts = append(link.opts, anchor.FontWeight(fontWeight))
	}
}

func Href(href string) Option {
	return func(link *link) {
		link.opts = append(link.opts, anchor.Href(href))
	}
}

func ID(id string) Option {
	return func(link *link) {
		link.opts = append(link.opts, anchor.ID(id))
	}
}

func Label(label string) Option {
	return func(link *link) {
		link.opts = append(link.opts, anchor.Label(label))
	}
}

func LineHeight(lineHeight float64) Option {
	return func(link *link) {
		link.opts = append(link.opts, anchor.LineHeight(lineHeight))
	}
}

func OnClick(handler options.OnClickHandler) Option {
	return func(link *link) {
		link.opts = append(link.opts, anchor.OnClick(handler))
	}
}

func NewTab(newTab bool) Option {
	return func(link *link) {
		link.opts = append(link.opts, anchor.NewTab(newTab))
	}
}

func Tooltip(tooltip string) Option {
	return func(link *link) {
		link.opts = append(link.opts, anchor.Tooltip(tooltip))
	}
}
