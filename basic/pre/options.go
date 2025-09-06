package pre

import (
	. "github.com/gofred-io/gofred/options"
)

func (p *BPre) Class(class string) *BPre {
	Class(class)(p)
	return p
}

func (p *BPre) ID(id string) *BPre {
	ID(id)(p)
	return p
}
