package main
import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, gin")
	})
	r.GET("/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(200, "Hello, %s", name)
	})
	v1 := r.Group("/v1")
	v1.GET("/test",func(c *gin.Context) {
		c.JSON(200, gin.H{
			"test":"1",
		})
	})
	r.Run(":8999")
}