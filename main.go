package main

import (
	"goweb/framework"
	"net/http"
)

func main() {
	r := framework.New()

	r.GET("/", func(c *framework.Context) {
		c.HTML(http.StatusOK, "<h1>Hello from framework</h1>")
	})
	r.GET("/hello", func(c *framework.Context) {
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})
	r.GET("/hello/:name", func(c *framework.Context) {
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	})
	r.GET("/assets/*filepath", func(c *framework.Context) {
		c.JSON(http.StatusOK, framework.H{
			"filepath": c.Param("filepath"),
		})
	})
	r.POST("/login", func(c *framework.Context) {
		c.JSON(http.StatusOK, framework.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("passowrd"),
		})
	})

	r.Run(":9999")
}
