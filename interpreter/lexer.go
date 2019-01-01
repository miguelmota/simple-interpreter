package interpreter

import "log"

// Lexer ...
type Lexer struct {
	Text        string
	Pos         int
	CurrentChar string
}

// NewLexer ...
func NewLexer(text string) *Lexer {
	currentChar := string(text[0])
	return &Lexer{
		Pos:         0,
		Text:        text,
		CurrentChar: currentChar,
	}
}

// Advance ...
func (i *Lexer) Advance() {
	// Advance the 'pos' pointer and set the 'current_char' variable
	i.Pos++
	if i.Pos > len(i.Text)-1 {
		i.CurrentChar = "" // end of input
	} else {
		i.CurrentChar = string(i.Text[i.Pos])
	}
}

// SkipWhitespace ...
func (i *Lexer) SkipWhitespace() {
	for i.CurrentChar != "" && isSpace(i.CurrentChar) {
		i.Advance()
	}
}

// Integer ...
func (i *Lexer) Integer() int64 {
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
func (i *Lexer) GetNextToken() *Token {
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

		if i.CurrentChar == "*" {
			i.Advance()
			return NewToken(Mul, "*")
		}

		if i.CurrentChar == "/" {
			i.Advance()
			return NewToken(Div, "/")
		}

		log.Fatal("unreachable1")
	}

	return NewToken(EOF, nil)
}
