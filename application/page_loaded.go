package application

import "sync"

type PageLoadedCallback func()
type PageLoadedListeners struct {
	listeners []PageLoadedCallback
	mutex     sync.RWMutex
}

var (
	pageLoadedListeners = &PageLoadedListeners{
		listeners: make([]PageLoadedCallback, 0),
		mutex:     sync.RWMutex{},
	}
)

func AddPageLoadedListener(listener PageLoadedCallback) {
	pageLoadedListeners.mutex.Lock()
	defer pageLoadedListeners.mutex.Unlock()

	pageLoadedListeners.listeners = append(pageLoadedListeners.listeners, listener)
}

func notifyPageLoaded() {
	for _, listener := range pageLoadedListeners.listeners {
		listener()
	}
}
