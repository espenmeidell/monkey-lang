package repl

import (
	"bufio"
	"fmt"
	"io"
	lexer "monkey-lang/lexer"
	"monkey-lang/token"
)

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprint(out, ">> ")
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		l := lexer.NewLexer(scanner.Text())
		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Fprintln(out, tok)
		}
	}
}
