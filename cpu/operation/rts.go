package operation

import "io"

const (
	OpMnemonicRTS = "RTS"

	OpCodeRTS = 0x60
)

func IsOpCodeValidRTS(opCode uint8) bool {
	return opCode == OpCodeRTS
}

func IsMnemonicValidRTS(mnemonic string) bool {
	return mnemonic == OpMnemonicRTS
}

type RTS struct {
	baseOperation
}

// 0x60: RTS
func NewRTS() *RTS {
	return &RTS{
		baseOperation{
			code:        OpCodeRTS,
			addressMode: AddrModeImplied,
			mnemonic:    OpMnemonicRTS,
		},
	}
}

func NewRTSBinary(opCode uint8, data io.Reader) (*RTS, error) {
	switch opCode {
	case OpCodeRTS:
		return NewRTS(), nil
	default:
		return nil, InvalidOpCodeError{
			OpCode: opCode,
		}
	}
}

func NewRTSFromBytes(addrMode uint8, arg0 uint8, arg1 uint8) (*RTS, error) {
	switch addrMode {
	case AddrModeImplied:
		return NewRTS(), nil
	default:
		return nil, InvalidAddressModeError{
			AddressMode: addrMode,
		}
	}
}
