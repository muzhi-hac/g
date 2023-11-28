package main

import (
	"gee"
)

func main() {
	router := gee.New()
	router.GET("/1", func(c *gee.Context) {
		c.JSON(200, gee.H{
			"name": "wangzeyu",
		})
	})
	v1 := router.Group("/hello")
	v1.GET("/1", func(c *gee.Context) {
		c.JSON(200, gee.H{
			"name": "wangzeyu",
		})
	})
	//router.GET("/", func(c *gee.Context) {
	//	c.JSON(200, gee.H{})
	//	//fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
	//})
	//
	//router.GET("/hello", func(c *gee.Context) {
	//	for k, v := range c.Req.Header {
	//		fmt.Printf("Header[%q] = %q\n", k, v)
	//	}
	//})

	router.Run(":8080")
}
