package listenable

import "github.com/gofred-io/gofred/application"

type ListenableBuilderCallback[T any] func() application.BaseWidget

type ListenableBuilder[T any] struct {
	listeners []Listener[T]
}

func Builder[T any](listenable Listenable[T], builder ListenableBuilderCallback[T]) application.BaseWidget {
	refWidget := builder()

	listener := NewListener(func(value T) {
		updatedWidget := builder()
		refWidget.ReplaceWith(updatedWidget.Widget)
		refWidget = updatedWidget
	})

	listenable.AddListener(listener)
	return refWidget
}
