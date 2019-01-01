package interpreter

import (
	"log"
	"regexp"
	"strconv"
)

// Interpreter ...
type Interpreter struct {
	lexer        *Lexer
	CurrentToken *Token
}

// NewInterpreter ...
func NewInterpreter(text string) *Interpreter {
	lexer := NewLexer(text)
	currentToken := lexer.GetNextToken()
	return &Interpreter{
		lexer:        lexer,
		CurrentToken: currentToken,
	}
}

// Eat ...
func (i *Interpreter) Eat(tokenKind string) {
	// compare the current token type with the passed token
	// type and if they match then "eat" the current token
	// and assign the next token to the self.current_token,
	// otherwise raise an exception.
	if i.CurrentToken.Kind == tokenKind {
		i.CurrentToken = i.lexer.GetNextToken()
	} else {
		log.Fatal("unreachable2")
	}
}

// Factor ...
func (i *Interpreter) Factor() interface{} {
	// factor : INTEGER
	token := i.CurrentToken
	i.Eat(Integer)
	return token.Value
}

// Term ...
func (i *Interpreter) Term() interface{} {
	// term : factor ((MUL | DIV) factor)*
	result := i.Factor()

	for i.CurrentToken.Kind == Mul || i.CurrentToken.Kind == Div {
		token := i.CurrentToken
		if token.Kind == Mul {
			i.Eat(Mul)
			result = toInteger(result) * toInteger(i.Factor())
		} else if token.Kind == Div {
			i.Eat(Div)
			result = toInteger(result) / toInteger(i.Factor())
		}
	}

	return result
}

// Expr ...
// expr -> INTEGER PLUS INTEGER
// expr -> INTEGER MINUS INTEGER
func (i *Interpreter) Expr() interface{} {
	/*
		 Arithmetic expression parser / interpreter.

			calc>  14 + 2 * 3 - 6 / 2
			17

			expr   : term ((PLUS | MINUS) term)*
			term   : factor ((MUL | DIV) factor)*
			factor : INTEGER
	*/

	result := i.Term()

	for i.CurrentToken.Kind == Plus || i.CurrentToken.Kind == Minus {
		token := i.CurrentToken
		if token.Kind == Plus {
			i.Eat(Plus)
			result = toInteger(result) + toInteger(i.Term())
		} else if token.Kind == Minus {
			i.Eat(Minus)
			result = toInteger(result) - toInteger(i.Term())
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
