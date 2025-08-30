package router

import (
	"github.com/gofred-io/gofred/container"
	"github.com/gofred-io/gofred/hooks"
	"github.com/gofred-io/gofred/listenable"
	"github.com/gofred-io/gofred/widget"
)

type Router struct {
	widget.Widget
	routes map[string]RouteBuilder
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
		routeBuilder := r.routes[path]
		if routeBuilder != nil {
			widget := routeBuilder()
			r.Widget.ReplaceWith(widget)
			r.Widget = widget
		} else {
			widget := container.New(widget.Nil)
			r.Widget.ReplaceWith(widget)
			r.Widget = widget
		}
	}

	listener := listenable.NewListener(callback)
	navigate.AddListener(listener)

	currentPath := widget.Context().CurrentPath()
	callback(currentPath)
}
