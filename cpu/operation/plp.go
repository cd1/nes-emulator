package operation

import "io"

const (
	OpMnemonicPLP = "PLP"

	OpCodePLP = 0x28
)

func IsOpCodeValidPLP(opCode uint8) bool {
	return opCode == OpCodePLP
}

func IsMnemonicValidPLP(mnemonic string) bool {
	return mnemonic == OpMnemonicPLP
}

type PLP struct {
	baseOperation
}

// 0x28: PLP
func NewPLP() *PLP {
	return &PLP{
		baseOperation{
			code:        OpCodePLP,
			addressMode: AddrModeImplied,
			mnemonic:    OpMnemonicPLP,
		},
	}
}

func NewPLPBinary(opCode uint8, data io.Reader) (*PLP, error) {
	switch opCode {
	case OpCodePLP:
		return NewPLP(), nil
	default:
		return nil, InvalidOpCodeError{
			OpCode: opCode,
		}
	}
}

func NewPLPFromBytes(addrMode uint8, arg0 uint8, arg1 uint8) (*PLP, error) {
	switch addrMode {
	case AddrModeImplied:
		return NewPLP(), nil
	default:
		return nil, InvalidAddressModeError{
			AddressMode: addrMode,
		}
	}
}
