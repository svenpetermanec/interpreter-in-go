package main

import (
	"os"

	"interpreter/repl"
)

func main() {
	repl.Start(os.Stdin, os.Stdout)
}
