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
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.char = 0
	} else {
		l.char = l.input[l.readPosition]
		l.currentPosition = l.readPosition
		l.readPosition++
	}
}

func (l *Lexer) NextToken() *token.Token {
	l.readChar()
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
		t = &token.Token{
			Type: token.ILLEGAL,
		}
	}
	return t
}
