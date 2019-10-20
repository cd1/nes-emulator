package operation

import "io"

const (
	OpMnemonicSEI = "SEI"

	OpCodeSEI = 0x78
)

func IsOpCodeValidSEI(opCode uint8) bool {
	return opCode == OpCodeSEI
}

func IsMnemonicValidSEI(mnemonic string) bool {
	return mnemonic == OpMnemonicSEI
}

type SEI struct {
	baseOperation
}

// 0x78: SEI
func NewSEI() *SEI {
	return &SEI{
		baseOperation{
			code:        OpCodeSEI,
			addressMode: AddrModeImplied,
			mnemonic:    OpMnemonicSEI,
		},
	}
}

func NewSEIBinary(opCode uint8, data io.Reader) (*SEI, error) {
	switch opCode {
	case OpCodeSEI:
		return NewSEI(), nil
	default:
		return nil, InvalidOpCodeError{
			OpCode: opCode,
		}
	}
}

func NewSEIFromBytes(addrMode uint8, arg0 uint8, arg1 uint8) (*SEI, error) {
	switch addrMode {
	case AddrModeImplied:
		return NewSEI(), nil
	default:
		return nil, InvalidAddressModeError{
			AddressMode: addrMode,
		}
	}
}
