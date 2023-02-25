package main

import (
	"crypto/sha256"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type User struct {
	id             int
	name           string
	hashedPassword string
	weight         int
}

var users = []User{
	{
		id:             0,
		name:           "taro",
		hashedPassword: "82de639f7b9e3e6208a31db244ac5d2e0dea9f54c90fa33def6cecf98ddc5b7a",
		weight:         0,
	},
	{
		id:             1,
		name:           "jiro",
		hashedPassword: "17f136cdb2905b019ad2aa5af7a96cd17c2885558d679f50bd6cf991d4da5cb7",
		weight:         0,
	},
	{
		id:             2,
		name:           "saburo",
		hashedPassword: "d3cd0977bd92193b27894ad52eb43fb6977687a62fb6f27b17a45fd14d60d0a1",
		weight:         0,
	},
}

// func print() {
// 	for _, v := range users {
// 		fmt.Printf("name: %v\n", v.name)
// 		fmt.Printf("hashedPassword: %v\n", v.hashedPassword)
// 		fmt.Printf("weight: %v\n\n", v.weight)
// 	}
// }

func average() int {
	result := 0

	for _, v := range users {
		if v.weight == 0 {
			return 0
		}

		result += v.weight
	}

	return result / len(users)
}

func auth(name, password string) (int, error) {
	raw := sha256.Sum256([]byte(password))
	hashedPassword := fmt.Sprintf("%x", raw)

	for _, v := range users {
		if v.name == name && v.hashedPassword == hashedPassword {
			return v.id, nil
		}

	}

	return -1, fmt.Errorf("%v", "Not valid!")
}

func update(id, weight int) {
	users[id].weight = weight
}

func averageEndpoint(c *gin.Context) {
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

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*.html")

	r.GET("/", averageEndpoint)

	r.POST("/", func(c *gin.Context) {
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
		averageEndpoint(c)
	})
	r.Run()
}
