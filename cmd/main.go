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

	// Norman
	f("I - II")

	//Bad
	f("I + 1")
	f("I")
	f("1 + 1 + 1")
	f("3 & 5")
	f("3 / 55")
}

func f(s string) {
	fmt.Println(s + " = " + calc.Calculate(s))
}
