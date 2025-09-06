package link

import (
	. "github.com/gofred-io/gofred/options"
)

func (link *FLink) Class(class string) *FLink {
	link.BAnchor.Class(class)
	return link
}

func (link *FLink) Flex(flex int) *FLink {
	link.BAnchor.Flex(flex)
	return link
}

func (link *FLink) FontColor(color string) *FLink {
	link.BAnchor.FontColor(color)
	return link
}

func (link *FLink) FontFamily(fontFamily string) *FLink {
	link.BAnchor.FontFamily(fontFamily)
	return link
}

func (link *FLink) FontSize(fontSize int) *FLink {
	link.BAnchor.FontSize(fontSize)
	return link
}

func (link *FLink) FontWeight(fontWeight string) *FLink {
	link.BAnchor.FontWeight(fontWeight)
	return link
}

func (link *FLink) Href(href string) *FLink {
	link.BAnchor.Href(href)
	return link
}

func (link *FLink) ID(id string) *FLink {
	link.BAnchor.ID(id)
	return link
}

func (link *FLink) LineHeight(lineHeight float64) *FLink {
	link.BAnchor.LineHeight(lineHeight)
	return link
}

func (link *FLink) OnClick(handler OnClickHandler) *FLink {
	link.BAnchor.OnClick(handler)
	return link
}

func (link *FLink) NewTab(newTab bool) *FLink {
	link.BAnchor.NewTab(newTab)
	return link
}

func (link *FLink) Tooltip(tooltip string) *FLink {
	link.BAnchor.Tooltip(tooltip)
	return link
}
