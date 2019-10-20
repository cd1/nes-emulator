package operation

import (
	"encoding/binary"
	"io"

	"github.com/cd1/nes-emulator/util"
)

const (
	OpMnemonicBIT = "BIT"

	OpCodeBITZero     = 0x24
	OpCodeBITAbsolute = 0x2C
)

func IsOpCodeValidBIT(opCode uint8) bool {
	return opCode == OpCodeBITZero ||
		opCode == OpCodeBITAbsolute
}

func IsMnemonicValidBIT(mnemonic string) bool {
	return mnemonic == OpMnemonicBIT
}

type BIT struct {
	baseOperation
}

func NewBITZero(zeroAddress uint8) *BIT {
	return &BIT{
		baseOperation{
			code:        OpCodeBITZero,
			addressMode: AddrModeZero,
			mnemonic:    OpMnemonicBIT,
			args:        []uint8{zeroAddress},
		},
	}
}

func NewBITAbsolute(absoluteAddress uint16) *BIT {
	return &BIT{
		baseOperation{
			code:        OpCodeBITAbsolute,
			addressMode: AddrModeAbsolute,
			mnemonic:    OpMnemonicBIT,
			args:        util.BreakWordIntoBytes(absoluteAddress),
		},
	}
}

func NewBITBinary(opCode uint8, data io.Reader) (*BIT, error) {
	switch opCode {
	case OpCodeBITZero:
		var addr uint8

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewBITZero(addr), nil
	case OpCodeBITAbsolute:
		var addr uint16

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewBITAbsolute(addr), nil
	default:
		return nil, InvalidOpCodeError{
			OpCode: opCode,
		}
	}
}

func NewBITFromBytes(addrMode uint8, arg0 uint8, arg1 uint8) (*BIT, error) {
	switch addrMode {
	case AddrModeZero, AddrModeRelative:
		return NewBITZero(arg0), nil
	case AddrModeAbsolute:
		return NewBITAbsolute(util.JoinBytesInWord([]uint8{arg0, arg1})), nil
	default:
		return nil, InvalidAddressModeError{
			AddressMode: addrMode,
		}
	}
}
