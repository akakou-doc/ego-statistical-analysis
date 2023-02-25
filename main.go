package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

const attestationProviderURL = "https://shareduks.uks.attest.azure.net"
const serverAddr = "0.0.0.0:8080"

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

	tlsConfig := setupTLS()
	setupAttestaion(r, tlsConfig)

	server := http.Server{
		Addr:      serverAddr,
		TLSConfig: tlsConfig,
		Handler:   r,
	}

	fmt.Printf("📎 Token now available under https://%s/token\n", serverAddr)
	fmt.Printf("👂 Listening on https://%s/...\n", serverAddr)

	err := server.ListenAndServeTLS("", "")
	fmt.Println(err)
}
