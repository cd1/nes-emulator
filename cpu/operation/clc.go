package operation

import (
	"io"
)

const (
	OpMnemonicCLC = "CLC"

	OpCodeCLC = 0x18
)

func IsOpCodeValidCLC(opCode uint8) bool {
	return opCode == OpCodeCLC
}

func IsMnemonicValidCLC(mnemonic string) bool {
	return mnemonic == OpMnemonicCLC
}

type CLC struct {
	baseOperation
}

// 0x18: CLC
func NewCLC() *CLC {
	return &CLC{
		baseOperation{
			code:        OpCodeCLC,
			addressMode: AddrModeImplied,
			mnemonic:    OpMnemonicCLC,
		},
	}
}

func NewCLCBinary(opCode uint8, data io.Reader) (*CLC, error) {
	switch opCode {
	case OpCodeCLC:
		return NewCLC(), nil
	default:
		return nil, InvalidOpCodeError{
			OpCode: opCode,
		}
	}
}

func NewCLCFromBytes(addrMode uint8, arg0 uint8, arg1 uint8) (*CLC, error) {
	switch addrMode {
	case AddrModeImplied:
		return NewCLC(), nil
	default:
		return nil, InvalidAddressModeError{
			AddressMode: addrMode,
		}
	}
}
