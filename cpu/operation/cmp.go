package operation

import (
	"encoding/binary"
	"io"

	"github.com/cd1/nes-emulator/util"
)

const (
	OpMnemonicCMP = "CMP"

	OpCodeCMPIndirectX = 0xC1
	OpCodeCMPZero      = 0xC5
	OpCodeCMPImmediate = 0xC9
	OpCodeCMPAbsolute  = 0xCD
	OpCodeCMPIndirectY = 0xD1
	OpCodeCMPZeroX     = 0xD5
	OpCodeCMPAbsoluteY = 0xD9
	OpCodeCMPAbsoluteX = 0xDD
)

func IsOpCodeValidCMP(opCode uint8) bool {
	return opCode == OpCodeCMPIndirectX ||
		opCode == OpCodeCMPZero ||
		opCode == OpCodeCMPImmediate ||
		opCode == OpCodeCMPAbsolute ||
		opCode == OpCodeCMPIndirectY ||
		opCode == OpCodeCMPZeroX ||
		opCode == OpCodeCMPAbsoluteY ||
		opCode == OpCodeCMPAbsoluteX
}

func IsMnemonicValidCMP(mnemonic string) bool {
	return mnemonic == OpMnemonicCMP
}

type CMP struct {
	baseOperation
}

// 0xC1: CMP ($NN, X)
func NewCMPIndirectX(indirectAddress uint8) *CMP {
	return &CMP{
		baseOperation{
			code:        OpCodeCMPIndirectX,
			addressMode: AddrModeIndirectX,
			mnemonic:    OpMnemonicCMP,
			args:        []uint8{indirectAddress},
		},
	}
}

// 0xC5: CMP $NN
func NewCMPZero(zeroAddress uint8) *CMP {
	return &CMP{
		baseOperation{
			code:        OpCodeCMPZero,
			addressMode: AddrModeZero,
			mnemonic:    OpMnemonicCMP,
			args:        []uint8{zeroAddress},
		},
	}
}

// 0xC9: CMP #$NN
func NewCMPImmediate(value uint8) *CMP {
	return &CMP{
		baseOperation{
			code:        OpCodeCMPImmediate,
			addressMode: AddrModeImmediate,
			mnemonic:    OpMnemonicCMP,
			args:        []uint8{value},
		},
	}
}

// 0xCD: CMP $NNNN
func NewCMPAbsolute(absoluteAddress uint16) *CMP {
	return &CMP{
		baseOperation{
			code:        OpCodeCMPAbsolute,
			addressMode: AddrModeAbsolute,
			mnemonic:    OpMnemonicCMP,
			args:        util.BreakWordIntoBytes(absoluteAddress),
		},
	}
}

// 0xD1: CMP ($NN), Y
func NewCMPIndirectY(indirectAddress uint8) *CMP {
	return &CMP{
		baseOperation{
			code:        OpCodeCMPIndirectY,
			addressMode: AddrModeIndirectY,
			mnemonic:    OpMnemonicCMP,
			args:        []uint8{indirectAddress},
		},
	}
}

// 0xD5: CMP $NN, X
func NewCMPZeroX(zeroAddress uint8) *CMP {
	return &CMP{
		baseOperation{
			code:        OpCodeCMPZeroX,
			addressMode: AddrModeZeroX,
			mnemonic:    OpMnemonicCMP,
			args:        []uint8{zeroAddress},
		},
	}
}

// 0xD9: CMP $NNNN, Y
func NewCMPAbsoluteY(absoluteAddress uint16) *CMP {
	return &CMP{
		baseOperation{
			code:        OpCodeCMPAbsoluteY,
			addressMode: AddrModeAbsoluteY,
			mnemonic:    OpMnemonicCMP,
			args:        util.BreakWordIntoBytes(absoluteAddress),
		},
	}
}

// 0xDD: CMP $NNNN, X
func NewCMPAbsoluteX(absoluteAddress uint16) *CMP {
	return &CMP{
		baseOperation{
			code:        OpCodeCMPAbsoluteX,
			addressMode: AddrModeAbsoluteX,
			mnemonic:    OpMnemonicCMP,
			args:        util.BreakWordIntoBytes(absoluteAddress),
		},
	}
}

func NewCMPBinary(opCode uint8, data io.Reader) (*CMP, error) {
	switch opCode {
	case OpCodeCMPIndirectX:
		var addr uint8

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewCMPIndirectX(addr), nil
	case OpCodeCMPZero:
		var addr uint8

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewCMPZero(addr), nil
	case OpCodeCMPImmediate:
		var addr uint8

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewCMPImmediate(addr), nil
	case OpCodeCMPAbsolute:
		var addr uint16

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewCMPAbsolute(addr), nil
	case OpCodeCMPIndirectY:
		var addr uint8

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewCMPIndirectY(addr), nil
	case OpCodeCMPZeroX:
		var addr uint8

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewCMPZeroX(addr), nil
	case OpCodeCMPAbsoluteY:
		var addr uint16

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewCMPAbsoluteY(addr), nil
	case OpCodeCMPAbsoluteX:
		var addr uint16

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewCMPAbsoluteX(addr), nil
	default:
		return nil, InvalidOpCodeError{
			OpCode: opCode,
		}
	}
}

func NewCMPFromBytes(addrMode uint8, arg0 uint8, arg1 uint8) (*CMP, error) {
	switch addrMode {
	case AddrModeIndirectX:
		return NewCMPIndirectX(arg0), nil
	case AddrModeZero, AddrModeRelative:
		return NewCMPZero(arg0), nil
	case AddrModeImmediate:
		return NewCMPImmediate(arg0), nil
	case AddrModeAbsolute:
		return NewCMPAbsolute(util.JoinBytesInWord([]uint8{arg0, arg1})), nil
	case AddrModeIndirectY:
		return NewCMPIndirectY(arg0), nil
	case AddrModeZeroX:
		return NewCMPZeroX(arg0), nil
	case AddrModeAbsoluteY:
		return NewCMPAbsoluteY(util.JoinBytesInWord([]uint8{arg0, arg1})), nil
	case AddrModeAbsoluteX:
		return NewCMPAbsoluteX(util.JoinBytesInWord([]uint8{arg0, arg1})), nil
	default:
		return nil, InvalidAddressModeError{
			AddressMode: addrMode,
		}
	}
}
