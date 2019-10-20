package operation

import "io"

const (
	OpMnemonicDEY = "DEY"

	OpCodeDEY = 0x88
)

func IsOpCodeValidDEY(opCode uint8) bool {
	return opCode == OpCodeDEY
}

func IsMnemonicValidDEY(mnemonic string) bool {
	return mnemonic == OpMnemonicDEY
}

type DEY struct {
	baseOperation
}

// 0x88: DEY
func NewDEY() *DEY {
	return &DEY{
		baseOperation{
			code:        OpCodeDEY,
			addressMode: AddrModeImplied,
			mnemonic:    OpMnemonicDEY,
		},
	}
}

func NewDEYBinary(opCode uint8, data io.Reader) (*DEY, error) {
	switch opCode {
	case OpCodeDEY:
		return NewDEY(), nil
	default:
		return nil, InvalidOpCodeError{
			OpCode: opCode,
		}
	}
}

func NewDEYFromBytes(addrMode uint8, arg0 uint8, arg1 uint8) (*DEY, error) {
	switch addrMode {
	case AddrModeImplied:
		return NewDEY(), nil
	default:
		return nil, InvalidAddressModeError{
			AddressMode: addrMode,
		}
	}
}
