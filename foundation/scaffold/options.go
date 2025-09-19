package scaffold

import "github.com/gofred-io/gofred/foundation/drawer"

type Option func(scaffold *Scaffold)

func Drawer(name string, drawer *drawer.Drawer) Option {
	return func(scaffold *Scaffold) {
		scaffold.mutex.Lock()
		defer scaffold.mutex.Unlock()
		scaffold.drawers[name] = drawer
	}
}
