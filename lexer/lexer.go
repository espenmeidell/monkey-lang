package lexer

import (
	"monkey-lang/token"
)

type Lexer struct {
	input           string // The source code to process
	currentPosition int    // The current character
	readPosition    int    // The next character to examine
	char            byte   // The current character we are working on
}

func NewLexer(s string) *Lexer {
	l := &Lexer{
		input: s,
	}
	l.readChar()
	return l
}

// readChar updates the current char of the lexer and
// increments the currentPosition and readPosition. When end
// is reached the value of char is 0
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.char = 0
	} else {
		l.char = l.input[l.readPosition]
	}
	l.currentPosition = l.readPosition
	l.readPosition++
}

// peekChar will look at the next character without incrementing current char
func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

// readWord will call readChar as long as the current character is a letter
func (l *Lexer) readWord() string {
	var bytes []byte
	for isLetter(l.char) {
		bytes = append(bytes, l.char)
		l.readChar()
	}
	return string(bytes)
}

// readNumber will call readChar as long as the current character is a digit
func (l *Lexer) readNumber() string {
	var bytes []byte
	for isDigit(l.char) {
		bytes = append(bytes, l.char)
		l.readChar()
	}
	return string(bytes)
}

// skipWhitespace will call readChar as long as the current char is whitespace
func (l *Lexer) skipWhitespace() {
	for l.char == ' ' || l.char == '\t' || l.char == '\r' || l.char == '\n' {
		l.readChar()
	}
}

// NextToken scans the input string for the next token
//
// First match all single char types such as =
// If not one of those, read the entire word and figure out what it is
func (l *Lexer) NextToken() *token.Token {
	l.skipWhitespace()
	var t *token.Token
	switch l.char {
	case '=':
		switch l.peekChar() {
		case '=':
			t = &token.Token{
				Type:    token.EQUALS,
				Literal: token.EQUALS,
			}
			l.readChar()
		default:
			t = &token.Token{
				Type:    token.ASSIGN,
				Literal: token.ASSIGN,
			}
		}
	case '+':
		t = &token.Token{
			Type:    token.PLUS,
			Literal: token.PLUS,
		}
	case '-':
		t = &token.Token{
			Type:    token.MINUS,
			Literal: token.MINUS,
		}
	case '*':
		t = &token.Token{
			Type:    token.MULTIPLY,
			Literal: token.MULTIPLY,
		}
	case '/':
		t = &token.Token{
			Type:    token.DIVIDE,
			Literal: token.DIVIDE,
		}
	case '!':
		switch l.peekChar() {
		case '=':
			t = &token.Token{
				Type:    token.NOT_EQUALS,
				Literal: token.NOT_EQUALS,
			}
			l.readChar()
		default:
			t = &token.Token{
				Type:    token.NOT,
				Literal: token.NOT,
			}
		}
	case '<':
		t = &token.Token{
			Type:    token.LESS_THAN,
			Literal: token.LESS_THAN,
		}
	case '>':
		t = &token.Token{
			Type:    token.GREATER_THAN,
			Literal: token.GREATER_THAN,
		}
	case '(':
		t = &token.Token{
			Type:    token.LPAREN,
			Literal: token.LPAREN,
		}
	case '{':
		t = &token.Token{
			Type:    token.LBRACE,
			Literal: token.LBRACE,
		}
	case ')':
		t = &token.Token{
			Type:    token.RPAREN,
			Literal: token.RPAREN,
		}
	case '}':
		t = &token.Token{
			Type:    token.RBRACE,
			Literal: token.RBRACE,
		}
	case ',':
		t = &token.Token{
			Type:    token.COMMA,
			Literal: token.COMMA,
		}
	case ';':
		t = &token.Token{
			Type:    token.SEMICOLON,
			Literal: token.SEMICOLON,
		}
	case 0:
		t = &token.Token{
			Type:    token.EOF,
			Literal: "",
		}
	default:
		if isLetter(l.char) {
			word := l.readWord()
			t = &token.Token{
				Type:    token.LookupIdent(word),
				Literal: word,
			}
			return t // readWord has incremented char position, return early
		} else if isDigit(l.char) {
			number := l.readNumber()
			t = &token.Token{
				Type:    token.INT,
				Literal: number,
			}
			return t // readNumber has incremented char position, return early
		} else {
			t = &token.Token{
				Type:    token.ILLEGAL,
				Literal: string(l.char),
			}
		}
	}
	l.readChar()
	return t
}

func isLetter(c byte) bool {
	if c >= 'A' && c <= 'Z' {
		return true
	}
	if c >= 'a' && c <= 'z' {
		return true
	}
	return false
}

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}
