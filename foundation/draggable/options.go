package draggable

import (
	"github.com/gofred-io/gofred/basic/div"
	"github.com/gofred-io/gofred/breakpoint"
	"github.com/gofred-io/gofred/options"
	"github.com/gofred-io/gofred/options/spacing"
	"github.com/gofred-io/gofred/theme"
	"github.com/gofred-io/gofred/theme/theme_style"
)

type Option func(draggable *draggable)

func containerStyle(containerStyle theme_style.ContainerStyle) Option {
	return func(draggable *draggable) {
		draggable.opts = append(draggable.opts, div.ContainerStyle(containerStyle))
	}
}

func display(display theme.DisplayType) Option {
	return func(draggable *draggable) {
		draggable.opts = append(draggable.opts, div.Display(display))
	}
}

func setDraggable(isDraggable bool) Option {
	return func(draggable *draggable) {
		draggable.opts = append(draggable.opts, div.Draggable(isDraggable))
	}
}

func Class(class string) Option {
	return func(draggable *draggable) {
		draggable.opts = append(draggable.opts, div.Class(class))
	}
}

func Height(opts ...breakpoint.BreakpointOptions[int]) Option {
	return func(draggable *draggable) {
		draggable.opts = append(draggable.opts, div.Height(opts...))
	}
}

func ID(id string) Option {
	return func(draggable *draggable) {
		draggable.opts = append(draggable.opts, div.ID(id))
	}
}

func Margin(opts ...breakpoint.BreakpointOptions[spacing.Spacing]) Option {
	return func(draggable *draggable) {
		draggable.opts = append(draggable.opts, div.Margin(opts...))
	}
}

func MaxWidth(opts ...breakpoint.BreakpointOptions[int]) Option {
	return func(draggable *draggable) {
		draggable.opts = append(draggable.opts, div.MaxWidth(opts...))
	}
}

func OnDragEnd(handler options.OnDragEndHandler) Option {
	return func(draggable *draggable) {
		draggable.opts = append(draggable.opts, div.OnDragEnd(handler))
	}
}

func OnDrag(handler options.OnDragHandler) Option {
	return func(draggable *draggable) {
		draggable.opts = append(draggable.opts, div.OnDrag(handler))
	}
}

func OnDragStart(handler options.OnDragStartHandler) Option {
	return func(draggable *draggable) {
		draggable.opts = append(draggable.opts, div.OnDragStart(handler))
	}
}

func Overflow(overflow theme.OverflowType) Option {
	return func(draggable *draggable) {
		draggable.opts = append(draggable.opts, div.Overflow(overflow))
	}
}

func Padding(opts ...breakpoint.BreakpointOptions[spacing.Spacing]) Option {
	return func(draggable *draggable) {
		draggable.opts = append(draggable.opts, div.Padding(opts...))
	}
}

func Visible(opts ...breakpoint.BreakpointOptions[bool]) Option {
	return func(draggable *draggable) {
		draggable.opts = append(draggable.opts, div.Visible(opts...))
	}
}

func Width(opts ...breakpoint.BreakpointOptions[int]) Option {
	return func(draggable *draggable) {
		draggable.opts = append(draggable.opts, div.Width(opts...))
	}
}

func WidthP(opts ...breakpoint.BreakpointOptions[float64]) Option {
	return func(draggable *draggable) {
		draggable.opts = append(draggable.opts, div.WidthP(opts...))
	}
}
