// package cli contains a standalone CLI client for the core fizzbuzz library
package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/constantoine/fizzbuzz/pkg"
)

func main() {
	if len(os.Args) != 6 {
		panic(`5 arguments are needed. int1, int2, limit, str1, str2
all multiples of int1 are replaced by str1
all multiples of int2 are replaced by str2
all multiples of int1 and int2 are replaced by str1str2`)
	}
	fizzNum, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic(err)
	}
	buzzNum, err := strconv.Atoi(os.Args[2])
	if err != nil {
		panic(err)
	}
	limit, err := strconv.Atoi(os.Args[3])
	if err != nil {
		panic(err)
	}
	fizzStr := os.Args[4]
	buzzStr := os.Args[5]
	req := pkg.NewRequest(fizzNum, buzzNum, fizzStr, buzzStr, limit)
	res := pkg.FizzBuzz(req)
	fmt.Println(strings.Join(res, ","))
}
