package pre

import (
	. "github.com/gofred-io/gofred/widget"
)

type BPre struct {
	*BaseWidget
}

func Pre(child Widget) *BPre {
	pre := &BPre{
		BaseWidget: New("pre"),
	}

	pre.AppendChild(child)
	return pre
}
