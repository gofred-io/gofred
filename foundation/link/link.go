package link

import (
	. "github.com/gofred-io/gofred/basic/anchor"
	. "github.com/gofred-io/gofred/widget"
)

type FLink struct {
	*BAnchor
}

func Link(child Widget) *FLink {
	link := &FLink{
		BAnchor: Anchor(child),
	}

	link.Class("gf-link")
	return link
}
