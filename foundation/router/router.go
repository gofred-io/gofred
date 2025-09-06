package router

import (
	"regexp"
	"strings"

	. "github.com/gofred-io/gofred/basic/div"
	. "github.com/gofred-io/gofred/hooks"
	. "github.com/gofred-io/gofred/listenable"
	. "github.com/gofred-io/gofred/widget"
)

type FRouter struct {
	*BDiv
	routePatterns   map[string]*regexp.Regexp
	routes          map[string]RouteBuilder
	notFoundBuilder RouteBuilder
}

func Router(opts ...RouterOption) *FRouter {
	router := &FRouter{
		BDiv:          Div(nil),
		routePatterns: make(map[string]*regexp.Regexp),
		routes:        make(map[string]RouteBuilder),
	}

	for _, opt := range opts {
		opt(router)
	}

	router.createListeners()
	return router
}

func (r *FRouter) createListeners() {
	navigate := UseNavigate()
	callback := func(path string) {
		var (
			ok           bool
			newWidget    Widget
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
			nilWidget := Nil
			newWidget = &BaseWidget{JSWidget: &nilWidget}
		}

		r.ReplaceWith(newWidget)
		r.BaseWidget = newWidget.GetBaseWidget()
	}

	listener := NewListener(callback)
	navigate.AddListener(listener)

	currentPath := Context().CurrentPath()
	callback(currentPath)
}

func (r *FRouter) parsePath(path string) (RouteBuilder, RouteParams) {
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
