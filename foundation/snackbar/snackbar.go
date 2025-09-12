package snackbar

import (
	"time"

	"github.com/gofred-io/gofred/application"
	"github.com/gofred-io/gofred/theme"
)

type Snackbar struct {
	application.BaseWidget
	child      *application.BaseWidget
	duration   time.Duration
	position   theme.PositionType
	transition float64
}

func New(opts ...Option) *Snackbar {
	s := &Snackbar{
		BaseWidget: application.New("div"),
		position:   theme.PositionTypeBottomRight,
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

	application.Context().Root.AppendChild(s.Widget)
	return s
}

func (s *Snackbar) setInitialPosition() {

	switch s.position {
	case theme.PositionTypeTopLeft:
		s.UpdateStyleProperty("top", "0")
		s.UpdateStyleProperty("left", "-100%")
	case theme.PositionTypeTopRight:
		s.UpdateStyleProperty("right", "-100%")
		s.UpdateStyleProperty("top", "0")
	case theme.PositionTypeBottomLeft:
		s.UpdateStyleProperty("left", "-100%")
		s.UpdateStyleProperty("bottom", "0")
	case theme.PositionTypeBottomRight:
		s.UpdateStyleProperty("right", "-100%")
		s.UpdateStyleProperty("bottom", "0")
	case theme.PositionTypeTop:
		s.UpdateStyleProperty("left", "50%")
		s.UpdateStyleProperty("transform", "translateX(-50%)")
		s.UpdateStyleProperty("top", "-100%")
	case theme.PositionTypeBottom:
		s.UpdateStyleProperty("left", "50%")
		s.UpdateStyleProperty("transform", "translateX(-50%)")
		s.UpdateStyleProperty("bottom", "-100%")
	case theme.PositionTypeLeft:
		s.UpdateStyleProperty("left", "-100%")
		s.UpdateStyleProperty("top", "50%")
		s.UpdateStyleProperty("transform", "translateY(-50%)")
	case theme.PositionTypeRight:
		s.UpdateStyleProperty("right", "-100%")
		s.UpdateStyleProperty("top", "50%")
		s.UpdateStyleProperty("transform", "translateY(-50%)")
	}
}

func (s *Snackbar) setFinalPosition() {
	switch s.position {
	case theme.PositionTypeTopLeft:
		s.UpdateStyleProperty("left", "0")
	case theme.PositionTypeTopRight:
		s.UpdateStyleProperty("right", "0")
	case theme.PositionTypeBottomLeft:
		s.UpdateStyleProperty("left", "0")
	case theme.PositionTypeBottomRight:
		s.UpdateStyleProperty("right", "0")
	case theme.PositionTypeTop:
		s.UpdateStyleProperty("top", "0")
	case theme.PositionTypeBottom:
		s.UpdateStyleProperty("bottom", "0")
	case theme.PositionTypeLeft:
		s.UpdateStyleProperty("left", "0")
	case theme.PositionTypeRight:
		s.UpdateStyleProperty("right", "0")
	}
}

func (s *Snackbar) Show(child application.BaseWidget) {
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
