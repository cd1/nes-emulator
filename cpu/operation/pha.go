package operation

import "io"

const (
	OpMnemonicPHA = "PHA"

	OpCodePHA = 0x48
)

func IsOpCodeValidPHA(opCode uint8) bool {
	return opCode == OpCodePHA
}

func IsMnemonicValidPHA(mnemonic string) bool {
	return mnemonic == OpMnemonicPHA
}

type PHA struct {
	baseOperation
}

// 0x48: PHA
func NewPHA() *PHA {
	return &PHA{
		baseOperation{
			code:        OpCodePHA,
			addressMode: AddrModeImplied,
			mnemonic:    OpMnemonicPHA,
		},
	}
}

func NewPHABinary(opCode uint8, data io.Reader) (*PHA, error) {
	switch opCode {
	case OpCodePHA:
		return NewPHA(), nil
	default:
		return nil, InvalidOpCodeError{
			OpCode: opCode,
		}
	}
}

func NewPHAFromBytes(addrMode uint8, arg0 uint8, arg1 uint8) (*PHA, error) {
	switch addrMode {
	case AddrModeImplied:
		return NewPHA(), nil
	default:
		return nil, InvalidAddressModeError{
			AddressMode: addrMode,
		}
	}
}
