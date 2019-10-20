package operation

import "io"

const (
	OpMnemonicCLI = "CLI"

	OpCodeCLI = 0x58
)

func IsOpCodeValidCLI(opCode uint8) bool {
	return opCode == OpCodeCLI
}

func IsMnemonicValidCLI(mnemonic string) bool {
	return mnemonic == OpMnemonicCLI
}

type CLI struct {
	baseOperation
}

// 0x58: CLI
func NewCLI() *CLI {
	return &CLI{
		baseOperation{
			code:        OpCodeCLI,
			addressMode: AddrModeImplied,
			mnemonic:    OpMnemonicCLI,
		},
	}
}

func NewCLIBinary(opCode uint8, data io.Reader) (*CLI, error) {
	switch opCode {
	case OpCodeCLI:
		return NewCLI(), nil
	default:
		return nil, InvalidOpCodeError{
			OpCode: opCode,
		}
	}
}

func NewCLIFromBytes(addrMode uint8, arg0 uint8, arg1 uint8) (*CLI, error) {
	switch addrMode {
	case AddrModeImplied:
		return NewCLI(), nil
	default:
		return nil, InvalidAddressModeError{
			AddressMode: addrMode,
		}
	}
}
