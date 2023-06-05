package calc

import (
	"errors"
	"fmt"
	"log"
	"strconv"
)

type digits int

const (
	_ = iota
	ARABIC_DIGITS
	ROMAN_DIGITS
)

const (
	MIN_NUMBER = 1
	MAX_NUMBER = 10
)

var operations = make(map[string]func(int, string, int) int)
var arabicNumbers = make(map[string]int)
var romanNumbers = make(map[string]int)

func calculate(s string) string {
	argX, argY, operation, calcSystem, err := parseInput(s)
	if err != nil {
		log.Println(err)
		return ""
	}

	return intToString(operations[operation](argX, operation, argY), calcSystem)
}

func intToString(x int, calcSystem digits) string {
	if calcSystem == ARABIC_DIGITS {
		return intToArabicDigits(x)
	}

	if calcSystem == ROMAN_DIGITS {
		return intToRomanDigits(x)
	}

	return ""
}

func intToArabicDigits(x int) string {
	return fmt.Sprintf("<%d>", x)
}

func intToRomanDigits(x int) string {
	if x <= 0 {
		log.Println(errors.New("Roman system cannot represent a non positive number."))
		return "<>"
	}

	return fmt.Sprintf("Roma<%d>", x)
}

func init() {
	// We can use code generation here, of course:-)
	operations["+"] = func(x int, op string, y int) int { return x + y }
	operations["-"] = func(x int, op string, y int) int { return x - y }
	operations["*"] = func(x int, op string, y int) int { return x * y }
	operations["/"] = func(x int, op string, y int) int { return x / y }

	for i := MIN_NUMBER; i <= MAX_NUMBER; i++ {
		arabicNumbers[strconv.Itoa(i)] = i
	}

	romanNumbers["I"] = 1
	romanNumbers["II"] = 2
	romanNumbers["III"] = 3
	romanNumbers["IV"] = 4
	romanNumbers["V"] = 5
	romanNumbers["VI"] = 6
	romanNumbers["VII"] = 7
	romanNumbers["VIII"] = 8
	romanNumbers["IX"] = 9
	romanNumbers["X"] = 10
}
