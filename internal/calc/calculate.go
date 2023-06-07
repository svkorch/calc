package calc

import (
	"errors"
	"log"
	"strconv"

	rom "github.com/brandenc40/romannumeral"
)

// Data type to specify the type of number
type NumeralSystem int

// Constants to specify the type of number
const (
	ARABIC_DIGITS NumeralSystem = iota
	ROMAN_DIGITS
)

const (
	MIN_NUMBER = 1
	MAX_NUMBER = 10
)

// Set of functions to do allowed calculations
var operations = make(map[string]func(int, string, int) int)

// Set of arabic numbers allowed for input
var arabicNumbers = make(map[string]int)

// Set of arabic numbers allowed for input
var romanNumbers = make(map[string]int)

func Calculate(s string) string {
	argX, argY, operation, numSystem, err := parseInput(s)
	if err != nil {
		log.Println(err)
		return "!"
	}

	return intToString(operations[operation](argX, operation, argY), numSystem)
}

func intToString(x int, numSystem NumeralSystem) string {
	if numSystem == ARABIC_DIGITS {
		return intToArabicDigits(x)
	}

	if numSystem == ROMAN_DIGITS {
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
	}

	return romanStr
}

func init() {
	// We can use code generation here, of course:-)
	operations["+"] = func(x int, op string, y int) int { return x + y }
	operations["-"] = func(x int, op string, y int) int { return x - y }
	operations["*"] = func(x int, op string, y int) int { return x * y }
	operations["/"] = func(x int, op string, y int) int { return x / y }

	// Precalculations of allowed arabic input numbers
	for i := MIN_NUMBER; i <= MAX_NUMBER; i++ {
		arabicNumbers[strconv.Itoa(i)] = i
	}

	// Precalculations of allowed roman input numbers
	for i := MIN_NUMBER; i <= MAX_NUMBER; i++ {
		romanStr, err := rom.IntToString(i)
		if err != nil {
			log.Fatalln(errors.New("Cannot convert integer to roman number string."))
		}
		romanNumbers[romanStr] = i
	}
}
