package drawer

import (
	"fmt"

	"github.com/gofred-io/gofred/breakpoint"
)

type Option func(drawer *Drawer)

func Class(class string) Option {
	return func(drawer *Drawer) {
		drawer.AddClass(class)
	}
}

func ID(id string) Option {
	return func(drawer *Drawer) {
		drawer.SetID(id)
	}
}

func Transition(transition float64) Option {
	return func(drawer *Drawer) {
		drawer.transition = transition
		drawer.menu.UpdateStyleProperty("transition", fmt.Sprintf("all %.1fs", transition))
	}
}

func Width(opts ...breakpoint.BreakpointOptions[int]) Option {
	return func(drawer *Drawer) {
		for _, opt := range opts {
			opt(drawer.menu.Width)
		}
	}
}

func WidthP(opts ...breakpoint.BreakpointOptions[float64]) Option {
	return func(drawer *Drawer) {
		for _, opt := range opts {
			opt(drawer.menu.WidthP)
		}
	}
}
