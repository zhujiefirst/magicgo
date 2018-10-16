package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	v1 := r.Group("/v1")
	{
		v1.GET("read", func(c *gin.Context) {
			c.String(http.StatusOK, "v1/read")
		})
	}

	v2 := r.Group("/v2")
	{
		v2.GET("read", func(c *gin.Context) {
			c.String(http.StatusOK, "v2/read")
		})
	}

	r.Run()
}
