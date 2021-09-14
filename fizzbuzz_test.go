package fizzbuzz_test

import (
	"math/rand"
	"testing"

	fizzbuzz "github.com/constantoine/fizzbuzz/pkg"
)

func TestFizzBuzz(t *testing.T) {
	req := fizzbuzz.NewRequest(3, 5, "fizz", "buzz", 16)
	res := fizzbuzz.FizzBuzz(req)
	answer := "1,2,fizz,4,buzz,fizz,7,8,fizz,buzz,11,fizz,13,14,fizzbuzz,16"
	if res != answer {
		t.Logf("Output mismatch\nExpected:\n%s\nGot:\n%s\n", answer, res)
		t.Fail()
	}
}

func TestFizzBuzzWithStats(t *testing.T) {
	req := fizzbuzz.NewRequest(3, 5, "fizz", "buzz", 16)
	res, err := fizzbuzz.FizzBuzzWithStats(req)
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	answer := "1,2,fizz,4,buzz,fizz,7,8,fizz,buzz,11,fizz,13,14,fizzbuzz,16"
	if res != answer {
		t.Logf("Output mismatch\nExpected:\n%s\nGot:\n%s\n", answer, res)
		t.Fail()
	}
}

func TestFizzBuzzDuplicate(t *testing.T) {
	req := fizzbuzz.NewRequest(3, 5, "fizz", "fizz", 16)
	res := fizzbuzz.FizzBuzz(req)
	answer := "1,2,fizz,4,fizz,fizz,7,8,fizz,fizz,11,fizz,13,14,fizzfizz,16"
	if res != answer {
		t.Logf("Output mismatch\nExpected:\n%s\nGot:\n%s\n", answer, res)
		t.Fail()
	}
}

func TestFizzBuzzDuplicateWithStats(t *testing.T) {
	req := fizzbuzz.NewRequest(3, 5, "fizz", "fizz", 16)
	res, err := fizzbuzz.FizzBuzzWithStats(req)
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	answer := "1,2,fizz,4,fizz,fizz,7,8,fizz,fizz,11,fizz,13,14,fizzfizz,16"
	if res != answer {
		t.Logf("Output mismatch\nExpected:\n%s\nGot:\n%s\n", answer, res)
		t.Fail()
	}
}

func BenchmarkFizzBuzz16l(b *testing.B) {
	req := fizzbuzz.NewRequest(3, 5, "fizz", "fizz", 16)
	for n := 0; n < b.N; n++ {
		fizzbuzz.FizzBuzz(req)
	}
}

func BenchmarkFizzBuzz100(b *testing.B) {
	req := fizzbuzz.NewRequest(3, 5, "fizz", "fizz", 100)
	for n := 0; n < b.N; n++ {
		fizzbuzz.FizzBuzz(req)
	}
}

func BenchmarkFizzBuzz1k(b *testing.B) {
	req := fizzbuzz.NewRequest(3, 5, "fizz", "fizz", 1000)
	for n := 0; n < b.N; n++ {
		fizzbuzz.FizzBuzz(req)
	}
}

func BenchmarkFizzBuzz10k(b *testing.B) {
	req := fizzbuzz.NewRequest(3, 5, "fizz", "fizz", 10000)
	for n := 0; n < b.N; n++ {
		fizzbuzz.FizzBuzz(req)
	}
}

func BenchmarkFizzBuzz16WithStats(b *testing.B) {
	req := fizzbuzz.NewRequest(3, 5, "fizz", "fizz", 16)
	for n := 0; n < b.N; n++ {
		fizzbuzz.FizzBuzzWithStats(req)
	}
}

func BenchmarkFizzBuzz100WithStats(b *testing.B) {
	req := fizzbuzz.NewRequest(3, 5, "fizz", "fizz", 100)
	for n := 0; n < b.N; n++ {
		fizzbuzz.FizzBuzzWithStats(req)
	}
}

func BenchmarkFizzBuzz1kWithStats(b *testing.B) {
	req := fizzbuzz.NewRequest(3, 5, "fizz", "fizz", 1000)
	for n := 0; n < b.N; n++ {
		fizzbuzz.FizzBuzzWithStats(req)
	}
}

func BenchmarkFizzBuzz10kWithStats(b *testing.B) {
	req := fizzbuzz.NewRequest(3, 5, "fizz", "fizz", 10000)
	for n := 0; n < b.N; n++ {
		fizzbuzz.FizzBuzzWithStats(req)
	}
}

func BenchmarkFizzBuzzGetMostRequested(b *testing.B) {
	for i := 0; i < 10000; i++ {
		req := fizzbuzz.NewRequest(rand.Intn(10)+1, rand.Intn(10)+1, "fizz", "buzz", 10000)
		_, err := fizzbuzz.FizzBuzzWithStats(req)
		if err != nil {
			b.Log(err)
			b.Fail()
		}
	}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_, _, err := fizzbuzz.GetMostRequested()
		if err != nil {
			b.Log(err)
			b.Fail()
		}
	}
}
