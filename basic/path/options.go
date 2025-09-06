package path

import (
	. "github.com/gofred-io/gofred/options"
)

func (p *BPath) d(data string) *BPath {
	D(data)(p)
	return p
}

func (p *BPath) Class(class string) *BPath {
	Class(class)(p)
	return p
}

func (p *BPath) Fill(fill string) *BPath {
	Fill(fill)(p)
	return p
}

func (p *BPath) ID(id string) *BPath {
	ID(id)(p)
	return p
}
