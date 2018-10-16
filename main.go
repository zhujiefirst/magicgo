package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RequestIdMiddleware(c *gin.Context) {
	c.Writer.Header().Set("X-Request-Id", "This is RequestId")
	c.Next()
}

func main() {
	r := gin.Default()

	r.Use(RequestIdMiddleware)

	r.GET("/middleware", func(c *gin.Context) {
		c.String(http.StatusOK, "")
	})

	r.Run()
}
