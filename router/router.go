package router

import (
	"github.com/gofred-io/gofred/container"
	"github.com/gofred-io/gofred/hooks"
	"github.com/gofred-io/gofred/listenable"
	"github.com/gofred-io/gofred/widget"
)

type Router struct {
	widget.Widget
	routes          map[string]RouteBuilder
	notFoundBuilder RouteBuilder
}

func New(options ...Options) widget.Widget {
	router := &Router{
		Widget: widget.Context().CreateElement("div"),
		routes: make(map[string]RouteBuilder),
	}

	for _, option := range options {
		option(router)
	}
	router.createListeners()

	return router.Widget
}

func (r *Router) createListeners() {
	navigate := hooks.UseNavigate()
	callback := func(path string) {
		var (
			newWidget    widget.Widget
			routeBuilder = r.routes[path]
		)

		if routeBuilder != nil {
			newWidget = routeBuilder()
		} else if r.notFoundBuilder != nil {
			newWidget = r.notFoundBuilder()
		} else {
			newWidget = container.New(widget.Nil)
		}

		r.Widget.ReplaceWith(newWidget)
		r.Widget = newWidget
	}

	listener := listenable.NewListener(callback)
	navigate.AddListener(listener)

	currentPath := widget.Context().CurrentPath()
	callback(currentPath)
}
