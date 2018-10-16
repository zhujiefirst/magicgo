package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/post", func(c *gin.Context) {
		id := c.Query("id")
		page := c.DefaultQuery("page", "0")
		name := c.PostForm("name")
		msg := c.PostForm("msg")

		fmt.Printf("id: %s; page: %s; name: %s; message: %s\n", id, page, name, msg)
	})

	r.Run(":8081")
}
