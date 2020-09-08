package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

var (
	serverList = []string{
		"http://127.0.0.1:5001",
		"http://127.0.0.1:5002",
		"http://127.0.0.1:5003",
		"http://127.0.0.1:5004",
		"http://127.0.0.1:5005",
	}
	lastServedIndex = 0
)

func main() {
	http.HandleFunc("/", forwardRequest)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func forwardRequest(res http.ResponseWriter, req *http.Request) {
	url := getServer()
	rProxy := httputil.NewSingleHostReverseProxy(url)
	log.Printf("Routing the request to the URL: %s", url.String())
	rProxy.ServeHTTP(res, req)
}

func getServer() *url.URL {
	nextIndex := (lastServedIndex + 1) % 5
	url, _ := url.Parse(serverList[lastServedIndex])
	lastServedIndex = nextIndex
	return url
}
