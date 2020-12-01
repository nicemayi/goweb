package main

import (
	"goweb/framework"
	"log"
	"net/http"
	"time"
)

func onlyForV2() framework.HandlerFunc {
	return func(c *framework.Context) {
		t := time.Now()
		c.Fail(500, "Internal Server Error")
		log.Printf("[%d] %s in %v for group v2", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}

func main() {
	r := framework.New()
	r.Use(framework.Logger())
	r.GET("/index", func(c *framework.Context) {
		c.HTML(http.StatusOK, "<h1>Index Page</h1>")
	})

	v1 := r.Group("/v1")
	{
		v1.GET("/", func(c *framework.Context) {
			c.HTML(http.StatusOK, "<h1>Hello from framework!")
		})

		v1.GET("/hello", func(c *framework.Context) {
			// /hello?name=zhewang
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
		})
	}

	v2 := r.Group("/v2")
	{
		v2.GET("/hello/:name", func(c *framework.Context) {
			// /hello/zhewang
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
		v2.POST("/login", func(c *framework.Context) {
			c.JSON(http.StatusOK, framework.H{
				"username": c.PostForm("username"),
				"password": c.PostForm("password"),
			})
		})
	}

	r.Run(":9999")
}
