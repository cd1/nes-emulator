package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/cd1/nes-emulator"
)

var verbose bool

func init() {
	flag.BoolVar(&verbose, "v", false, "Display information when executing each instruction")
}

func main() {
	flag.Parse()

	game, err := nes.LoadGame(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load the game: %v.\n", err)
		os.Exit(1)
	}

	system := nes.NES{
		Verbose: verbose,
	}

	if err := system.Run(*game); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to run the game: %v.\n", err)
		os.Exit(1)
	}
}
