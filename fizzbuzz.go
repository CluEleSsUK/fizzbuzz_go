package fizzbuzz

import (
	"strconv"
)

func FizzBuzz(start int, end int) string {
	var output string
	for i := start; i <= end; i++ {
		output += ConvertToString(i)
		if i < end {
			output += " "
		}
	}
	return output
}

func ConvertToString(val int) string {
	if val%3 == 0 {
		newVal := "fizz"
		if val%5 == 0 {
			newVal += "buzz"
		}
		return newVal
	} else if val%5 == 0 {
		return "buzz"
	} else {
		return strconv.Itoa(val)
	}
}
