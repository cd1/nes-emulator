package operation

import (
	"encoding/binary"
	"io"

	"github.com/cd1/nes-emulator/util"
)

const (
	OpMnemonicSTX = "STX"

	OpCodeSTXZero     = 0x86
	OpCodeSTXAbsolute = 0x8E
	OpCodeSTXZeroY    = 0x96
)

func IsOpCodeValidSTX(opCode uint8) bool {
	return opCode == OpCodeSTXZero ||
		opCode == OpCodeSTXAbsolute ||
		opCode == OpCodeSTXZeroY
}

func IsMnemonicValidSTX(mnemonic string) bool {
	return mnemonic == OpMnemonicSTX
}

type STX struct {
	baseOperation
}

// 0x86: STX $NN
func NewSTXZero(zeroAddress uint8) *STX {
	return &STX{
		baseOperation{
			code:        OpCodeSTXZero,
			addressMode: AddrModeZero,
			mnemonic:    OpMnemonicSTX,
			args:        []uint8{zeroAddress},
		},
	}
}

// 0x8E: STX $NNNN
func NewSTXAbsolute(absoluteAddress uint16) *STX {
	return &STX{
		baseOperation{
			code:        OpCodeSTXAbsolute,
			addressMode: AddrModeAbsolute,
			mnemonic:    OpMnemonicSTX,
			args:        util.BreakWordIntoBytes(absoluteAddress),
		},
	}
}

// 0x96: STX $NN, Y
func NewSTXZeroY(zeroAddress uint8) *STX {
	return &STX{
		baseOperation{
			code:        OpCodeSTXZeroY,
			addressMode: AddrModeZeroY,
			mnemonic:    OpMnemonicSTX,
			args:        []uint8{zeroAddress},
		},
	}
}

func NewSTXBinary(opCode uint8, data io.Reader) (*STX, error) {
	switch opCode {
	case OpCodeSTXZero:
		var addr uint8

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewSTXZero(addr), nil
	case OpCodeSTXAbsolute:
		var addr uint16

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewSTXAbsolute(addr), nil
	case OpCodeSTXZeroY:
		var addr uint8

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewSTXZeroY(addr), nil
	default:
		return nil, InvalidOpCodeError{
			OpCode: opCode,
		}
	}
}

func NewSTXFromBytes(addrMode uint8, arg0 uint8, arg1 uint8) (*STX, error) {
	switch addrMode {
	case AddrModeZero, AddrModeRelative:
		return NewSTXZero(arg0), nil
	case AddrModeAbsolute:
		return NewSTXAbsolute(util.JoinBytesInWord([]uint8{arg0, arg1})), nil
	case AddrModeZeroY:
		return NewSTXZeroY(arg0), nil
	default:
		return nil, InvalidAddressModeError{
			AddressMode: addrMode,
		}
	}
}
