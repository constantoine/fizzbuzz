package internal

import (
	"encoding/json"
	"net/http"

	"github.com/constantoine/fizzbuzz/pkg"
)

func RouteFizzBuzz(w http.ResponseWriter, r *http.Request) {
	res, err := pkg.FizzBuzzWithStats(nil)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	enc := json.NewEncoder(w)
	enc.SetIndent("", "\t")
	enc.Encode(res)
}

func RouteStats(w http.ResponseWriter, r *http.Request) {
	res, err := pkg.GetMostRequested()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	enc := json.NewEncoder(w)
	enc.SetIndent("", "\t")
	enc.Encode(res)
}
