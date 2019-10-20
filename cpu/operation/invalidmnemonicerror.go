package operation

import "fmt"

type InvalidMnemonicError struct {
	Mnemonic string
}

func (err InvalidMnemonicError) Error() string {
	return fmt.Sprintf("invalid mnemonic: %v", err.Mnemonic)
}
