package calc

import (
	"errors"
	"fmt"
	"strings"
)

// argX, argY, operation, numSystem, err
func parseInput(s string) (int, int, string, NumeralSystem, error) {
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

	argX, numSystemX, err := parseOperand(args[0])
	if err != nil {
		return 0, 0, "", 0, err
	}

	argY, numSystemY, err := parseOperand(args[2])
	if err != nil {
		return 0, 0, "", 0, err
	}

	if numSystemX != numSystemY {
		return 0, 0, "", 0, errors.New(
			"Operands should consist both of Arabic " +
				"or both of uppercase Roman digits.")
	}

	return argX, argY, operation, numSystemX, nil
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

func parseOperand(s string) (int, NumeralSystem, error) {
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
