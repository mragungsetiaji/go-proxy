package main

import (
	"flag"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	var ipAddress string
	flag.StringVar(&ipAddress, "ip", "8.8.8.8", "IP address of the target server")
	flag.Parse()

	remote, err := url.Parse("http://" + ipAddress)
	if err != nil {
		panic(err)
	}

	handler := func(p *httputil.ReverseProxy) func(http.ResponseWriter, *http.Request) {
		return func(w http.ResponseWriter, r *http.Request) {
			log.Println(r.URL)
			r.Host = remote.Host
			p.ServeHTTP(w, r)
		}
	}

	proxy := httputil.NewSingleHostReverseProxy(remote)
	http.HandleFunc("/", handler(proxy))
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
