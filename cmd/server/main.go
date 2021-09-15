package main

import (
	"net/http"

	"github.com/constantoine/fizzbuzz/internal"
)

func main() {
	http.HandleFunc("/fizz", internal.RouteFizzBuzz)
	http.HandleFunc("/stats", internal.RouteStats)
	http.ListenAndServe(":8080", nil)
}
