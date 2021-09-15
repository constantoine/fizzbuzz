// package server contains a webservice using core fizzbuzz library
package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/constantoine/fizzbuzz/internal"
)

func main() {
	var port int
	flag.IntVar(&port, "port", 8080, "define the port the server will listen to")
	flag.Parse()

	http.HandleFunc("/fizz", internal.RouteFizzBuzz)
	http.HandleFunc("/stats", internal.RouteStats)
	srv := &http.Server{
		Addr:              fmt.Sprintf(":%d", port),
		ReadTimeout:       1 * time.Second,
		WriteTimeout:      5 * time.Second,
		IdleTimeout:       30 * time.Second,
		ReadHeaderTimeout: 2 * time.Second,
	}
	srv.ListenAndServe()
}
