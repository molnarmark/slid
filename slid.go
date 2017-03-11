package slid

import (
	"fmt"
	"net/http"
)

const (
	GET = "GET"
)

type Slid struct {
	Router *Router
	Config *Config
}

func New(cfg *Config) *Slid {
	return &Slid{
		Router: new(Router),
		Config: cfg,
	}
}

func (s *Slid) Get(path string, handler func(*Context)) {
	s.Router.routes = append(s.Router.routes, Route{"GET", path, handler})
}

func (s *Slid) GetFile(path, filename string) {
	s.Router.routes = append(s.Router.routes, Route{"GET", path, func(c *Context) {
		c.Serve(filename)
	}})
}

func (s *Slid) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := NewContext(r, w, s)

	for _, route := range s.Router.Routes() {
		if route.Path == r.URL.String() && route.Method == r.Method {
			route.Handler(ctx)
		}
	}
}

func (s *Slid) Run(callback func()) {
	addr := fmt.Sprintf("localhost:%d", s.Config.Port)
	callback()
	fmt.Println(http.ListenAndServe(addr, s))
}
