package operation

import "io"

const (
	OpMnemonicNOP = "NOP"

	OpCodeNOP = 0xEA
)

func IsOpCodeValidNOP(opCode uint8) bool {
	return opCode == OpCodeNOP
}

func IsMnemonicValidNOP(mnemonic string) bool {
	return mnemonic == OpMnemonicNOP
}

type NOP struct {
	baseOperation
}

// 0xEA: NOP
func NewNOP() *NOP {
	return &NOP{
		baseOperation{
			code:        OpCodeNOP,
			addressMode: AddrModeImplied,
			mnemonic:    OpMnemonicNOP,
		},
	}
}

func NewNOPBinary(opCode uint8, data io.Reader) (*NOP, error) {
	switch opCode {
	case OpCodeNOP:
		return NewNOP(), nil
	default:
		return nil, InvalidOpCodeError{
			OpCode: opCode,
		}
	}
}

func NewNOPFromBytes(addrMode uint8, arg0 uint8, arg1 uint8) (*NOP, error) {
	switch addrMode {
	case AddrModeImplied:
		return NewNOP(), nil
	default:
		return nil, InvalidAddressModeError{
			AddressMode: addrMode,
		}
	}
}
