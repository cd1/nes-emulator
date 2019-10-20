package operation

import (
	"encoding/binary"
	"io"

	"github.com/cd1/nes-emulator/util"
)

const (
	OpMnemonicROL = "ROL"

	OpCodeROLZero        = 0x26
	OpCodeROLAccumulator = 0x2A
	OpCodeROLAbsolute    = 0x2E
	OpCodeROLZeroX       = 0x36
	OpCodeROLAbsoluteX   = 0x3E
)

func IsOpCodeValidROL(opCode uint8) bool {
	return opCode == OpCodeROLZero ||
		opCode == OpCodeROLAccumulator ||
		opCode == OpCodeROLAbsolute ||
		opCode == OpCodeROLZeroX ||
		opCode == OpCodeROLAbsoluteX
}

func IsMnemonicValidROL(mnemonic string) bool {
	return mnemonic == OpMnemonicROL
}

type ROL struct {
	baseOperation
}

// 0x26: ROL $NN
func NewROLZero(zeroAddress uint8) *ROL {
	return &ROL{
		baseOperation{
			code:        OpCodeROLZero,
			addressMode: AddrModeZero,
			mnemonic:    OpMnemonicROL,
			args:        []uint8{zeroAddress},
		},
	}
}

// 0x2A: ROL A
func NewROLAccumulator() *ROL {
	return &ROL{
		baseOperation{
			code:        OpCodeROLAccumulator,
			addressMode: AddrModeAccumulator,
			mnemonic:    OpMnemonicROL,
		},
	}
}

// 0x2E: ROL $NNNN
func NewROLAbsolute(absoluteAddress uint16) *ROL {
	return &ROL{
		baseOperation{
			code:        OpCodeROLAbsolute,
			addressMode: AddrModeAbsolute,
			mnemonic:    OpMnemonicROL,
			args:        util.BreakWordIntoBytes(absoluteAddress),
		},
	}
}

// 0x36: ROL $NN, X
func NewROLZeroX(zeroAddress uint8) *ROL {
	return &ROL{
		baseOperation{
			code:        OpCodeROLZeroX,
			addressMode: AddrModeZeroX,
			mnemonic:    OpMnemonicROL,
			args:        []uint8{zeroAddress},
		},
	}
}

// 0x3E: ROL $NNNN, X
func NewROLAbsoluteX(absoluteAddress uint16) *ROL {
	return &ROL{
		baseOperation{
			code:        OpCodeROLAbsoluteX,
			addressMode: AddrModeAbsoluteX,
			mnemonic:    OpMnemonicROL,
			args:        util.BreakWordIntoBytes(absoluteAddress),
		},
	}
}

func NewROLBinary(opCode uint8, data io.Reader) (*ROL, error) {
	switch opCode {
	case OpCodeROLZero:
		var addr uint8

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewROLZero(addr), nil
	case OpCodeROLAccumulator:
		return NewROLAccumulator(), nil
	case OpCodeROLAbsolute:
		var addr uint16

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewROLAbsolute(addr), nil
	case OpCodeROLZeroX:
		var addr uint8

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewROLZeroX(addr), nil
	case OpCodeROLAbsoluteX:
		var addr uint16

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewROLAbsoluteX(addr), nil
	default:
		return nil, InvalidOpCodeError{
			OpCode: opCode,
		}
	}
}

func NewROLFromBytes(addrMode uint8, arg0 uint8, arg1 uint8) (*ROL, error) {
	switch addrMode {
	case AddrModeZero, AddrModeRelative:
		return NewROLZero(arg0), nil
	case AddrModeAccumulator:
		return NewROLAccumulator(), nil
	case AddrModeAbsolute:
		return NewROLAbsolute(util.JoinBytesInWord([]uint8{arg0, arg1})), nil
	case AddrModeZeroX:
		return NewROLZeroX(arg0), nil
	case AddrModeAbsoluteX:
		return NewROLAbsoluteX(util.JoinBytesInWord([]uint8{arg0, arg1})), nil
	default:
		return nil, InvalidAddressModeError{
			AddressMode: addrMode,
		}
	}
}
