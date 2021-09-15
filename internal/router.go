package internal

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/constantoine/fizzbuzz/pkg"
)

func RouteFizzBuzz(w http.ResponseWriter, r *http.Request) {
	fizzNum, err := strconv.Atoi(r.Form.Get("int1"))
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	buzzNum, err := strconv.Atoi(r.Form.Get("int2"))
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	limit, err := strconv.Atoi(r.Form.Get("limit"))
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	fizzStr := r.Form.Get("str1")
	buzzStr := r.Form.Get("str2")
	req := pkg.NewRequest(fizzNum, buzzNum, fizzStr, buzzStr, limit)
	res, err := pkg.FizzBuzzWithStats(req)
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
