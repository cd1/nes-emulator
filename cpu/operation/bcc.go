package operation

import (
	"encoding/binary"
	"io"
)

const (
	OpMnemonicBCC = "BCC"

	OpCodeBCC = 0x90
)

func IsOpCodeValidBCC(opCode uint8) bool {
	return opCode == OpCodeBCC
}

func IsMnemonicValidBCC(mnemonic string) bool {
	return mnemonic == OpMnemonicBCC
}

type BCC struct {
	baseOperation
}

// 0x90: BCC $NN
func NewBCC(relativeAddress uint8) *BCC {
	return &BCC{
		baseOperation{
			code:        OpCodeBCC,
			addressMode: AddrModeRelative,
			mnemonic:    OpMnemonicBCC,
			args:        []uint8{relativeAddress},
		},
	}
}

func NewBCCBinary(opCode uint8, data io.Reader) (*BCC, error) {
	switch opCode {
	case OpCodeBCC:
		var addr uint8

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewBCC(addr), nil
	default:
		return nil, InvalidOpCodeError{
			OpCode: opCode,
		}
	}
}

func NewBCCFromBytes(addrMode uint8, arg0 uint8, arg1 uint8) (*BCC, error) {
	switch addrMode {
	case AddrModeZero, AddrModeRelative:
		return NewBCC(arg0), nil
	default:
		return nil, InvalidAddressModeError{
			AddressMode: addrMode,
		}
	}
}
