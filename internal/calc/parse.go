package calc

import (
	"errors"
	"fmt"
	"strings"
)

// argX, argY, operation, calcSystem, err
func parseInput(s string) (int, int, string, digits, error) {
	args := strings.Split(s, " ")
	if len(args) != 3 {
		return 0, 0, "", 0, errors.New(
			"Input \"<operand_1> <operation> <operand_2>\" expected.")
	}

	if !checkOperation(args[1]) {
		return 0, 0, "", 0, errors.New(
			"Allowed operations are: " + opList())
	}
	operation := args[1]

	argX, calcSystemX, err := parseOperand(args[0])
	if err != nil {
		return 0, 0, "", 0, err
	}

	argY, calcSystemY, err := parseOperand(args[2])
	if err != nil {
		return 0, 0, "", 0, err
	}

	if calcSystemX != calcSystemY {
		return 0, 0, "", 0, errors.New(
			"Operands should consist both of Arabic " +
				"or both of uppercase Roman digits.")
	}

	return argX, argY, operation, calcSystemX, nil
}

// opList returns a string of comma-space separated
// quoted operations signs
func opList() string {
	list := make([]string, len(operations))

	i := 0
	for op := range operations {
		list[i] = "'" + op + "'"

		i++
	}

	return strings.Join(list, ", ")
}

func checkOperation(s string) bool {
	_, ok := operations[s]

	return ok
}

func parseOperand(s string) (int, digits, error) {
	value, ok := arabicNumbers[s]
	if !ok {
		value, ok = romanNumbers[s]
		if !ok {
			return 0, 0, errors.New(
				fmt.Sprintf(
					"Numbers should be Arabic or uppercase Roman "+
						"from %d to %d.",
					MIN_NUMBER,
					MAX_NUMBER))
		}

		return value, ROMAN_DIGITS, nil
	}

	return value, ARABIC_DIGITS, nil
}
