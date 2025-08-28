package hooks

import (
	"slices"
	"sync"

	"github.com/gofred-io/gofred/listenable"
)

type StateValue[T any] struct {
	listeners []listenable.Listener[T]
	mutex     sync.RWMutex
	value     T
}

func UseState[T any](initialValue T) (*StateValue[T], func(T)) {
	state := &StateValue[T]{value: initialValue}

	return state, func(value T) {
		state.value = value
		state.NotifyListeners(value)
	}
}

func (s *StateValue[T]) AddListener(listener listenable.Listener[T]) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.listeners = append(s.listeners, listener)
}

func (s *StateValue[T]) NotifyListeners(newValue T) {
	for _, listener := range s.listeners {
		select {
		case listener <- newValue:
		default:
		}
	}
}

func (s *StateValue[T]) RemoveListener(listener listenable.Listener[T]) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.listeners = slices.DeleteFunc(s.listeners, func(l listenable.Listener[T]) bool {
		return l == listener
	})
}

func (s *StateValue[T]) Value() T {
	return s.value
}
