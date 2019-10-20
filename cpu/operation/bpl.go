package operation

import (
	"encoding/binary"
	"io"
)

const (
	OpMnemonicBPL = "BPL"

	OpCodeBPL = 0x10
)

func IsOpCodeValidBPL(opCode uint8) bool {
	return opCode == OpCodeBPL
}

func IsMnemonicValidBPL(mnemonic string) bool {
	return mnemonic == OpMnemonicBPL
}

type BPL struct {
	baseOperation
}

// 0x10: BPL $NN
func NewBPL(relativeAddress uint8) *BPL {
	return &BPL{
		baseOperation{
			code:        OpCodeBPL,
			addressMode: AddrModeRelative,
			mnemonic:    OpMnemonicBPL,
			args:        []uint8{relativeAddress},
		},
	}
}

func NewBPLBinary(opCode uint8, data io.Reader) (*BPL, error) {
	switch opCode {
	case OpCodeBPL:
		var addr uint8

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewBPL(addr), nil
	default:
		return nil, InvalidOpCodeError{
			OpCode: opCode,
		}
	}
}

func NewBPLFromBytes(addrMode uint8, arg0 uint8, arg1 uint8) (*BPL, error) {
	switch addrMode {
	case AddrModeZero, AddrModeRelative:
		return NewBPL(arg0), nil
	default:
		return nil, InvalidAddressModeError{
			AddressMode: addrMode,
		}
	}
}
