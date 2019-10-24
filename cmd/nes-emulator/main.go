package main

import (
	"fmt"
	"os"

	"github.com/cd1/nes-emulator"
)

func main() {
	game, err := emulator.LoadGame(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load the game: %v.\n", err)
		os.Exit(1)
	}

	fmt.Printf("%q\n", game.Header.MagicNumber())
}
