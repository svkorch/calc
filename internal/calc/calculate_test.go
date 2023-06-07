package calc

import (
	"fmt"
	"strconv"
	"testing"

	rom "github.com/brandenc40/romannumeral"
)

// Test section

var testCaseGood = map[string]string{
	"X * X":     "C",
	"1 + 2":     "3",
	"VI / III":  "II",
	"VII / III": "II",
	"I + II":    "III",
	"3 + 8":     "11",
	"6 - 9":     "-3",
}

var testCaseNormal = []string{
	"I - II",
	"X - X",
	"III - IV",
}

var testCaseBad = []string{
	"I + 1",
	"I",
	"1 + 1 + 1",
	"3 & 5",
	"3 / 55",
	"ZVO / 1",
}

func TestGoodCases(t *testing.T) {
	for input, expected := range testCaseGood {
		got := Calculate(input)

		if got != expected {
			t.Fatalf("Input: [%s], got [%s], expected [%s]\n", input, got, expected)
		}
	}
}

func TestNormalCases(t *testing.T) {
	expected := ""
	for _, input := range testCaseNormal {
		got := Calculate(input)

		if got != expected {
			t.Fatalf("Input: [%s], got [%s], expected [%s]\n", input, got, expected)
		}
	}
}

func TestBadCases(t *testing.T) {
	expected := "!"
	for _, input := range testCaseBad {
		got := Calculate(input)

		if got != expected {
			t.Fatalf("Input: [%s], got [%s], expected [%s]\n", input, got, expected)
		}
	}
}

// Benchmark section

var benchNumbers = []struct {
	arabic string
	roman  string
}{
	{"1", "I"},
	{"2", "II"},
	{"3", "III"},
	{"4", "IV"},
	{"5", "V"},
	{"6", "VI"},
	{"7", "VII"},
	{"8", "VIII"},
	{"9", "IX"},
	{"10", "X"},
}

func Benchmark_libRomanStrToInt(b *testing.B) {
	for _, v := range benchNumbers {
		b.Run(fmt.Sprintf("input_[%s]", v.roman), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				rom.StringToInt(v.roman)
			}
		})
	}
}

func Benchmark_libArabicStrToInt(b *testing.B) {
	for _, v := range benchNumbers {
		b.Run(fmt.Sprintf("input_[%s]", v.arabic), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				strconv.Atoi(v.arabic)
			}
		})
	}
}

func Benchmark_mapRomanStrToInt(b *testing.B) {
	for _, v := range benchNumbers {
		b.Run(fmt.Sprintf("input_[%s]", v.roman), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = romanNumbers[v.roman]
			}
		})
	}
}

func Benchmark_mapArabicStrToInt(b *testing.B) {
	for _, v := range benchNumbers {
		b.Run(fmt.Sprintf("input_[%s]", v.arabic), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = arabicNumbers[v.arabic]
			}
		})
	}
}
