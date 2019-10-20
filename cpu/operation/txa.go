package operation

import "io"

const (
	OpMnemonicTXA = "TXA"

	OpCodeTXA = 0x8A
)

func IsOpCodeValidTXA(opCode uint8) bool {
	return opCode == OpCodeTXA
}

func IsMnemonicValidTXA(mnemonic string) bool {
	return mnemonic == OpMnemonicTXA
}

type TXA struct {
	baseOperation
}

// 0x8A: TXA
func NewTXA() *TXA {
	return &TXA{
		baseOperation{
			code:        OpCodeTXA,
			addressMode: AddrModeImplied,
			mnemonic:    OpMnemonicTXA,
		},
	}
}

func NewTXABinary(opCode uint8, data io.Reader) (*TXA, error) {
	switch opCode {
	case OpCodeTXA:
		return NewTXA(), nil
	default:
		return nil, InvalidOpCodeError{
			OpCode: opCode,
		}
	}
}

func NewTXAFromBytes(addrMode uint8, arg0 uint8, arg1 uint8) (*TXA, error) {
	switch addrMode {
	case AddrModeImplied:
		return NewTXA(), nil
	default:
		return nil, InvalidAddressModeError{
			AddressMode: addrMode,
		}
	}
}
