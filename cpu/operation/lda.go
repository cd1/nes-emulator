package operation

import (
	"encoding/binary"
	"io"

	"github.com/cd1/nes-emulator/util"
)

const (
	OpMnemonicLDA = "LDA"

	OpCodeLDAIndirectX = 0xA1
	OpCodeLDAZero      = 0xA5
	OpCodeLDAImmediate = 0xA9
	OpCodeLDAAbsolute  = 0xAD
	OpCodeLDAIndirectY = 0xB1
	OpCodeLDAZeroX     = 0xB5
	OpCodeLDAAbsoluteY = 0xB9
	OpCodeLDAAbsoluteX = 0xBD
)

func IsOpCodeValidLDA(opCode uint8) bool {
	return opCode == OpCodeLDAIndirectX ||
		opCode == OpCodeLDAZero ||
		opCode == OpCodeLDAImmediate ||
		opCode == OpCodeLDAAbsolute ||
		opCode == OpCodeLDAIndirectY ||
		opCode == OpCodeLDAZeroX ||
		opCode == OpCodeLDAAbsoluteY ||
		opCode == OpCodeLDAAbsoluteX
}

func IsMnemonicValidLDA(mnemonic string) bool {
	return mnemonic == OpMnemonicLDA
}

type LDA struct {
	baseOperation
}

// 0xA1: LDA ($NN, X)
func NewLDAIndirectX(indirectAddress uint8) *LDA {
	return &LDA{
		baseOperation{
			code:        OpCodeLDAIndirectX,
			addressMode: AddrModeIndirectX,
			mnemonic:    OpMnemonicLDA,
			args:        []uint8{indirectAddress},
		},
	}
}

// 0xA5: LDA $NN
func NewLDAZero(zeroAddress uint8) *LDA {
	return &LDA{
		baseOperation{
			code:        OpCodeLDAZero,
			addressMode: AddrModeZero,
			mnemonic:    OpMnemonicLDA,
			args:        []uint8{zeroAddress},
		},
	}
}

// 0xA9: LDA #$NN
func NewLDAImmediate(value uint8) *LDA {
	return &LDA{
		baseOperation{
			code:        OpCodeLDAImmediate,
			addressMode: AddrModeImmediate,
			mnemonic:    OpMnemonicLDA,
			args:        []uint8{value},
		},
	}
}

// 0xAD: LDA $NNNN
func NewLDAAbsolute(absoluteAddress uint16) *LDA {
	return &LDA{
		baseOperation{
			code:        OpCodeLDAAbsolute,
			addressMode: AddrModeAbsolute,
			mnemonic:    OpMnemonicLDA,
			args:        util.BreakWordIntoBytes(absoluteAddress),
		},
	}
}

// 0xB1: LDA ($NN), Y
func NewLDAIndirectY(indirectAddress uint8) *LDA {
	return &LDA{
		baseOperation{
			code:        OpCodeLDAIndirectY,
			addressMode: AddrModeIndirectY,
			mnemonic:    OpMnemonicLDA,
			args:        []uint8{indirectAddress},
		},
	}
}

// 0xB5: LDA $NN, X
func NewLDAZeroX(zeroAddress uint8) *LDA {
	return &LDA{
		baseOperation{
			code:        OpCodeLDAZeroX,
			addressMode: AddrModeZeroX,
			mnemonic:    OpMnemonicLDA,
			args:        []uint8{zeroAddress},
		},
	}
}

// 0xB9: LDA $NNNN, Y
func NewLDAAbsoluteY(absoluteAddress uint16) *LDA {
	return &LDA{
		baseOperation{
			code:        OpCodeLDAAbsoluteY,
			addressMode: AddrModeAbsoluteY,
			mnemonic:    OpMnemonicLDA,
			args:        util.BreakWordIntoBytes(absoluteAddress),
		},
	}
}

// 0xBD: LDA $NNNN, X
func NewLDAAbsoluteX(absoluteAddress uint16) *LDA {
	return &LDA{
		baseOperation{
			code:        OpCodeLDAAbsoluteX,
			addressMode: AddrModeAbsoluteX,
			mnemonic:    OpMnemonicLDA,
			args:        util.BreakWordIntoBytes(absoluteAddress),
		},
	}
}

func NewLDABinary(opCode uint8, data io.Reader) (*LDA, error) {
	switch opCode {
	case OpCodeLDAIndirectX:
		var addr uint8

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewLDAIndirectX(addr), nil
	case OpCodeLDAZero:
		var addr uint8

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewLDAZero(addr), nil
	case OpCodeLDAImmediate:
		var addr uint8

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewLDAImmediate(addr), nil
	case OpCodeLDAAbsolute:
		var addr uint16

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewLDAAbsolute(addr), nil
	case OpCodeLDAIndirectY:
		var addr uint8

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewLDAIndirectY(addr), nil
	case OpCodeLDAZeroX:
		var addr uint8

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewLDAZeroX(addr), nil
	case OpCodeLDAAbsoluteY:
		var addr uint16

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewLDAAbsoluteY(addr), nil
	case OpCodeLDAAbsoluteX:
		var addr uint16

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewLDAAbsoluteX(addr), nil
	default:
		return nil, InvalidOpCodeError{
			OpCode: opCode,
		}
	}
}

func NewLDAFromBytes(addrMode uint8, arg0 uint8, arg1 uint8) (*LDA, error) {
	switch addrMode {
	case AddrModeIndirectX:
		return NewLDAIndirectX(arg0), nil
	case AddrModeZero, AddrModeRelative:
		return NewLDAZero(arg0), nil
	case AddrModeImmediate:
		return NewLDAImmediate(arg0), nil
	case AddrModeAbsolute:
		return NewLDAAbsolute(util.JoinBytesInWord([]uint8{arg0, arg1})), nil
	case AddrModeIndirectY:
		return NewLDAIndirectY(arg0), nil
	case AddrModeZeroX:
		return NewLDAZeroX(arg0), nil
	case AddrModeAbsoluteY:
		return NewLDAAbsoluteY(util.JoinBytesInWord([]uint8{arg0, arg1})), nil
	case AddrModeAbsoluteX:
		return NewLDAAbsoluteX(util.JoinBytesInWord([]uint8{arg0, arg1})), nil
	default:
		return nil, InvalidAddressModeError{
			AddressMode: addrMode,
		}
	}
}
