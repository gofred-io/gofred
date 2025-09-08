package snackbar

import (
	"time"

	"github.com/gofred-io/gofred/options"
	"github.com/gofred-io/gofred/widget"
)

type Snackbar struct {
	widget.BaseWidget
	child      *widget.BaseWidget
	duration   time.Duration
	position   options.PositionType
	transition float64
}

func New(opts ...Option) *Snackbar {
	s := &Snackbar{
		BaseWidget: widget.New("div"),
		position:   options.PositionTypeBottomRight,
	}

	opts = append([]Option{
		Class("gf-snackbar-container"),
		Duration(3 * time.Second),
		Transition(0.5),
	}, opts...)

	for _, option := range opts {
		option(s)
	}

	s.setInitialPosition()

	widget.Context().Root.AppendChild(s.Widget)
	return s
}

func (s *Snackbar) setInitialPosition() {

	switch s.position {
	case options.PositionTypeTopLeft:
		s.UpdateStyleProperty("top", "0")
		s.UpdateStyleProperty("left", "-100%")
	case options.PositionTypeTopRight:
		s.UpdateStyleProperty("right", "-100%")
		s.UpdateStyleProperty("top", "0")
	case options.PositionTypeBottomLeft:
		s.UpdateStyleProperty("left", "-100%")
		s.UpdateStyleProperty("bottom", "0")
	case options.PositionTypeBottomRight:
		s.UpdateStyleProperty("right", "-100%")
		s.UpdateStyleProperty("bottom", "0")
	case options.PositionTypeTop:
		s.UpdateStyleProperty("left", "50%")
		s.UpdateStyleProperty("transform", "translateX(-50%)")
		s.UpdateStyleProperty("top", "-100%")
	case options.PositionTypeBottom:
		s.UpdateStyleProperty("left", "50%")
		s.UpdateStyleProperty("transform", "translateX(-50%)")
		s.UpdateStyleProperty("bottom", "-100%")
	case options.PositionTypeLeft:
		s.UpdateStyleProperty("left", "-100%")
		s.UpdateStyleProperty("top", "50%")
		s.UpdateStyleProperty("transform", "translateY(-50%)")
	case options.PositionTypeRight:
		s.UpdateStyleProperty("right", "-100%")
		s.UpdateStyleProperty("top", "50%")
		s.UpdateStyleProperty("transform", "translateY(-50%)")
	}
}

func (s *Snackbar) setFinalPosition() {
	switch s.position {
	case options.PositionTypeTopLeft:
		s.UpdateStyleProperty("left", "0")
	case options.PositionTypeTopRight:
		s.UpdateStyleProperty("right", "0")
	case options.PositionTypeBottomLeft:
		s.UpdateStyleProperty("left", "0")
	case options.PositionTypeBottomRight:
		s.UpdateStyleProperty("right", "0")
	case options.PositionTypeTop:
		s.UpdateStyleProperty("top", "0")
	case options.PositionTypeBottom:
		s.UpdateStyleProperty("bottom", "0")
	case options.PositionTypeLeft:
		s.UpdateStyleProperty("left", "0")
	case options.PositionTypeRight:
		s.UpdateStyleProperty("right", "0")
	}
}

func (s *Snackbar) Show(child widget.BaseWidget) {
	if s.child != nil {
		s.child.Remove()
	}

	s.child = &child
	s.AppendChild(child.Widget)
	s.setFinalPosition()

	time.AfterFunc(s.duration, func() {
		if s.child == &child {
			s.setInitialPosition()
		}
	})
}
