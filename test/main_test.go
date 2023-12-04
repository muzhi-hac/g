package test

import (
	"gee"
	"github.com/gin-gonic/gin"
	"testing"
)

func BenchmarkTestGee(b *testing.B) {
	router := gee.New()
	router.GET("/1", func(c *gee.Context) {
		c.JSON(200, gee.H{
			"name": "wangzeyu",
		})
	})
}
func BenchmarkTestGin(b *testing.B) {
	router := gin.Default()
	router.GET("/q1", func(c *gin.Context) {
		c.JSON(200, gee.H{
			"name": "wangzeyu",
		})
	})

}
