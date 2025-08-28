package listenable

type Listener[T any] chan T
type ListenerFunc[T any] func(T)

func NewListener[T any](callback ListenerFunc[T]) Listener[T] {
	listener := make(chan T)

	go func() {
		for value := range listener {
			callback(value)
		}
	}()

	return listener
}
