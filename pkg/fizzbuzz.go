package pkg

import (
	"strconv"
)

// NewRequest will create a fizzbuzz request according to given parameters
func NewRequest(fizzNumber, buzzNumber int, fizzString, BuzzString string, limit int) Request {
	return request{
		FizzNum: fizzNumber,
		BuzzNum: buzzNumber,
		FizzStr: fizzString,
		BuzzStr: BuzzString,
		Lim:     limit,
	}
}

// FizzBuzzWithStats is an alias to FizzBuzz
// Will register hits in the global counter
func FizzBuzzWithStats(r Request) ([]string, error) {
	if err := counter.Add(r); err != nil {
		return nil, err
	}
	return FizzBuzz(r), nil
}

// FizzBuzz will generate the requested list of strings
func FizzBuzz(r Request) []string {
	str := make([]string, r.Limit())
	fizzbuz := r.FizzNumber() * r.BuzzNumber()
	// Precompute fizzBuzzStr to avoid appending strings at each fizzbuzz
	// Saves one reallocation and one memmove. Slower will small limit
	// But faster with larger sets
	fizzBuzzStr := r.FizzString() + r.BuzzString()
	for i := 1; i <= r.Limit(); i++ {
		switch {
		case i%fizzbuz == 0:
			str[i-1] = fizzBuzzStr
		case i%r.FizzNumber() == 0:
			str[i-1] = r.FizzString()
		case i%r.BuzzNumber() == 0:
			str[i-1] = r.BuzzString()
		default:
			str[i-1] = strconv.Itoa(i)
		}
	}
	return str
}
