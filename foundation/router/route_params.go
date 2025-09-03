package router

type RouteParams map[string]string

func (r RouteParams) Get(name string) string {
	return r[name]
}
