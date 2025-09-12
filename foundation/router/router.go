package router

import (
	"regexp"
	"strings"

	"github.com/gofred-io/gofred/application"
	"github.com/gofred-io/gofred/foundation/container"
	"github.com/gofred-io/gofred/hooks"
	"github.com/gofred-io/gofred/listenable"
	"github.com/gofred-io/gofred/theme/theme_data"
)

type Router struct {
	application.BaseWidget
	routePatterns   map[string]*regexp.Regexp
	routes          map[string]RouteBuilder
	notFoundBuilder RouteBuilder
}

func New(opts ...Option) application.BaseWidget {
	router := &Router{
		BaseWidget:    application.New("div"),
		routePatterns: make(map[string]*regexp.Regexp),
		routes:        make(map[string]RouteBuilder),
	}

	for _, option := range opts {
		option(router)
	}

	application.AddPageLoadedListener(router.createListeners)
	return router.BaseWidget
}

func (r *Router) createListeners() {
	navigate := hooks.UseNavigate()
	callback := func(path string) {
		var (
			ok           bool
			newWidget    application.BaseWidget
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
			newWidget = container.New(application.Nil)
		}

		r.Widget.ReplaceWith(newWidget.Widget)
		r.Widget = newWidget.Widget
	}

	themeCallback := func(theme *theme_data.ThemeData) {
		callback(application.Context().CurrentPath())
	}

	listener := listenable.NewListener(callback)
	navigate.AddListener(listener)

	themeHook, _ := hooks.UseTheme()
	themeListener := listenable.NewListener(themeCallback)
	themeHook.AddListener(themeListener)

	currentPath := application.Context().CurrentPath()
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
