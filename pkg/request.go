package pkg

import "encoding/json"

// Request interface describes all necessary methods
// For a fizzbuzz request.
type Request interface {
	JSON() ([]byte, error)
	Key() ([]byte, error)
	FromKey(key []byte) (Request, error)
	FizzNumber() int
	FizzString() string
	BuzzNumber() int
	BuzzString() string
	Limit() int
}

type request struct {
	FizzNum int    `json:"int1"`
	BuzzNum int    `json:"int2"`
	FizzStr string `json:"str1"`
	BuzzStr string `json:"str2"`
	Lim     int    `json:"limit"`
}

// Key is an alias for the JSON method
func (r request) Key() ([]byte, error) {
	return r.JSON()
}

// JSON will return a json representation of a request
// Never returns an error since ints and strings can't
// cause json.Marshal to return an error
func (r request) JSON() ([]byte, error) {
	b, _ := json.Marshal(r)
	return b, nil
}

// RequestFromKey will return that key back into a request
func (request) FromKey(key []byte) (Request, error) {
	var req request
	if err := json.Unmarshal(key, &req); err != nil {
		return nil, err
	}
	return req, nil
}

func (r request) FizzString() string {
	return r.FizzStr
}

func (r request) FizzNumber() int {
	return r.FizzNum
}

func (r request) BuzzString() string {
	return r.BuzzStr
}

func (r request) BuzzNumber() int {
	return r.BuzzNum
}

func (r request) Limit() int {
	return r.Lim
}
