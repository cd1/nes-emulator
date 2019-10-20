package operation

import (
	"encoding/binary"
	"io"

	"github.com/cd1/nes-emulator/util"
)

const (
	OpMnemonicLDX = "LDX"

	OpCodeLDXImmediate = 0xA2
	OpCodeLDXZero      = 0xA6
	OpCodeLDXAbsolute  = 0xAE
	OpCodeLDXZeroY     = 0xB6
	OpCodeLDXAbsoluteY = 0xBE
)

func IsOpCodeValidLDX(opCode uint8) bool {
	return opCode == OpCodeLDXImmediate ||
		opCode == OpCodeLDXZero ||
		opCode == OpCodeLDXAbsolute ||
		opCode == OpCodeLDXZeroY ||
		opCode == OpCodeLDXAbsoluteY
}

func IsMnemonicValidLDX(mnemonic string) bool {
	return mnemonic == OpMnemonicLDX
}

type LDX struct {
	baseOperation
}

func NewLDXImmediate(value uint8) *LDX {
	return &LDX{
		baseOperation{
			code:        OpCodeLDXImmediate,
			addressMode: AddrModeImmediate,
			mnemonic:    OpMnemonicLDX,
			args:        []uint8{value},
		},
	}
}

func NewLDXZero(zeroAddress uint8) *LDX {
	return &LDX{
		baseOperation{
			code:        OpCodeLDXZero,
			addressMode: AddrModeZero,
			mnemonic:    OpMnemonicLDX,
			args:        []uint8{zeroAddress},
		},
	}
}

func NewLDXAbsolute(absoluteAddress uint16) *LDX {
	return &LDX{
		baseOperation{
			code:        OpCodeLDXAbsolute,
			addressMode: AddrModeAbsolute,
			mnemonic:    OpMnemonicLDX,
			args:        util.BreakWordIntoBytes(absoluteAddress),
		},
	}
}

func NewLDXZeroY(zeroAddress uint8) *LDX {
	return &LDX{
		baseOperation{
			code:        OpCodeLDXZeroY,
			addressMode: AddrModeZeroY,
			mnemonic:    OpMnemonicLDX,
			args:        []uint8{zeroAddress},
		},
	}
}

func NewLDXAbsoluteY(absoluteAddress uint16) *LDX {
	return &LDX{
		baseOperation{
			code:        OpCodeLDXAbsoluteY,
			addressMode: AddrModeAbsoluteY,
			mnemonic:    OpMnemonicLDX,
			args:        util.BreakWordIntoBytes(absoluteAddress),
		},
	}
}

func NewLDXBinary(opCode uint8, data io.Reader) (*LDX, error) {
	switch opCode {
	case OpCodeLDXImmediate:
		var addr uint8

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewLDXImmediate(addr), nil
	case OpCodeLDXZero:
		var addr uint8

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewLDXZero(addr), nil
	case OpCodeLDXAbsolute:
		var addr uint16

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewLDXAbsolute(addr), nil
	case OpCodeLDXZeroY:
		var addr uint8

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewLDXZeroY(addr), nil
	case OpCodeLDXAbsoluteY:
		var addr uint16

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewLDXAbsoluteY(addr), nil
	default:
		return nil, InvalidOpCodeError{
			OpCode: opCode,
		}
	}
}

func NewLDXFromBytes(addrMode uint8, arg0 uint8, arg1 uint8) (*LDX, error) {
	switch addrMode {
	case AddrModeImmediate:
		return NewLDXImmediate(arg0), nil
	case AddrModeZero, AddrModeRelative:
		return NewLDXZero(arg0), nil
	case AddrModeAbsolute:
		return NewLDXAbsolute(util.JoinBytesInWord([]uint8{arg0, arg1})), nil
	case AddrModeZeroY:
		return NewLDXZeroY(arg0), nil
	case AddrModeAbsoluteY:
		return NewLDXAbsoluteY(util.JoinBytesInWord([]uint8{arg0, arg1})), nil
	default:
		return nil, InvalidAddressModeError{
			AddressMode: addrMode,
		}
	}
}
