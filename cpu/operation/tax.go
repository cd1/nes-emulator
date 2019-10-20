package operation

import "io"

const (
	OpMnemonicTAX = "TAX"

	OpCodeTAX = 0xAA
)

func IsOpCodeValidTAX(opCode uint8) bool {
	return opCode == OpCodeTAX
}

func IsMnemonicValidTAX(mnemonic string) bool {
	return mnemonic == OpMnemonicTAX
}

type TAX struct {
	baseOperation
}

// 0xAA: TAX
func NewTAX() *TAX {
	return &TAX{
		baseOperation{
			code:        OpCodeTAX,
			addressMode: AddrModeImplied,
			mnemonic:    OpMnemonicTAX,
		},
	}
}

func NewTAXBinary(opCode uint8, data io.Reader) (*TAX, error) {
	switch opCode {
	case OpCodeTAX:
		return NewTAX(), nil
	default:
		return nil, InvalidOpCodeError{
			OpCode: opCode,
		}
	}
}

func NewTAXFromBytes(addrMode uint8, arg0 uint8, arg1 uint8) (*TAX, error) {
	switch addrMode {
	case AddrModeImplied:
		return NewTAX(), nil
	default:
		return nil, InvalidAddressModeError{
			AddressMode: addrMode,
		}
	}
}
