package hooks

import (
	"slices"
	"sync"
	"time"

	"github.com/gofred-io/gofred/breakpoint"
	"github.com/gofred-io/gofred/listenable"
	"github.com/gofred-io/gofred/utils"
	"github.com/gofred-io/gofred/widget"
)

var (
	breakpointValue *BreakpointValue
)

type BreakpointValue struct {
	listeners []listenable.Listener[breakpoint.BreakPoint]
	mutex     sync.RWMutex
	current   breakpoint.BreakPoint
}

func UseBreakpoint() *BreakpointValue {
	if breakpointValue != nil {
		return breakpointValue
	}

	breakpointValue = &BreakpointValue{
		listeners: make([]listenable.Listener[breakpoint.BreakPoint], 0),
		mutex:     sync.RWMutex{},
		current:   breakpoint.GetCurrentBreakPoint(),
	}

	c := make(chan int)

	widget.Context().OnResize(func() {
		width := widget.Context().ClientWidth()
		select {
		case c <- width:
		default:
		}
	})

	utils.SafeGo(func() {
		tick := time.NewTicker(time.Millisecond * 100)
		for range tick.C {
			width := <-c
			current := breakpoint.GetBreakPointFromWidth(width)
			if current != breakpointValue.current {
				breakpointValue.current = current
				breakpointValue.NotifyListeners(current)
			}
		}
	})

	return breakpointValue
}

func (s *BreakpointValue) AddListener(listener listenable.Listener[breakpoint.BreakPoint]) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.listeners = append(s.listeners, listener)
}

func (s *BreakpointValue) NotifyListeners(newValue breakpoint.BreakPoint) {
	for _, listener := range s.listeners {
		select {
		case listener <- newValue:
		default:
		}
	}
}

func (s *BreakpointValue) RemoveListener(listener listenable.Listener[breakpoint.BreakPoint]) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.listeners = slices.DeleteFunc(s.listeners, func(l listenable.Listener[breakpoint.BreakPoint]) bool {
		return l == listener
	})
}
