package main

import (
	"fmt"

	"git.sk.sk/sk/calc/internal/calc"
)

func main() {
	// Good
	fmt.Println(calc.Calculate("X * X"))
	fmt.Println(calc.Calculate("1 + 2"))
	fmt.Println(calc.Calculate("VI / III"))
	fmt.Println(calc.Calculate("VII / III"))
	fmt.Println(calc.Calculate("I + II"))

	// Norman
	fmt.Println(calc.Calculate("I - II"))

	//Bad
	fmt.Println(calc.Calculate("I + 1"))
	fmt.Println(calc.Calculate("I"))
	fmt.Println(calc.Calculate("1 + 1 + 1"))
	fmt.Println(calc.Calculate("3 & 5"))
	fmt.Println(calc.Calculate("3 / 55"))
}
