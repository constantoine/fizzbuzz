package pkg

import (
	"strconv"
	"strings"
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
func FizzBuzzWithStats(r Request) (string, error) {
	if err := counter.Add(r); err != nil {
		return "", err
	}
	return FizzBuzz(r), nil
}

// FizzBuzz will generate the requested string
func FizzBuzz(r Request) string {
	str := new(strings.Builder)
	// Start with a big enough allocation to avoid
	// Lots of small rellocations, especially at the beggining of the loop
	str.Grow(r.Limit() * 2)
	fizzbuz := r.FizzNumber() * r.BuzzNumber()
	// Precompute fizzBuzzStr to avoid appending strings at each fizzbuzz
	// Saves one reallocation and one memmove. Slower will small limit
	// But faster with larger sets
	fizzBuzzStr := r.FizzString() + r.BuzzString()
	for i := 1; i <= r.Limit(); i++ {
		switch {
		case i%fizzbuz == 0:
			str.WriteString(fizzBuzzStr)
		case i%r.FizzNumber() == 0:
			str.WriteString(r.FizzString())
		case i%r.BuzzNumber() == 0:
			str.WriteString(r.BuzzString())
		default:
			str.WriteString(strconv.Itoa(i))
		}
		if i < r.Limit() {
			str.WriteRune(',')
		}
	}
	return str.String()
}
