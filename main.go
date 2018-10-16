package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.POST("/form_post", func(c *gin.Context) {
		msg := c.PostForm("msg")
		nick := c.DefaultPostForm("nick", "anonymous")

		c.JSON(http.StatusOK, gin.H{
			"status": "posted",
			"msg": msg,
			"nick": nick,
		})
	})

	r.Run()
}
