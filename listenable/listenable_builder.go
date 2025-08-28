package listenable

import (
	"github.com/gofred-io/gofred/widget"
)

type ListenableBuilderCallback[T any] func() widget.Widget

type ListenableBuilder[T any] struct {
	listeners []Listener[T]
}

func Builder[T any](listenable Listenable[T], builder ListenableBuilderCallback[T]) widget.Widget {
	refWidget := builder()

	listener := NewListener(func(value T) {
		updatedWidget := builder()
		refWidget.ReplaceWith(updatedWidget)
		refWidget = updatedWidget
	})

	listenable.AddListener(listener)
	return refWidget
}
