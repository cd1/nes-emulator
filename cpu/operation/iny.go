package operation

import "io"

const (
	OpMnemonicINY = "INY"

	OpCodeINY = 0xC8
)

func IsOpCodeValidINY(opCode uint8) bool {
	return opCode == OpCodeINY
}

func IsMnemonicValidINY(mnemonic string) bool {
	return mnemonic == OpMnemonicINY
}

type INY struct {
	baseOperation
}

// 0xC8: INY
func NewINY() *INY {
	return &INY{
		baseOperation{
			code:        OpCodeINY,
			addressMode: AddrModeImplied,
			mnemonic:    OpMnemonicINY,
		},
	}
}

func NewINYBinary(opCode uint8, data io.Reader) (*INY, error) {
	switch opCode {
	case OpCodeINY:
		return NewINY(), nil
	default:
		return nil, InvalidOpCodeError{
			OpCode: opCode,
		}
	}
}

func NewINYFromBytes(addrMode uint8, arg0 uint8, arg1 uint8) (*INY, error) {
	switch addrMode {
	case AddrModeImplied:
		return NewINY(), nil
	default:
		return nil, InvalidAddressModeError{
			AddressMode: addrMode,
		}
	}
}
