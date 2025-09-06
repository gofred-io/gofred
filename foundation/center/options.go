package center

import (
	. "github.com/gofred-io/gofred/options"
)

func (c *FCenter) alignItems(alignItems AxisAlignmentType) *FCenter {
	c.BDiv.AlignItems(alignItems)
	return c
}

func (c *FCenter) display(display DisplayType) *FCenter {
	c.BDiv.Display(display)
	return c
}

func (c *FCenter) flex(flex int) *FCenter {
	c.BDiv.Flex(flex)
	return c
}

func (c *FCenter) justifyContent(justifyContent AxisAlignmentType) *FCenter {
	c.BDiv.JustifyContent(justifyContent)
	return c
}

func (c *FCenter) Class(class string) *FCenter {
	c.BDiv.Class(class)
	return c
}

func (c *FCenter) ID(id string) *FCenter {
	c.BDiv.ID(id)
	return c
}
