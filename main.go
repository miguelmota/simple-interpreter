package main

import (
	"fmt"

	"github.com/miguelmota/simple-interpreter/interpreter"
)

func run(text string, expected int64) {
	preter := interpreter.NewInterpreter(text)
	result := preter.Expr()
	fmt.Println("->", result, result == expected)
}

func main() {
	run("2+3", 5)
	run("2 + 3", 5)
	run("2 +3", 5)
	run("5-2", 3)
	run("12+3", 15)
	run("32 -3", 29)
	run("2 + 2 + 4", 8)
	run("2 + 2 + 14", 18)
	run("7 * 4 / 2", 14)
	run("7 * 4 / 2 * 3", 42)
	run("10 * 4  * 2 * 3 / 8", 30)
	run("10 * 4 + 1", 41)
	run("1 + 2 * 4 + 1", 10)
	run("14 + 2 * 3 - 6 / 2", 17)
}
