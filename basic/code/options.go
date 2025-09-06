package code

import (
	. "github.com/gofred-io/gofred/options"
)

func (c *BCode) Class(class string) *BCode {
	Class(class)(c)
	return c
}

func (c *BCode) FontColor(color string) *BCode {
	FontColor(color)(c)
	return c
}

func (c *BCode) FontFamily(fontFamily string) *BCode {
	FontFamily(fontFamily)(c)
	return c
}

func (c *BCode) FontSize(fontSize int) *BCode {
	FontSize(fontSize)(c)
	return c
}

func (c *BCode) FontWeight(fontWeight string) *BCode {
	FontWeight(fontWeight)(c)
	return c
}

func (c *BCode) ID(id string) *BCode {
	ID(id)(c)
	return c
}
