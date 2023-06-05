package calc

import (
	"errors"
	"log"
	"strconv"

	rom "github.com/brandenc40/romannumeral"
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
		log.Fatalln(err)
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
	return strconv.Itoa(x)
}

func intToRomanDigits(x int) string {
	if x <= 0 {
		log.Println(errors.New("Roman system cannot represent a non positive number."))
		return ""
	}

	romanStr, err := rom.IntToString(x)
	if err != nil {
		log.Fatalln(errors.New("Cannot convert integer to roman number string."))
		// return ""
	}

	return romanStr
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

	for i := MIN_NUMBER; i <= MAX_NUMBER; i++ {
		romanStr, err := rom.IntToString(i)
		if err != nil {
			log.Fatalln(errors.New("Cannot convert integer to roman number string."))
		}
		romanNumbers[romanStr] = i
	}
}
