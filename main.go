package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)



func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*.html")

	r.Any("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"average": "100",
		})
	})
	r.Run()
}
