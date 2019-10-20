package operation

import "io"

const (
	OpMnemonicPHP = "PHP"

	OpCodePHP = 0x08
)

func IsOpCodeValidPHP(opCode uint8) bool {
	return opCode == OpCodePHP
}

func IsMnemonicValidPHP(mnemonic string) bool {
	return mnemonic == OpMnemonicPHP
}

type PHP struct {
	baseOperation
}

// 0x08: PHP
func NewPHP() *PHP {
	return &PHP{
		baseOperation{
			code:        OpCodePHP,
			addressMode: AddrModeImplied,
			mnemonic:    OpMnemonicPHP,
		},
	}
}

func NewPHPBinary(opCode uint8, data io.Reader) (*PHP, error) {
	switch opCode {
	case OpCodePHP:
		return NewPHP(), nil
	default:
		return nil, InvalidOpCodeError{
			OpCode: opCode,
		}
	}
}

func NewPHPFromBytes(addrMode uint8, arg0 uint8, arg1 uint8) (*PHP, error) {
	switch addrMode {
	case AddrModeImplied:
		return NewPHP(), nil
	default:
		return nil, InvalidAddressModeError{
			AddressMode: addrMode,
		}
	}
}
