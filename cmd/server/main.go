// package server contains a webservice using core fizzbuzz library
package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/constantoine/fizzbuzz/internal"
)

func main() {
	var port int
	flag.IntVar(&port, "port", 8080, "define the port the server will listen to")
	flag.Parse()

	http.HandleFunc("/fizz", internal.RouteFizzBuzz)
	http.HandleFunc("/stats", internal.RouteStats)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
