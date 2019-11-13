package cpu

import "fmt"

type InvalidAddressModeError struct {
	AddressMode uint8
}

func (err InvalidAddressModeError) Error() string {
	return fmt.Sprintf("invalid address mode: %v", AddressModeString(err.AddressMode))
}

type InvalidMnemonicError struct {
	Mnemonic string
}

func (err InvalidMnemonicError) Error() string {
	return fmt.Sprintf("invalid mnemonic: %v", err.Mnemonic)
}

type InvalidOpCodeError struct {
	OpCode uint8
}

func (err InvalidOpCodeError) Error() string {
	return fmt.Sprintf("invalid op code: $%02X", err.OpCode)
}

type InvalidSyntaxError struct {
	Line string
}

func (err InvalidSyntaxError) Error() string {
	return fmt.Sprintf("invalid syntax in line: %v", err.Line)
}
