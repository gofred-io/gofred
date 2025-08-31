package text

import (
	"github.com/gofred-io/gofred/options"
)

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

func ID(id string) Option {
	return func() options.Option {
		return options.ID(id)
	}
}

func UserSelect(userSelect options.UserSelectType) Option {
	return func() options.Option {
		return options.UserSelect(userSelect)
	}
}
