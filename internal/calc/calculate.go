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

// Data type to specify an integer value and a numeral system of a number
type number struct {
	value  int
	system NumeralSystem
}

// Set of numbers allowed for input
var numbers = make(map[string]number)

// Set of functions to do allowed calculations
var operations = make(map[string]func(int, string, int) int)

// Calculate interprets the input string as an arithmetic expression,
// calculates it and returns the result as a string.
// It supports Roman and Arabic numeral systems.
// Both of input expression operands must be in the same numeral system.
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
		numbers[strconv.Itoa(i)] = number{i, ARABIC_DIGITS}
	}

	// Precalculations of allowed roman input numbers
	for i := MIN_NUMBER; i <= MAX_NUMBER; i++ {
		romanStr, err := rom.IntToString(i)
		if err != nil {
			log.Fatalln(errors.New("Cannot convert integer to roman number string."))
		}
		numbers[romanStr] = number{i, ROMAN_DIGITS}
	}
}
