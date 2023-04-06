package main

import (
	"fmt"
	_ "log"
	"webfra/router"
)

func main() {
	r := router.New()

	r.GET("/", func(c *router.Context) {
		fmt.Fprintf(c.Writer, "URL.Path = %q\n", c.Req.URL.Path)
	})

	r.GET("/hello", func(c *router.Context) {
		for k, v := range c.Req.Header {
			fmt.Fprintf(c.Writer, "Header[%q] = %q\n", k, v)
		}
	})

	r.Run(":8080")

}
