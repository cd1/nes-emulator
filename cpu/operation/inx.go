package operation

import "io"

const (
	OpMnemonicINX = "INX"

	OpCodeINX = 0xE8
)

func IsOpCodeValidINX(opCode uint8) bool {
	return opCode == OpCodeINX
}

func IsMnemonicValidINX(mnemonic string) bool {
	return mnemonic == OpMnemonicINX
}

type INX struct {
	baseOperation
}

// 0xE8: INX
func NewINX() *INX {
	return &INX{
		baseOperation{
			code:        OpCodeINX,
			addressMode: AddrModeImplied,
			mnemonic:    OpMnemonicINX,
		},
	}
}

func NewINXBinary(opCode uint8, data io.Reader) (*INX, error) {
	switch opCode {
	case OpCodeINX:
		return NewINX(), nil
	default:
		return nil, InvalidOpCodeError{
			OpCode: opCode,
		}
	}
}

func NewINXFromBytes(addrMode uint8, arg0 uint8, arg1 uint8) (*INX, error) {
	switch addrMode {
	case AddrModeImplied:
		return NewINX(), nil
	default:
		return nil, InvalidAddressModeError{
			AddressMode: addrMode,
		}
	}
}
