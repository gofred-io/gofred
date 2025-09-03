package router

import (
	"regexp"
	"strings"

	"github.com/gofred-io/gofred/foundation/container"
	"github.com/gofred-io/gofred/hooks"
	"github.com/gofred-io/gofred/listenable"
	"github.com/gofred-io/gofred/widget"
)

type Router struct {
	widget.BaseWidget
	routePatterns   map[string]*regexp.Regexp
	routes          map[string]RouteBuilder
	notFoundBuilder RouteBuilder
}

func New(opts ...Options) widget.BaseWidget {
	router := &Router{
		BaseWidget:    widget.New("div"),
		routePatterns: make(map[string]*regexp.Regexp),
		routes:        make(map[string]RouteBuilder),
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
			ok           bool
			newWidget    widget.BaseWidget
			routeBuilder RouteBuilder
			params       RouteParams
		)

		if !strings.HasSuffix(path, "/") {
			path = path + "/"
		}

		routeBuilder, ok = r.routes[path]
		if !ok {
			routeBuilder, params = r.parsePath(path)
		}

		if routeBuilder != nil {
			newWidget = routeBuilder(params)
		} else if r.notFoundBuilder != nil {
			newWidget = r.notFoundBuilder(params)
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

func (r *Router) parsePath(path string) (RouteBuilder, RouteParams) {
	for route, pattern := range r.routePatterns {
		matches := pattern.FindStringSubmatch(path)
		if matches == nil {
			continue
		}

		params := make(map[string]string)
		for i, name := range pattern.SubexpNames() {
			if i != 0 && name != "" && matches[i] != "" {
				params[name] = matches[i]
			}
		}

		return r.routes[route], params
	}

	return nil, nil
}
