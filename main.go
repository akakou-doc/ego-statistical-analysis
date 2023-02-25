package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// func print() {
// 	for _, v := range users {
// 		fmt.Printf("name: %v\n", v.name)
// 		fmt.Printf("hashedPassword: %v\n", v.hashedPassword)
// 		fmt.Printf("weight: %v\n\n", v.weight)
// 	}
// }

func baseEndPoint(c *gin.Context) {
	avg := average()
	msg := ""

	if avg == 0 {
		msg = "no all user's weights"
	} else {
		msg = fmt.Sprintf("%v", avg)
	}

	c.HTML(http.StatusOK, "index.html", gin.H{
		"average": msg,
	})
}

func updateEndPoint(c *gin.Context) {
	name := c.PostForm("name")
	password := c.PostForm("password")

	weightStr := c.PostForm("weight")
	weight, _ := strconv.Atoi(weightStr)

	id, err := auth(name, password)
	if err != nil {
		c.String(400, "password not valid")
		return

	}

	update(id, weight)
	baseEndPoint(c)
}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*.html")

	r.GET("/", baseEndPoint)
	r.POST("/", updateEndPoint)

	r.Run()
}
