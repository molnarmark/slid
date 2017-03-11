package slid

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Context struct {
	Request  *http.Request
	Response http.ResponseWriter
	App      *Slid
}

func NewContext(r *http.Request, rw http.ResponseWriter, s *Slid) *Context {
	return &Context{r, rw, s}
}

func (c *Context) Text(text string) {
	fmt.Fprintf(c.Response, text)
}

func (c *Context) Serve(name string) {
	dir, _ := os.Getwd()
	file, _ := ioutil.ReadFile(dir + "/" + c.App.Config.Dir + "/" + name + ".html")
	fmt.Fprintf(c.Response, string(file))
}
