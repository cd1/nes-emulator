package operation

import "io"

const (
	OpMnemonicTYA = "TYA"

	OpCodeTYA = 0x98
)

func IsOpCodeValidTYA(opCode uint8) bool {
	return opCode == OpCodeTYA
}

func IsMnemonicValidTYA(mnemonic string) bool {
	return mnemonic == OpMnemonicTYA
}

type TYA struct {
	baseOperation
}

// 0x98: TYA
func NewTYA() *TYA {
	return &TYA{
		baseOperation{
			code:        OpCodeTYA,
			addressMode: AddrModeImplied,
			mnemonic:    OpMnemonicTYA,
		},
	}
}

func NewTYABinary(opCode uint8, data io.Reader) (*TYA, error) {
	switch opCode {
	case OpCodeTYA:
		return NewTYA(), nil
	default:
		return nil, InvalidOpCodeError{
			OpCode: opCode,
		}
	}
}

func NewTYAFromBytes(addrMode uint8, arg0 uint8, arg1 uint8) (*TYA, error) {
	switch addrMode {
	case AddrModeImplied:
		return NewTYA(), nil
	default:
		return nil, InvalidAddressModeError{
			AddressMode: addrMode,
		}
	}
}
