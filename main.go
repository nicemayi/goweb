package main

import (
	"fmt"
	"goweb/framework"
	"html/template"
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

type student struct {
	Name string
	Age  int8
}

func FormatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d-%02d-%02d", year, month, day)
}

func main() {
	r := framework.New()
	r.Use(framework.Logger())
	r.SetFuncMap(template.FuncMap{
		"FormatAsDate": FormatAsDate,
	})
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./static")

	stu1 := &student{
		Name: "ZZZ",
		Age:  20,
	}
	stu2 := &student{
		Name: "WWW",
		Age:  36,
	}

	r.GET("/", func(c *framework.Context) {
		c.HTML(http.StatusOK, "css.tmpl", nil)
	})
	r.GET("/students", func(c *framework.Context) {
		c.HTML(http.StatusOK, "arr.tmpl", framework.H{
			"title":  "zzz",
			"stuArr": [2]*student{stu1, stu2},
		})
	})
	r.GET("/date", func(c *framework.Context) {
		c.HTML(http.StatusOK, "custom_func.tmpl", framework.H{
			"title": "www",
			"now":   time.Date(2020, 12, 1, 0, 0, 0, 0, time.UTC),
		})
	})

	v1 := r.Group("/v1")
	{
		v1.GET("/", func(c *framework.Context) {
			c.HTML(http.StatusOK, "<h1>Hello from framework!", nil)
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
