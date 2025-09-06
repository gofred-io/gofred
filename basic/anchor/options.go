package anchor

import (
	. "github.com/gofred-io/gofred/options"
)

func (a *BAnchor) Class(class string) *BAnchor {
	Class(class)(a)
	return a
}

func (a *BAnchor) Flex(flex int) *BAnchor {
	Flex(flex)(a)
	return a
}

func (a *BAnchor) FontColor(color string) *BAnchor {
	FontColor(color)(a)
	return a
}

func (a *BAnchor) FontFamily(fontFamily string) *BAnchor {
	FontFamily(fontFamily)(a)
	return a
}

func (a *BAnchor) FontSize(fontSize int) *BAnchor {
	FontSize(fontSize)(a)
	return a
}

func (a *BAnchor) FontWeight(fontWeight string) *BAnchor {
	FontWeight(fontWeight)(a)
	return a
}

func (a *BAnchor) Href(href string) *BAnchor {
	Href(href)(a)
	return a
}

func (a *BAnchor) ID(id string) *BAnchor {
	ID(id)(a)
	return a
}

func (a *BAnchor) LineHeight(lineHeight float64) *BAnchor {
	LineHeight(lineHeight)(a)
	return a
}

func (a *BAnchor) OnClick(handler OnClickHandler) *BAnchor {
	OnClick(handler)(a)
	return a
}

func (a *BAnchor) NewTab(newTab bool) *BAnchor {
	NewTab(newTab)(a)
	return a
}

func (a *BAnchor) Tooltip(tooltip string) *BAnchor {
	Tooltip(tooltip)(a)
	return a
}
