package cpu

import (
	"encoding/binary"
	"io"
	"log"

	"github.com/cd1/nes-emulator/util"
)

const (
	OpMnemonicLAX = "LAX"

	// unofficial opcodes
	OpCodeUnLAXIndirectX = 0xA3
	OpCodeUnLAXZero      = 0xA7
	OpCodeUnLAXAbsolute  = 0xAF
	OpCodeUnLAXIndirectY = 0xB3
	OpCodeUnLAXZeroY     = 0xB7
	OpCodeUnLAXAbsoluteY = 0xBF
)

func IsOpCodeValidLAX(opCode uint8) bool {
	return opCode == OpCodeUnLAXIndirectX ||
		opCode == OpCodeUnLAXZero ||
		opCode == OpCodeUnLAXAbsolute ||
		opCode == OpCodeUnLAXIndirectY ||
		opCode == OpCodeUnLAXZeroY ||
		opCode == OpCodeUnLAXAbsoluteY
}

func IsMnemonicValidLAX(mnemonic string) bool {
	return mnemonic == OpMnemonicLAX
}

type LAX struct {
	baseOperation
}

// 0xA3: LAX ($NN, X)
func NewUnLAXIndirectX(indirectAddress uint8) *LAX {
	return &LAX{
		baseOperation{
			code:        OpCodeUnLAXIndirectX,
			addressMode: AddrModeIndirectX,
			mnemonic:    OpMnemonicLAX,
			args:        []uint8{indirectAddress},
			unofficial:  true,
		},
	}
}

// 0xA7: LAX $NN
func NewUnLAXZero(zeroAddress uint8) *LAX {
	return &LAX{
		baseOperation{
			code:        OpCodeUnLAXZero,
			addressMode: AddrModeZero,
			mnemonic:    OpMnemonicLAX,
			args:        []uint8{zeroAddress},
			unofficial:  true,
		},
	}
}

// 0xAF: LAX $NNNN
func NewUnLAXAbsolute(absoluteAddress uint16) *LAX {
	return &LAX{
		baseOperation{
			code:        OpCodeUnLAXAbsolute,
			addressMode: AddrModeAbsolute,
			mnemonic:    OpMnemonicLAX,
			args:        util.BreakWordIntoBytes(absoluteAddress),
			unofficial:  true,
		},
	}
}

// 0xB3: LAX ($NN), Y
func NewUnLAXIndirectY(indirectAddress uint8) *LAX {
	return &LAX{
		baseOperation{
			code:        OpCodeUnLAXIndirectY,
			addressMode: AddrModeIndirectY,
			mnemonic:    OpMnemonicLAX,
			args:        []uint8{indirectAddress},
			unofficial:  true,
		},
	}
}

// 0xB7: LAX $NN, Y
func NewUnLAXZeroY(zeroAddress uint8) *LAX {
	return &LAX{
		baseOperation{
			code:        OpCodeUnLAXZeroY,
			addressMode: AddrModeZeroY,
			mnemonic:    OpMnemonicLAX,
			args:        []uint8{zeroAddress},
			unofficial:  true,
		},
	}
}

// 0xBF: LAX $NNNN, Y
func NewUnLAXAbsoluteY(absoluteAddress uint16) *LAX {
	return &LAX{
		baseOperation{
			code:        OpCodeUnLAXAbsoluteY,
			addressMode: AddrModeAbsoluteY,
			mnemonic:    OpMnemonicLAX,
			args:        util.BreakWordIntoBytes(absoluteAddress),
			unofficial:  true,
		},
	}
}

func NewLAXBinary(opCode uint8, data io.Reader) (*LAX, error) {
	switch opCode {
	case OpCodeUnLAXIndirectX:
		var addr uint8

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewUnLAXIndirectX(addr), nil
	case OpCodeUnLAXZero:
		var addr uint8

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewUnLAXZero(addr), nil
	case OpCodeUnLAXAbsolute:
		var addr uint16

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewUnLAXAbsolute(addr), nil
	case OpCodeUnLAXIndirectY:
		var addr uint8

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewUnLAXIndirectY(addr), nil
	case OpCodeUnLAXZeroY:
		var addr uint8

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewUnLAXZeroY(addr), nil
	case OpCodeUnLAXAbsoluteY:
		var addr uint16

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewUnLAXAbsoluteY(addr), nil
	default:
		return nil, InvalidOpCodeError{
			OpCode: opCode,
		}
	}
}

func NewLAXFromBytes(addrMode uint8, arg0 uint8, arg1 uint8) (*LAX, error) {
	switch addrMode {
	case AddrModeIndirectX:
		return NewUnLAXIndirectX(arg0), nil
	case AddrModeZero, AddrModeRelative:
		return NewUnLAXZero(arg0), nil
	case AddrModeAbsolute:
		return NewUnLAXAbsolute(util.JoinBytesInWord([]uint8{arg0, arg1})), nil
	case AddrModeIndirectY:
		return NewUnLAXIndirectY(arg0), nil
	case AddrModeZeroY:
		return NewUnLAXZeroY(arg0), nil
	case AddrModeAbsoluteY:
		return NewUnLAXAbsoluteY(util.JoinBytesInWord([]uint8{arg0, arg1})), nil
	default:
		return nil, InvalidAddressModeError{
			AddressMode: addrMode,
		}
	}
}

func (op LAX) Cycles() uint8 {
	switch op.AddressMode() {
	case AddrModeIndirectX:
		return 6
	case AddrModeZero:
		return 3
	case AddrModeAbsolute:
		return 4
	case AddrModeIndirectY:
		return 5
	case AddrModeZeroY:
		return 4
	case AddrModeAbsoluteY:
		return 4
	default:
		log.Printf("cannot calculate number of cycles for %#v: invalid address mode (%v)", op, op.AddressMode())
		return 0
	}
}

func (op LAX) ExecuteIn(env OperationEnvironment) (uint8, error) {
	cycles := op.Cycles()
	_, operand, pageCrossed := env.FetchOperand(op)

	if pageCrossed {
		cycles++
	}

	env.SetAccumulator(operand)
	env.SetIndexX(operand)

	env.SetStatusZero(operand == 0x00)
	env.SetStatusNegative(operand&0x80 != 0x00)

	env.IncrementProgramCounter(op.Size())

	return cycles, nil
}
