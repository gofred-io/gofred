package hooks

import (
	"slices"
	"sync"

	"github.com/gofred-io/gofred/listenable"
	"github.com/gofred-io/gofred/theme/theme_data"
)

var (
	theme *ThemeHook
)

type ThemeHook struct {
	listeners []listenable.Listener[*theme_data.ThemeData]
	mutex     sync.RWMutex
	current   theme_data.ThemeData
}

func UseTheme() (*ThemeHook, func(*theme_data.ThemeData)) {
	if theme != nil {
		return theme, theme.setThemeData
	}

	theme = &ThemeHook{
		listeners: make([]listenable.Listener[*theme_data.ThemeData], 0),
		mutex:     sync.RWMutex{},
	}

	return theme, theme.setThemeData
}

func (t *ThemeHook) AddListener(listener listenable.Listener[*theme_data.ThemeData]) {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	t.listeners = append(t.listeners, listener)
}

func (t *ThemeHook) NotifyListeners(newValue *theme_data.ThemeData) {
	for _, listener := range t.listeners {
		select {
		case listener <- newValue:
		default:
		}
	}
}

func (t *ThemeHook) RemoveListener(listener listenable.Listener[*theme_data.ThemeData]) {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	t.listeners = slices.DeleteFunc(t.listeners, func(l listenable.Listener[*theme_data.ThemeData]) bool {
		return l == listener
	})
}

func (t *ThemeHook) ThemeData() *theme_data.ThemeData {
	return &t.current
}

func (t *ThemeHook) setThemeData(themeData *theme_data.ThemeData) {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	t.current = *themeData
	t.NotifyListeners(themeData)
}
