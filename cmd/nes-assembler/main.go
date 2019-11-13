package main

import (
	"fmt"
	"os"

	"github.com/cd1/nes-emulator/parser"
)

func main() {
	if err := parser.Assemble(os.Stdin, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to assemble the code (%v)\n", err)
		os.Exit(1)
	}
}
