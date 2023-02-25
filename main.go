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
	var body string
	avg := average()

	if avg == 0 {
		body = fmt.Sprintf(HTML, "no all user's weights")
	} else {
		body = fmt.Sprintf(HTML, avg)
	}

	c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(body))
}

func updateEndPoint(c *gin.Context) {
	name := c.PostForm("name")
	password := c.PostForm("password")

	weightStr := c.PostForm("weight")
	weight, _ := strconv.Atoi(weightStr)

	id, err := auth(name, password)
	if err != nil {
		c.String(http.StatusUnauthorized, "password not valid")
		return
	}

	update(id, weight)
	baseEndPoint(c)
}

func main() {
	r := gin.Default()

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
