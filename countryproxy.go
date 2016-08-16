package main

import (
	"log"
	"net/http"
	"os"

	"github.com/elazarl/goproxy"
)

func main() {
	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = true
	port := getDefaultConfig("PORT", "8080")
	log.Fatal(http.ListenAndServe(":"+port, proxy))
}

func getDefaultConfig(name, fallback string) string {
	if val := os.Getenv(name); val != "" {
		return val
	}
	return fallback
}
