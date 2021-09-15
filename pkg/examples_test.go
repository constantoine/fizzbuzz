package pkg_test

import (
	"fmt"

	fizzbuzz "github.com/constantoine/fizzbuzz/pkg"
)

func ExampleFizzBuzz() {
	req := fizzbuzz.NewRequest(3, 5, "fizz", "buzz", 16)
	res := fizzbuzz.FizzBuzz(req)
	fmt.Printf("%#+v", res)
	// Output: []string{"1", "2", "fizz", "4", "buzz", "fizz", "7", "8", "fizz", "buzz", "11", "fizz", "13", "14", "fizzbuzz", "16"}
}
