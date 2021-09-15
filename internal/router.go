package internal

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/constantoine/fizzbuzz/pkg"
)

// RouteFizzBuzz is a GET route that accepts 5 parameters
// three integers int1, int2 and limit, and two strings str1 and str2
// all multiples of int1 are replaced by str1
// all multiples of int2 are replaced by str2
// all multiples of int1 and int2 are replaced by str1str2
func RouteFizzBuzz(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	fizzNum, err := strconv.Atoi(r.FormValue("int1"))
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	buzzNum, err := strconv.Atoi(r.FormValue("int2"))
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	limit, err := strconv.Atoi(r.FormValue("limit"))
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	if fizzNum < 1 || buzzNum < 1 || limit < 1 {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	fizzStr := r.FormValue("str1")
	buzzStr := r.FormValue("str2")
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

// RouteStats is a GET route that will return the parameters or the most popular request
// Response will look like
//	{
//		"request": {
//				"int1": 3,
//				"int2": 5,
//				"str1": "fizz",
//				"str2": "buzz",
//				"limit": 15
//		},
//		"hits": 2
//	}
func RouteStats(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	res, err := pkg.GetMostRequested()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	enc := json.NewEncoder(w)
	enc.SetIndent("", "\t")
	enc.Encode(res)
}
