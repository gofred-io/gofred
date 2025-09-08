package snackbar

import (
	"fmt"
	"time"

	"github.com/gofred-io/gofred/options"
)

type Option func(s *Snackbar)

func Class(class string) Option {
	return func(s *Snackbar) {
		s.AddClass(class)
	}
}

func Duration(duration time.Duration) Option {
	return func(s *Snackbar) {
		s.duration = duration
	}
}

func ID(id string) Option {
	return func(s *Snackbar) {
		s.SetID(id)
	}
}

func Position(position options.PositionType) Option {
	return func(s *Snackbar) {
		s.position = position
	}
}

func Transition(transition float64) Option {
	return func(s *Snackbar) {
		s.transition = transition
		s.UpdateStyleProperty("transition", fmt.Sprintf("all %.1fs", transition))
	}
}
