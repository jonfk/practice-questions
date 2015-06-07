package main

import (
	"fmt"
)

var ones []string = []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten"}
var teens []string = []string{"zeroteen", "eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen"}
var tens []string = []string{"zeroten", "oneten", "twenty", "thirty", "fourty", "fifty", "sixty", "seventy", "eighty", "ninety"}

func convert(a int) string {
	var millions, thousands, hundreds, tensS, teensS string
	if a >= 1000000 {
		millions = convert(a / 1000000)
		a = a % 1000000
	}
	if a >= 1000 {
		thousands = convert(a / 1000)
		a = a % 1000
	}
	if a >= 100 {
		hundreds = convert(a / 100)
		a = a % 100
	}
	if a >= 20 {
		tensS = tens[a/10]
		a = a % 10
	}

	if a >= 11 {
		a = a % 10
		teensS = teens[a]
	}
	fmt.Printf("%d\n", a)
	result := ""
	if millions != "" {
		result = result + millions + " million,"
	}
	if thousands != "" {
		result = result + thousands + " thousands,"
	}
	if hundreds != "" {
		result = result + hundreds + " hundred,"
	}
	if tensS != "" {
		result = result + tensS + " "
	}
	if teensS != "" {
		result = result + teensS + " "
	}
	if a != 0 && teensS == "" {
		result = result + ones[a]
	}
	return result
}

func main() {
	fmt.Printf("%s\n", convert(12555))
}
