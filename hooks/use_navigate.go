package hooks

import (
	"slices"
	"sync"

	"github.com/gofred-io/gofred/application"
	"github.com/gofred-io/gofred/listenable"
)

var (
	navigate *Navigate
)

type Navigate struct {
	listeners []listenable.Listener[string]
	mutex     sync.RWMutex
	current   string
}

func UseNavigate() *Navigate {
	if navigate != nil {
		return navigate
	}

	navigate = &Navigate{
		current: application.Context().CurrentPath(),
	}

	application.Context().OnNavigate(func(path string) {
		navigate.current = path
		navigate.NotifyListeners(path)
	})

	return navigate
}

func (n *Navigate) AddListener(listener listenable.Listener[string]) {
	n.mutex.Lock()
	defer n.mutex.Unlock()

	n.listeners = append(n.listeners, listener)
}

func (n *Navigate) NotifyListeners(newValue string) {
	for _, listener := range n.listeners {
		select {
		case listener <- newValue:
		default:
		}
	}
}

func (n *Navigate) RemoveListener(listener listenable.Listener[string]) {
	n.mutex.Lock()
	defer n.mutex.Unlock()

	n.listeners = slices.DeleteFunc(n.listeners, func(l listenable.Listener[string]) bool {
		return l == listener
	})
}

func (n *Navigate) Path() string {
	return n.current
}
