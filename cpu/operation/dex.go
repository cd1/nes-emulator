package operation

import "io"

const (
	OpMnemonicDEX = "DEX"

	OpCodeDEX = 0xCA
)

func IsOpCodeValidDEX(opCode uint8) bool {
	return opCode == OpCodeDEX
}

func IsMnemonicValidDEX(mnemonic string) bool {
	return mnemonic == OpMnemonicDEX
}

type DEX struct {
	baseOperation
}

// 0xCA: DEX
func NewDEX() *DEX {
	return &DEX{
		baseOperation{
			code:        OpCodeDEX,
			addressMode: AddrModeImplied,
			mnemonic:    OpMnemonicDEX,
		},
	}
}

func NewDEXBinary(opCode uint8, data io.Reader) (*DEX, error) {
	switch opCode {
	case OpCodeDEX:
		return NewDEX(), nil
	default:
		return nil, InvalidOpCodeError{
			OpCode: opCode,
		}
	}
}

func NewDEXFromBytes(addrMode uint8, arg0 uint8, arg1 uint8) (*DEX, error) {
	switch addrMode {
	case AddrModeImplied:
		return NewDEX(), nil
	default:
		return nil, InvalidAddressModeError{
			AddressMode: addrMode,
		}
	}
}
