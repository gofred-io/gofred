package router

import "github.com/gofred-io/gofred/widget"

type RouteBuilder func() widget.Widget
type Options func(router *Router)

func Route(path string, routeBuilder RouteBuilder) Options {
	return func(router *Router) {
		router.routes[path] = routeBuilder
	}
}
