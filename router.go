package slid

type Route struct {
	Method  string
	Path    string
	Handler func(*Context)
}

type Router struct {
	routes []Route
}

func (r *Router) Routes() []Route {
	return r.routes
}
