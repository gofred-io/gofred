package scaffold

import (
	"sync"

	"github.com/gofred-io/gofred/application"
	"github.com/gofred-io/gofred/foundation/drawer"
)

type Scaffold struct {
	drawers map[string]*drawer.Drawer
	mutex   sync.RWMutex
}

func New(child application.BaseWidget, opts ...Option) application.BaseWidget {
	if application.Context().Scaffold != nil {
		panic("You can only have one Scaffold in the application")
	}

	s := &Scaffold{
		drawers: make(map[string]*drawer.Drawer),
	}

	for _, option := range opts {
		option(s)
	}

	application.Context().Scaffold = s
	return child
}

func Get() *Scaffold {
	s, ok := application.Context().Scaffold.(*Scaffold)
	if !ok {
		return nil
	}
	return s
}

func (s *Scaffold) Refresh() {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	for _, drawer := range s.drawers {
		drawer.Refresh()
	}
}

func (s *Scaffold) Drawer(name string) *drawer.Drawer {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	return s.drawers[name]
}
