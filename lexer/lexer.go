package lexer

import "monkey-lang/token"

type Lexer struct {
}

func NewLexer(s string) *Lexer {
	return &Lexer{}
}

func (l *Lexer) NextToken() *token.Token {
	return &token.Token{}
}
