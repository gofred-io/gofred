package drawer

import (
	"fmt"

	. "github.com/gofred-io/gofred/breakpoint"
)

func (drawer *FDrawer) Class(class string) *FDrawer {
	drawer.AddClass(class)
	return drawer
}

func (drawer *FDrawer) ID(id string) *FDrawer {
	drawer.SetID(id)
	return drawer
}

func (drawer *FDrawer) Transition(transition float64) *FDrawer {
	drawer.transition = transition
	drawer.menu.GetBaseWidget().UpdateStyleProperty("transition", fmt.Sprintf("all %.1fs", transition))
	return drawer
}

func (drawer *FDrawer) Width(opts ...BreakpointOptions[int]) *FDrawer {
	drawer.menu.Width(opts...)
	return drawer
}

func (drawer *FDrawer) WidthP(opts ...BreakpointOptions[float64]) *FDrawer {
	drawer.menu.WidthP(opts...)
	return drawer
}
