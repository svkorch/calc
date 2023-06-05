package main

import (
	"fmt"

	"git.sk.sk/sk/calc/internal/calc"
)

func main() {
	// Good
	f("X * X")
	f("1 + 2")
	f("VI / III")
	f("VII / III")
	f("I + II")
	f("3 + 8")
	f("6 - 9")

	// Norman
	f("I - II")
	f("X - X")
	f("III - IV")

	//Bad
	f("I + 1")
	f("I")
	f("1 + 1 + 1")
	f("3 & 5")
	f("3 / 55")
	f("ZVO / 1")
}

func f(s string) {
	fmt.Println(s + " = " + calc.Calculate(s) + "\n")
}
