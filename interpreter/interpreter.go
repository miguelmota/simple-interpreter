package interpreter

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
)

// Integer ...
var Integer = "INTEGER"

// Plus ...
var Plus = "PLUS"

// Minus ...
var Minus = "MINUS"

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

// Interpreter ...
type Interpreter struct {
	Text         string
	Pos          int
	CurrentToken *Token
	CurrentChar  string
}

// NewInterpreter ...
func NewInterpreter(text string) *Interpreter {
	currentChar := string(text[0])

	return &Interpreter{
		Text:         text,
		Pos:          0,
		CurrentToken: nil,
		CurrentChar:  currentChar,
	}
}

// Advance ...
func (i *Interpreter) Advance() {
	// Advance the 'pos' pointer and set the 'current_char' variable
	i.Pos++
	if i.Pos > len(i.Text)-1 {
		i.CurrentChar = "" // end of input
	} else {
		i.CurrentChar = string(i.Text[i.Pos])
	}
}

// SkipWhitespace ...
func (i *Interpreter) SkipWhitespace() {
	for i.CurrentChar != "" && isSpace(i.CurrentChar) {
		i.Advance()
	}
}

// Integer ...
func (i *Interpreter) Integer() int64 {
	// Return a (multidigit) integer consumed from the input
	var result string
	for i.CurrentChar != "" && isDigit(i.CurrentChar) {
		result += string(i.CurrentChar)
		i.Advance()
	}
	return toInteger(result)
}

// GetNextToken Lexical analyzer (also known as scanner or tokenizer)
// This method is responsible for breaking a sentence
// apart into tokens. One token at a time.
func (i *Interpreter) GetNextToken() *Token {

	for i.CurrentChar != "" {
		if isSpace(i.CurrentChar) {
			i.SkipWhitespace()
			continue
		}

		if isDigit(i.CurrentChar) {
			return NewToken(Integer, i.Integer())
		}

		if i.CurrentChar == "+" {
			i.Advance()
			return NewToken(Plus, "+")
		}

		if i.CurrentChar == "-" {
			i.Advance()
			return NewToken(Minus, "-")
		}

		log.Fatal("unreachable1")
	}

	return NewToken(EOF, nil)
}

// Eat ...
func (i *Interpreter) Eat(tokenKind string) {
	// compare the current token type with the passed token
	// type and if they match then "eat" the current token
	// and assign the next token to the self.current_token,
	// otherwise raise an exception.
	if i.CurrentToken.Kind == tokenKind {
		i.CurrentToken = i.GetNextToken()
	} else {
		log.Fatal("unreachable2")
	}
}

// Term ...
func (i *Interpreter) Term() interface{} {
	// return an INTEGER token value
	token := i.CurrentToken
	i.Eat(Integer)
	return token.Value
}

// Expr ...
// expr -> INTEGER PLUS INTEGER
// expr -> INTEGER MINUS INTEGER
func (i *Interpreter) Expr() interface{} {
	// set current token to the first token taken from the input
	i.CurrentToken = i.GetNextToken()

	result := toInteger(i.Term())
	for i.CurrentToken.Kind == Plus || i.CurrentToken.Kind == Minus {
		token := i.CurrentToken
		if token.Kind == Plus {
			i.Eat(Plus)
			result = result + toInteger(i.Term())
		} else if token.Kind == Minus {
			i.Eat(Minus)
			result = result - toInteger(i.Term())
		}
	}

	return result
}

// toInteger
func toInteger(value interface{}) int64 {
	switch v := value.(type) {
	case int:
		return int64(v)
	case int8:
		return int64(v)
	case int64:
		return int64(v)
	case uint:
		return int64(v)
	case uint8:
		return int64(v)
	case uint64:
		return int64(v)
	case string:
		i, err := strconv.Atoi(v)
		if err != nil {
			log.Fatal(err)
		}
		return int64(i)
	default:
		log.Fatal("unreachable3")
		return 0
	}
}

// isDigit
func isDigit(value interface{}) bool {
	switch v := value.(type) {
	case int:
		return true
	case int8:
		return true
	case int64:
		return true
	case uint:
		return true
	case uint8:
		return true
	case uint64:
		return true
	case string:
		regex := regexp.MustCompile(`^[0-9]+$`)
		return regex.Match([]byte(v))
	default:
		return false
	}
}

// isSpace
func isSpace(value string) bool {
	return value == " "
}
