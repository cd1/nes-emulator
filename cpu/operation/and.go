package operation

import (
	"encoding/binary"
	"io"

	"github.com/cd1/nes-emulator/util"
)

const (
	OpMnemonicAND = "AND"

	OpCodeANDIndirectX = 0x21
	OpCodeANDZero      = 0x25
	OpCodeANDImmediate = 0x29
	OpCodeANDAbsolute  = 0x2D
	OpCodeANDIndirectY = 0x31
	OpCodeANDZeroX     = 0x35
	OpCodeANDAbsoluteY = 0x39
	OpCodeANDAbsoluteX = 0x3D
)

func IsOpCodeValidAND(opCode uint8) bool {
	return opCode == OpCodeANDIndirectX ||
		opCode == OpCodeANDZero ||
		opCode == OpCodeANDImmediate ||
		opCode == OpCodeANDAbsolute ||
		opCode == OpCodeANDIndirectY ||
		opCode == OpCodeANDZeroX ||
		opCode == OpCodeANDAbsoluteY ||
		opCode == OpCodeANDAbsoluteX
}

func IsMnemonicValidAND(mnemonic string) bool {
	return mnemonic == OpMnemonicAND
}

type AND struct {
	baseOperation
}

func NewANDIndirectX(indirectAddress uint8) *AND {
	return &AND{
		baseOperation{
			code:        OpCodeANDIndirectX,
			addressMode: AddrModeIndirectX,
			mnemonic:    OpMnemonicAND,
			args:        []uint8{indirectAddress},
		},
	}
}

func NewANDZero(zeroAddress uint8) *AND {
	return &AND{
		baseOperation{
			code:        OpCodeANDZero,
			addressMode: AddrModeZero,
			mnemonic:    OpMnemonicAND,
			args:        []uint8{zeroAddress},
		},
	}
}

func NewANDImmediate(value uint8) *AND {
	return &AND{
		baseOperation{
			code:        OpCodeANDImmediate,
			addressMode: AddrModeImmediate,
			mnemonic:    OpMnemonicAND,
			args:        []uint8{value},
		},
	}
}

func NewANDAbsolute(absoluteAddress uint16) *AND {
	return &AND{
		baseOperation{
			code:        OpCodeANDAbsolute,
			addressMode: AddrModeAbsolute,
			mnemonic:    OpMnemonicAND,
			args:        util.BreakWordIntoBytes(absoluteAddress),
		},
	}
}

func NewANDIndirectY(indirectAddress uint8) *AND {
	return &AND{
		baseOperation{
			code:        OpCodeANDIndirectY,
			addressMode: AddrModeIndirectY,
			mnemonic:    OpMnemonicAND,
			args:        []uint8{indirectAddress},
		},
	}
}

func NewANDZeroX(zeroAddress uint8) *AND {
	return &AND{
		baseOperation{
			code:        OpCodeANDZeroX,
			addressMode: AddrModeZeroX,
			mnemonic:    OpMnemonicAND,
			args:        []uint8{zeroAddress},
		},
	}
}

func NewANDAbsoluteY(absoluteAddress uint16) *AND {
	return &AND{
		baseOperation{
			code:        OpCodeANDAbsoluteY,
			addressMode: AddrModeAbsoluteY,
			mnemonic:    OpMnemonicAND,
			args:        util.BreakWordIntoBytes(absoluteAddress),
		},
	}
}

func NewANDAbsoluteX(absoluteAddress uint16) *AND {
	return &AND{
		baseOperation{
			code:        OpCodeANDAbsoluteX,
			addressMode: AddrModeAbsoluteX,
			mnemonic:    OpMnemonicAND,
			args:        util.BreakWordIntoBytes(absoluteAddress),
		},
	}
}

func NewANDBinary(opCode uint8, data io.Reader) (*AND, error) {
	switch opCode {
	case OpCodeANDIndirectX:
		var addr uint8

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewANDIndirectX(addr), nil
	case OpCodeANDZero:
		var addr uint8

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewANDZero(addr), nil
	case OpCodeANDImmediate:
		var addr uint8

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewANDImmediate(addr), nil
	case OpCodeANDAbsolute:
		var addr uint16

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewANDAbsolute(addr), nil
	case OpCodeANDIndirectY:
		var addr uint8

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewANDIndirectY(addr), nil
	case OpCodeANDZeroX:
		var addr uint8

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewANDZeroX(addr), nil
	case OpCodeANDAbsoluteY:
		var addr uint16

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewANDAbsoluteY(addr), nil
	case OpCodeANDAbsoluteX:
		var addr uint16

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewANDAbsoluteX(addr), nil
	default:
		return nil, InvalidOpCodeError{opCode}
	}
}

func NewANDFromBytes(addrMode uint8, arg0 uint8, arg1 uint8) (*AND, error) {
	switch addrMode {
	case AddrModeIndirectX:
		return NewANDIndirectX(arg0), nil
	case AddrModeZero, AddrModeRelative:
		return NewANDZero(arg0), nil
	case AddrModeImmediate:
		return NewANDImmediate(arg0), nil
	case AddrModeAbsolute:
		return NewANDAbsolute(util.JoinBytesInWord([]uint8{arg0, arg1})), nil
	case AddrModeIndirectY:
		return NewANDIndirectY(arg0), nil
	case AddrModeZeroX:
		return NewANDZeroX(arg0), nil
	case AddrModeAbsoluteY:
		return NewANDAbsoluteY(util.JoinBytesInWord([]uint8{arg0, arg1})), nil
	case AddrModeAbsoluteX:
		return NewANDAbsoluteX(util.JoinBytesInWord([]uint8{arg0, arg1})), nil
	default:
		return nil, InvalidAddressModeError{
			AddressMode: addrMode,
		}
	}
}
