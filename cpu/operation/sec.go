package operation

import (
	"io"
)

const (
	OpMnemonicSEC = "SEC"

	OpCodeSEC = 0x38
)

func IsOpCodeValidSEC(opCode uint8) bool {
	return opCode == OpCodeSEC
}

func IsMnemonicValidSEC(mnemonic string) bool {
	return mnemonic == OpMnemonicSEC
}

type SEC struct {
	baseOperation
}

// 0x38: SEC
func NewSEC() *SEC {
	return &SEC{
		baseOperation{
			code:        OpCodeSEC,
			addressMode: AddrModeImplied,
			mnemonic:    OpMnemonicSEC,
		},
	}
}

func NewSECBinary(opCode uint8, data io.Reader) (*SEC, error) {
	switch opCode {
	case OpCodeSEC:
		return NewSEC(), nil
	default:
		return nil, InvalidOpCodeError{
			OpCode: opCode,
		}
	}
}

func NewSECFromBytes(addrMode uint8, arg0 uint8, arg1 uint8) (*SEC, error) {
	switch addrMode {
	case AddrModeImplied:
		return NewSEC(), nil
	default:
		return nil, InvalidAddressModeError{
			AddressMode: addrMode,
		}
	}
}
