package listenable

type Listenable[T any] interface {
	AddListener(listener Listener[T])
	NotifyListeners(value T)
	RemoveListener(listener Listener[T])
}
