package internal

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/constantoine/fizzbuzz/pkg"
)

func writeErr(w http.ResponseWriter, status int, text string) {
	http.Error(w, text, status)
	log.Printf("Error %d: %s", status, text)
}

// RouteFizzBuzz is a GET route that accepts 5 parameters
// three strictly positive integers int1, int2 and limit, and two strings str1 and str2
// all multiples of int1 are replaced by str1
// all multiples of int2 are replaced by str2
// all multiples of int1 and int2 are replaced by str1str2
func RouteFizzBuzz(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeErr(w, http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed))
		return
	}
	fizzNum, err := strconv.Atoi(r.FormValue("int1"))
	if err != nil {
		writeErr(w, http.StatusBadRequest, fmt.Sprintf("int1 (%s) could not be formatted into an int", r.FormValue("int1")))
		return
	}
	buzzNum, err := strconv.Atoi(r.FormValue("int2"))
	if err != nil {
		writeErr(w, http.StatusBadRequest, fmt.Sprintf("int2 (%s) could not be formatted into an int", r.FormValue("int2")))
		return
	}
	limit, err := strconv.Atoi(r.FormValue("limit"))
	if err != nil {
		writeErr(w, http.StatusBadRequest, fmt.Sprintf("limit (%s) could not be formatted into an int", r.FormValue("limit")))
		return
	}
	if fizzNum < 1 {
		writeErr(w, http.StatusBadRequest, fmt.Sprintf("int1 (%d) must be strictly positive", fizzNum))
		return
	} else if buzzNum < 1 {
		writeErr(w, http.StatusBadRequest, fmt.Sprintf("int2 (%d) must be strictly positive", buzzNum))
		return
	} else if limit < 1 {
		writeErr(w, http.StatusBadRequest, fmt.Sprintf("limit (%d) must be strictly positive", limit))
		return
	}
	fizzStr := r.FormValue("str1")
	buzzStr := r.FormValue("str2")
	req := pkg.NewRequest(fizzNum, buzzNum, fizzStr, buzzStr, limit)
	res, err := pkg.FizzBuzzWithStats(req)
	if err != nil {
		log.Printf("Error registering stats: %s", err.Error())
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
		log.Printf("Error retrieving stats: %s", err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	enc := json.NewEncoder(w)
	enc.SetIndent("", "\t")
	enc.Encode(res)
}
