package lexer

import (
	"monkey-lang/token"
)

type Lexer struct {
	input           string
	currentPosition int
	readPosition    int
	char            byte
}

func NewLexer(s string) *Lexer {
	l := &Lexer{
		input: s,
	}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.char = 0
	} else {
		l.char = l.input[l.readPosition]
	}
	l.currentPosition = l.readPosition
	l.readPosition++
}

func (l *Lexer) readWord() string {
	var bytes []byte
	for isLetter(l.char) {
		bytes = append(bytes, l.char)
		l.readChar()
	}
	return string(bytes)
}

func (l *Lexer) readNumber() string {
	var bytes []byte
	for isDigit(l.char) {
		bytes = append(bytes, l.char)
		l.readChar()
	}
	return string(bytes)
}

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
		t = &token.Token{
			Type:    token.ASSIGN,
			Literal: token.ASSIGN,
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
		t = &token.Token{
			Type:    token.NOT,
			Literal: token.NOT,
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
			return t
		} else if isDigit(l.char) {
			number := l.readNumber()
			t = &token.Token{
				Type:    token.INT,
				Literal: number,
			}
			return t
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
