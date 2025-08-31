package router

import (
	"github.com/gofred-io/gofred/foundation/container"
	"github.com/gofred-io/gofred/hooks"
	"github.com/gofred-io/gofred/listenable"
	"github.com/gofred-io/gofred/widget"
)

type Router struct {
	widget.BaseWidget
	routes          map[string]RouteBuilder
	notFoundBuilder RouteBuilder
}

func New(opts ...Options) widget.BaseWidget {
	router := &Router{
		BaseWidget: widget.New("div"),
		routes:     make(map[string]RouteBuilder),
	}

	for _, option := range opts {
		option(router)
	}
	router.createListeners()

	return router.BaseWidget
}

func (r *Router) createListeners() {
	navigate := hooks.UseNavigate()
	callback := func(path string) {
		var (
			newWidget    widget.BaseWidget
			routeBuilder = r.routes[path]
		)

		if routeBuilder != nil {
			newWidget = routeBuilder()
		} else if r.notFoundBuilder != nil {
			newWidget = r.notFoundBuilder()
		} else {
			newWidget = container.New(widget.Nil)
		}

		r.Widget.ReplaceWith(newWidget.Widget)
		r.Widget = newWidget.Widget
	}

	listener := listenable.NewListener(callback)
	navigate.AddListener(listener)

	currentPath := widget.Context().CurrentPath()
	callback(currentPath)
}
