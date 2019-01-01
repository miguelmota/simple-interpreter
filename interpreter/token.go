package interpreter

import "fmt"

// Integer ...
var Integer = "INTEGER"

// Plus ...
var Plus = "PLUS"

// Minus ...
var Minus = "MINUS"

// Mul ...
var Mul = "MUL"

// Div ...
var Div = "DIV"

// EOF EOF (end-of-file) token is used to indicate that
// there is no more input left for lexical analysis
var EOF = "EOF"

// Token ...
type Token struct {
	Kind  string
	Value interface{}
}

// NewToken ...
func NewToken(kind string, value interface{}) *Token {
	return &Token{
		Kind:  kind,
		Value: value,
	}
}

// String returns string representation of the class instance
func (t *Token) String() string {
	return fmt.Sprintf("Token(%s, %s)", t.Kind, t.Value)
}
