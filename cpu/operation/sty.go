package operation

import (
	"encoding/binary"
	"io"

	"github.com/cd1/nes-emulator/util"
)

const (
	OpMnemonicSTY = "STY"

	OpCodeSTYZero     = 0x84
	OpCodeSTYAbsolute = 0x8C
	OpCodeSTYZeroX    = 0x94
)

func IsOpCodeValidSTY(opCode uint8) bool {
	return opCode == OpCodeSTYZero ||
		opCode == OpCodeSTYAbsolute ||
		opCode == OpCodeSTYZeroX
}

func IsMnemonicValidSTY(mnemonic string) bool {
	return mnemonic == OpMnemonicSTY
}

type STY struct {
	baseOperation
}

// 0x84: STY $NN
func NewSTYZero(zeroAddress uint8) *STY {
	return &STY{
		baseOperation{
			code:        OpCodeSTYZero,
			addressMode: AddrModeZero,
			mnemonic:    OpMnemonicSTY,
			args:        []uint8{zeroAddress},
		},
	}
}

// 0x8C: STY $NNNN
func NewSTYAbsolute(absoluteAddress uint16) *STY {
	return &STY{
		baseOperation{
			code:        OpCodeSTYAbsolute,
			addressMode: AddrModeAbsolute,
			mnemonic:    OpMnemonicSTY,
			args:        util.BreakWordIntoBytes(absoluteAddress),
		},
	}
}

// 0x94: STY $NN, X
func NewSTYZeroX(zeroAddress uint8) *STY {
	return &STY{
		baseOperation{
			code:        OpCodeSTYZeroX,
			addressMode: AddrModeZeroX,
			mnemonic:    OpMnemonicSTY,
			args:        []uint8{zeroAddress},
		},
	}
}

func NewSTYBinary(opCode uint8, data io.Reader) (*STY, error) {
	switch opCode {
	case OpCodeSTYZero:
		var addr uint8

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewSTYZero(addr), nil
	case OpCodeSTYAbsolute:
		var addr uint16

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewSTYAbsolute(addr), nil
	case OpCodeSTYZeroX:
		var addr uint8

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewSTYZeroX(addr), nil
	default:
		return nil, InvalidOpCodeError{
			OpCode: opCode,
		}
	}
}

func NewSTYFromBytes(addrMode uint8, arg0 uint8, arg1 uint8) (*STY, error) {
	switch addrMode {
	case AddrModeZero, AddrModeRelative:
		return NewSTYZero(arg0), nil
	case AddrModeAbsolute:
		return NewSTYAbsolute(util.JoinBytesInWord([]uint8{arg0, arg1})), nil
	case AddrModeZeroX:
		return NewSTYZeroX(arg0), nil
	default:
		return nil, InvalidAddressModeError{
			AddressMode: addrMode,
		}
	}
}
