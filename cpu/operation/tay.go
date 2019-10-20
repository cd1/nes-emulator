package operation

import "io"

const (
	OpMnemonicTAY = "TAY"

	OpCodeTAY = 0xA8
)

func IsOpCodeValidTAY(opCode uint8) bool {
	return opCode == OpCodeTAY
}

func IsMnemonicValidTAY(mnemonic string) bool {
	return mnemonic == OpMnemonicTAY
}

type TAY struct {
	baseOperation
}

// 0xA8: TAY
func NewTAY() *TAY {
	return &TAY{
		baseOperation{
			code:        OpCodeTAY,
			addressMode: AddrModeImplied,
			mnemonic:    OpMnemonicTAY,
		},
	}
}

func NewTAYBinary(opCode uint8, data io.Reader) (*TAY, error) {
	switch opCode {
	case OpCodeTAY:
		return NewTAY(), nil
	default:
		return nil, InvalidOpCodeError{
			OpCode: opCode,
		}
	}
}

func NewTAYFromBytes(addrMode uint8, arg0 uint8, arg1 uint8) (*TAY, error) {
	switch addrMode {
	case AddrModeImplied:
		return NewTAY(), nil
	default:
		return nil, InvalidAddressModeError{
			AddressMode: addrMode,
		}
	}
}
