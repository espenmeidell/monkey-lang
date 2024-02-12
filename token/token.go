package token

import "fmt"

type Type string

type Token struct {
	Type    Type
	Literal string
}

func (t Token) String() string {
	return fmt.Sprintf("{type='%s', lit='%s'}", t.Type, t.Literal)
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers and literals
	IDENT = "IDENT"
	INT   = "INT"

	// Operators
	ASSIGN       = "="
	PLUS         = "+"
	MINUS        = "-"
	MULTIPLY     = "*"
	DIVIDE       = "/"
	NOT          = "!"
	LESS_THAN    = "<"
	GREATER_THAN = ">"
	EQUALS       = "=="
	NOT_EQUALS   = "!="

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"

	LBRACE = "{"
	RBRACE = "}"

	// Keywords

	FUNCTION = "FUNCTION"
	LET      = "LET"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
)

var keywords = map[string]Type{
	"fn":     FUNCTION,
	"let":    LET,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
	"true":   TRUE,
	"false":  FALSE,
}

func LookupIdent(ident string) Type {
	if t, ok := keywords[ident]; ok {
		return t
	}
	return IDENT
}
