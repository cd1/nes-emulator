package operation

import "io"

const (
	OpMnemonicBRK = "BRK"

	OpCodeBRK = 0x00
)

func IsOpCodeValidBRK(opCode uint8) bool {
	return opCode == OpCodeBRK
}

func IsMnemonicValidBRK(mnemonic string) bool {
	return mnemonic == OpMnemonicBRK
}

type BRK struct {
	baseOperation
}

// 0x00: BRK
func NewBRK() *BRK {
	return &BRK{
		baseOperation{
			code:        OpCodeBRK,
			addressMode: AddrModeImplied,
			mnemonic:    OpMnemonicBRK,
		},
	}
}

func NewBRKBinary(opCode uint8, data io.Reader) (*BRK, error) {
	switch opCode {
	case OpCodeBRK:
		return NewBRK(), nil
	default:
		return nil, InvalidOpCodeError{
			OpCode: opCode,
		}
	}
}

func NewBRKFromBytes(addrMode uint8, arg0 uint8, arg1 uint8) (*BRK, error) {
	switch addrMode {
	case AddrModeImplied:
		return NewBRK(), nil
	default:
		return nil, InvalidAddressModeError{
			AddressMode: addrMode,
		}
	}
}
