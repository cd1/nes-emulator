package operation

import "io"

const (
	OpMnemonicCLD = "CLD"

	OpCodeCLD = 0xD8
)

func IsOpCodeValidCLD(opCode uint8) bool {
	return opCode == OpCodeCLD
}

func IsMnemonicValidCLD(mnemonic string) bool {
	return mnemonic == OpMnemonicCLD
}

type CLD struct {
	baseOperation
}

// 0xD8: CLD
func NewCLD() *CLD {
	return &CLD{
		baseOperation{
			code:        OpCodeCLD,
			addressMode: AddrModeImplied,
			mnemonic:    OpMnemonicCLD,
		},
	}
}

func NewCLDBinary(opCode uint8, data io.Reader) (*CLD, error) {
	switch opCode {
	case OpCodeCLD:
		return NewCLD(), nil
	default:
		return nil, InvalidOpCodeError{
			OpCode: opCode,
		}
	}
}

func NewCLDFromBytes(addrMode uint8, arg0 uint8, arg1 uint8) (*CLD, error) {
	switch addrMode {
	case AddrModeImplied:
		return NewCLD(), nil
	default:
		return nil, InvalidAddressModeError{
			AddressMode: addrMode,
		}
	}
}
