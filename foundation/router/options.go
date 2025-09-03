package router

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/gofred-io/gofred/widget"
)

type RouteBuilder func(params RouteParams) widget.BaseWidget
type Options func(router *Router)

func getRoutePattern(path string) *regexp.Regexp {
	if !strings.Contains(path, ":") {
		return regexp.MustCompile(fmt.Sprintf("^%s$", path))
	}

	var parts []string

	for _, part := range strings.Split(path, "/") {
		if strings.HasPrefix(part, ":") {
			name := strings.TrimPrefix(part, ":")
			parts = append(parts, fmt.Sprintf("(?P<%s>[^/]*)", name))
		} else {
			parts = append(parts, part)
		}
	}

	joined := strings.Join(parts, "/")
	if strings.HasSuffix(joined, "/") {
		joined = joined[:len(joined)-1] + "[/]?"
	}

	return regexp.MustCompile("^" + joined + "$")
}

func Route(path string, routeBuilder RouteBuilder) Options {
	return func(router *Router) {
		if !strings.HasSuffix(path, "/") {
			path = path + "/"
		}

		router.routePatterns[path] = getRoutePattern(path)
		router.routes[path] = routeBuilder
	}
}

func NotFound(routeBuilder RouteBuilder) Options {
	return func(router *Router) {
		router.notFoundBuilder = routeBuilder
	}
}
