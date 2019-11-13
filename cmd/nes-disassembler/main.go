package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/cd1/nes-emulator/parser"
)

var cfg parser.DisassembleConfig

func init() {
	flag.BoolVar(&cfg.DisplayMemoryAddress, "m", false, "Display the memory address in the beginning of each instruction")
	flag.BoolVar(&cfg.DisplayBytes, "b", false, "Display the instruction bytes")
}

func main() {
	flag.Parse()

	if err := parser.Disassemble(os.Stdin, os.Stdout, cfg); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to disassemble the game file (%v).\n", err)
		os.Exit(1)
	}
}
