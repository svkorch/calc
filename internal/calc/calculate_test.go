package calc

import "testing"

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
		got := calculate(input)

		if got != expected {
			t.Fatalf("Input: [%s], got [%s], expected [%s]\n", input, got, expected)
		}
	}
}

func TestNormalCases(t *testing.T) {
	expected := ""
	for _, input := range testCaseNormal {
		got := calculate(input)

		if got != expected {
			t.Fatalf("Input: [%s], got [%s], expected [%s]\n", input, got, expected)
		}
	}
}

func TestBadCases(t *testing.T) {
	expected := "!"
	for _, input := range testCaseBad {
		got := calculate(input)

		if got != expected {
			t.Fatalf("Input: [%s], got [%s], expected [%s]\n", input, got, expected)
		}
	}
}
