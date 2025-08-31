package anchor

import "github.com/gofred-io/gofred/options"

type Option options.OptionWrapper

func Class(class string) Option {
	return func() options.Option {
		return options.Class(class)
	}
}

func FontColor(color string) Option {
	return func() options.Option {
		return options.FontColor(color)
	}
}

func FontFamily(fontFamily string) Option {
	return func() options.Option {
		return options.FontFamily(fontFamily)
	}
}

func FontSize(fontSize int) Option {
	return func() options.Option {
		return options.FontSize(fontSize)
	}
}

func FontWeight(fontWeight string) Option {
	return func() options.Option {
		return options.FontWeight(fontWeight)
	}
}

func Href(href string) Option {
	return func() options.Option {
		return options.Href(href)
	}
}

func ID(id string) Option {
	return func() options.Option {
		return options.ID(id)
	}
}

func LineHeight(lineHeight float64) Option {
	return func() options.Option {
		return options.LineHeight(lineHeight)
	}
}

func OnClick(handler options.OnClickHandler) Option {
	return func() options.Option {
		return options.OnClick(handler)
	}
}

func NewTab(newTab bool) Option {
	return func() options.Option {
		return options.NewTab(newTab)
	}
}

func Tooltip(tooltip string) Option {
	return func() options.Option {
		return options.Tooltip(tooltip)
	}
}
