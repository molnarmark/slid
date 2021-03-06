# Slid
[![Open Source Love](https://badges.frapsoft.com/os/v1/open-source.svg?v=102)](https://github.com/ellerbrock/open-source-badge/)
[![Open Source Love](https://badges.frapsoft.com/os/mit/mit.svg?v=102)](https://github.com/ellerbrock/open-source-badge/)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat-square)](http://makeapullrequest.com)

Slid is a static file server with routing.


### Example

```go
package main

import (
	"fmt"

	"github.com/molnarmark/slid"
)

func main() {
	config := &slid.Config{
		Port: 9898,
		Dir:  "files",
	}

	// Makes a new slid instance with the config
	app := slid.New(config)

	// Adding a route with a path and a specific handler function.
	app.Get("/", func(c *slid.Context) {
		c.Text("Homepage, with plain text sent.")
	})

	// You can also call the Serve method on the context to send back a file.
	app.Get("/anotherindex", func(c *slid.Context) {
		c.Serve("index")
	})

	// A quick way is also available to map a route to a filename
	app.GetFile("/index", "index")

	app.Run(func() {
		fmt.Println("Slid is running.")
	})
}
```





