package operation

import (
	"encoding/binary"
	"io"
)

const (
	OpMnemonicBCS = "BCS"

	OpCodeBCS = 0xB0
)

func IsOpCodeValidBCS(opCode uint8) bool {
	return opCode == OpCodeBCS
}

func IsMnemonicValidBCS(mnemonic string) bool {
	return mnemonic == OpMnemonicBCS
}

type BCS struct {
	baseOperation
}

// 0xB0: BCS $NN
func NewBCS(relativeAddress uint8) *BCS {
	return &BCS{
		baseOperation{
			code:        OpCodeBCS,
			addressMode: AddrModeRelative,
			mnemonic:    OpMnemonicBCS,
			args:        []uint8{relativeAddress},
		},
	}
}

func NewBCSBinary(opCode uint8, data io.Reader) (*BCS, error) {
	switch opCode {
	case OpCodeBCS:
		var addr uint8

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewBCS(addr), nil
	default:
		return nil, InvalidOpCodeError{
			OpCode: opCode,
		}
	}
}

func NewBCSFromBytes(addrMode uint8, arg0 uint8, arg1 uint8) (*BCS, error) {
	switch addrMode {
	case AddrModeZero, AddrModeRelative:
		return NewBCS(arg0), nil
	default:
		return nil, InvalidAddressModeError{
			AddressMode: addrMode,
		}
	}
}
