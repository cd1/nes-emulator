package operation

import (
	"encoding/binary"
	"io"

	"github.com/cd1/nes-emulator/util"
)

const (
	OpMnemonicASL = "ASL"

	OpCodeASLZero        = 0x06
	OpCodeASLAccumulator = 0x0A
	OpCodeASLAbsolute    = 0x0E
	OpCodeASLZeroX       = 0x16
	OpCodeASLAbsoluteX   = 0x1E
)

func IsOpCodeValidASL(opCode uint8) bool {
	return opCode == OpCodeASLZero ||
		opCode == OpCodeASLAccumulator ||
		opCode == OpCodeASLAbsolute ||
		opCode == OpCodeASLZeroX ||
		opCode == OpCodeASLAbsoluteX
}

func IsMnemonicValidASL(mnemonic string) bool {
	return mnemonic == OpMnemonicASL
}

type ASL struct {
	baseOperation
}

func NewASLZero(zeroAddress uint8) *ASL {
	return &ASL{
		baseOperation{
			code:        OpCodeASLZero,
			addressMode: AddrModeZero,
			mnemonic:    OpMnemonicASL,
			args:        []uint8{zeroAddress},
		},
	}
}

func NewASLAccumulator() *ASL {
	return &ASL{
		baseOperation{
			code:        OpCodeASLAccumulator,
			addressMode: AddrModeAccumulator,
			mnemonic:    OpMnemonicASL,
		},
	}
}

func NewASLAbsolute(absoluteAddress uint16) *ASL {
	return &ASL{
		baseOperation{
			code:        OpCodeASLAbsolute,
			addressMode: AddrModeAbsolute,
			mnemonic:    OpMnemonicASL,
			args:        util.BreakWordIntoBytes(absoluteAddress),
		},
	}
}

func NewASLZeroX(zeroAddress uint8) *ASL {
	return &ASL{
		baseOperation{
			code:        OpCodeASLZeroX,
			addressMode: AddrModeZeroX,
			mnemonic:    OpMnemonicASL,
			args:        []uint8{zeroAddress},
		},
	}
}

func NewASLAbsoluteX(absoluteAddress uint16) *ASL {
	return &ASL{
		baseOperation{
			code:        OpCodeASLAbsoluteX,
			addressMode: AddrModeAbsoluteX,
			mnemonic:    OpMnemonicASL,
			args:        util.BreakWordIntoBytes(absoluteAddress),
		},
	}
}

func NewASLBinary(opCode uint8, data io.Reader) (*ASL, error) {
	switch opCode {
	case OpCodeASLZero:
		var addr uint8

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewASLZero(addr), nil
	case OpCodeASLAccumulator:
		return NewASLAccumulator(), nil
	case OpCodeASLAbsolute:
		var addr uint16

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewASLAbsolute(addr), nil
	case OpCodeASLZeroX:
		var addr uint8

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewASLZeroX(addr), nil
	case OpCodeASLAbsoluteX:
		var addr uint16

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewASLAbsoluteX(addr), nil
	default:
		return nil, InvalidOpCodeError{
			OpCode: opCode,
		}
	}
}

func NewASLFromBytes(addrMode uint8, arg0 uint8, arg1 uint8) (*ASL, error) {
	switch addrMode {
	case AddrModeZero, AddrModeRelative:
		return NewASLZero(arg0), nil
	case AddrModeAccumulator:
		return NewASLAccumulator(), nil
	case AddrModeAbsolute:
		return NewASLAbsolute(util.JoinBytesInWord([]uint8{arg0, arg1})), nil
	case AddrModeZeroX:
		return NewASLZeroX(arg0), nil
	case AddrModeAbsoluteX:
		return NewASLAbsoluteX(util.JoinBytesInWord([]uint8{arg0, arg1})), nil
	default:
		return nil, InvalidAddressModeError{
			AddressMode: addrMode,
		}
	}
}
