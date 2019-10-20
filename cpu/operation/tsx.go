package operation

import "io"

const (
	OpMnemonicTSX = "TSX"

	OpCodeTSX = 0xBA
)

func IsOpCodeValidTSX(opCode uint8) bool {
	return opCode == OpCodeTSX
}

func IsMnemonicValidTSX(mnemonic string) bool {
	return mnemonic == OpMnemonicTSX
}

type TSX struct {
	baseOperation
}

// 0xBA: TSX
func NewTSX() *TSX {
	return &TSX{
		baseOperation{
			code:        OpCodeTSX,
			addressMode: AddrModeImplied,
			mnemonic:    OpMnemonicTSX,
		},
	}
}

func NewTSXBinary(opCode uint8, data io.Reader) (*TSX, error) {
	switch opCode {
	case OpCodeTSX:
		return NewTSX(), nil
	default:
		return nil, InvalidOpCodeError{
			OpCode: opCode,
		}
	}
}

func NewTSXFromBytes(addrMode uint8, arg0 uint8, arg1 uint8) (*TSX, error) {
	switch addrMode {
	case AddrModeImplied:
		return NewTSX(), nil
	default:
		return nil, InvalidAddressModeError{
			AddressMode: addrMode,
		}
	}
}
