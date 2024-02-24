package main

import (
	"monkey-lang/repl"
	"os"
)

func main() {

	repl.Start(os.Stdin, os.Stdout)
}
