package main

import (
	"fmt"

	"github.com/miguelmota/migscript/interpreter"
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
}
