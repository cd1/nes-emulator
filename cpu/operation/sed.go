package operation

import "io"

const (
	OpMnemonicSED = "SED"

	OpCodeSED = 0xF8
)

func IsOpCodeValidSED(opCode uint8) bool {
	return opCode == OpCodeSED
}

func IsMnemonicValidSED(mnemonic string) bool {
	return mnemonic == OpMnemonicSED
}

type SED struct {
	baseOperation
}

// 0xF8: SED
func NewSED() *SED {
	return &SED{
		baseOperation{
			code:        OpCodeSED,
			addressMode: AddrModeImplied,
			mnemonic:    OpMnemonicSED,
		},
	}
}

func NewSEDBinary(opCode uint8, data io.Reader) (*SED, error) {
	switch opCode {
	case OpCodeSED:
		return NewSED(), nil
	default:
		return nil, InvalidOpCodeError{
			OpCode: opCode,
		}
	}
}

func NewSEDFromBytes(addrMode uint8, arg0 uint8, arg1 uint8) (*SED, error) {
	switch addrMode {
	case AddrModeImplied:
		return NewSED(), nil
	default:
		return nil, InvalidAddressModeError{
			AddressMode: addrMode,
		}
	}
}
