package fizzbuzz

import (
	"testing"
)

func Test_ShouldReturnFizz_GivenThreeValue(t *testing.T) {
	actual := FizzBuzz(1, 3)
	expected := "1 2 fizz"

	if actual != expected {
		t.Fatalf("expected %s, but got %s", expected, actual)
	}
}

func Test_ShouldReturnBuzz_GivenFiveValue(t *testing.T) {
	actual := FizzBuzz(3, 5)
	expected := "fizz 4 buzz"

	if actual != expected {
		t.Fatalf("expected %s, but got %s", expected, actual)
	}
}

func Test_ShouldReturnFizzBuzz_GivenMultipleOfThreeAndFive(t *testing.T) {
	actual := FizzBuzz(14, 15)
	expected := "14 fizzbuzz"

	if actual != expected {
		t.Fatalf("expected %s, but got %s", expected, actual)
	}
}
