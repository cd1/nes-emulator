package operation

import (
	"encoding/binary"
	"io"
)

const (
	OpMnemonicBNE = "BNE"

	OpCodeBNE = 0xD0
)

func IsOpCodeValidBNE(opCode uint8) bool {
	return opCode == OpCodeBNE
}

func IsMnemonicValidBNE(mnemonic string) bool {
	return mnemonic == OpMnemonicBNE
}

type BNE struct {
	baseOperation
}

// 0xD0: BNE $NN
func NewBNE(relativeAddress uint8) *BNE {
	return &BNE{
		baseOperation{
			code:        OpCodeBNE,
			addressMode: AddrModeRelative,
			mnemonic:    OpMnemonicBNE,
			args:        []uint8{relativeAddress},
		},
	}
}

func NewBNEBinary(opCode uint8, data io.Reader) (*BNE, error) {
	switch opCode {
	case OpCodeBNE:
		var addr uint8

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewBNE(addr), nil
	default:
		return nil, InvalidOpCodeError{
			OpCode: opCode,
		}
	}
}

func NewBNEFromBytes(addrMode uint8, arg0 uint8, arg1 uint8) (*BNE, error) {
	switch addrMode {
	case AddrModeZero, AddrModeRelative:
		return NewBNE(arg0), nil
	default:
		return nil, InvalidAddressModeError{
			AddressMode: addrMode,
		}
	}
}
