package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"fmt"
	"math/big"
	"net/http"
	"strconv"
	"time"

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
	tlsConfig := setupTLS()

	r := gin.Default()
	r.LoadHTMLGlob("templates/*.html")

	r.GET("/", baseEndPoint)
	r.POST("/", updateEndPoint)

	server := http.Server{
		Addr:      serverAddr,
		TLSConfig: tlsConfig,
		Handler:   r,
	}

	fmt.Printf("📎 Token now available under https://%s/token\n", serverAddr)
	fmt.Printf("👂 Listening on https://%s/secret for secrets...\n", serverAddr)

	server.Handler = r

	err := server.ListenAndServeTLS("", "")
	fmt.Println(err)
}

func setupTLS() *tls.Config {
	template := &x509.Certificate{
		SerialNumber: &big.Int{},
		Subject:      pkix.Name{CommonName: "localhost"},
		NotAfter:     time.Now().Add(time.Hour),
		DNSNames:     []string{"localhost"},
	}

	priv, _ := rsa.GenerateKey(rand.Reader, 2048)
	cert, _ := x509.CreateCertificate(rand.Reader, template, template, &priv.PublicKey, priv)

	tlsCfg := tls.Config{
		Certificates: []tls.Certificate{
			{
				Certificate: [][]byte{cert},
				PrivateKey:  priv,
			},
		},
	}
	return &tlsCfg
}
