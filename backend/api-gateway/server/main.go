package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	userServiceURL := os.Getenv("USER_SERVICE_URL")
	if userServiceURL == "" {
		log.Fatal("Not get userServiceURL env variable")
	}

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		log.Fatal("Not get SERVER_PORT env variable")
	}

	r := gin.Default()

	remote, err := url.Parse(userServiceURL)
	if err != nil {
		log.Fatal("Invalid user service URL: ", err)
	}

	proxy := httputil.NewSingleHostReverseProxy(remote)

	r.NoRoute(func(c *gin.Context) {
		proxy.ServeHTTP(c.Writer, c.Request)
	})

	log.Printf("API listening on port %s, proxying to %s", port, userServiceURL)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}
