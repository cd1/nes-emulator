package cpu

import (
	"encoding/binary"
	"io"
	"log"

	"github.com/cd1/nes-emulator/util"
)

const (
	OpMnemonicORA = "ORA"

	OpCodeORAIndirectX = 0x01
	OpCodeORAZero      = 0x05
	OpCodeORAImmediate = 0x09
	OpCodeORAAbsolute  = 0x0D
	OpCodeORAIndirectY = 0x11
	OpCodeORAZeroX     = 0x15
	OpCodeORAAbsoluteY = 0x19
	OpCodeORAAbsoluteX = 0x1D
)

func IsOpCodeValidORA(opCode uint8) bool {
	return opCode == OpCodeORAIndirectX ||
		opCode == OpCodeORAZero ||
		opCode == OpCodeORAImmediate ||
		opCode == OpCodeORAAbsolute ||
		opCode == OpCodeORAIndirectY ||
		opCode == OpCodeORAZeroX ||
		opCode == OpCodeORAAbsoluteY ||
		opCode == OpCodeORAAbsoluteX
}

func IsMnemonicValidORA(mnemonic string) bool {
	return mnemonic == OpMnemonicORA
}

type ORA struct {
	baseOperation
}

// 0x01: ORA ($NN, X)
func NewORAIndirectX(indirectAddress uint8) *ORA {
	return &ORA{
		baseOperation{
			code:        OpCodeORAIndirectX,
			addressMode: AddrModeIndirectX,
			mnemonic:    OpMnemonicORA,
			args:        []uint8{indirectAddress},
		},
	}
}

// 0x05: ORA $NN
func NewORAZero(zeroAddress uint8) *ORA {
	return &ORA{
		baseOperation{
			code:        OpCodeORAZero,
			addressMode: AddrModeZero,
			mnemonic:    OpMnemonicORA,
			args:        []uint8{zeroAddress},
		},
	}
}

// 0x09: ORA #$NN
func NewORAImmediate(value uint8) *ORA {
	return &ORA{
		baseOperation{
			code:        OpCodeORAImmediate,
			addressMode: AddrModeImmediate,
			mnemonic:    OpMnemonicORA,
			args:        []uint8{value},
		},
	}
}

// 0x0D: ORA $NNNN
func NewORAAbsolute(absoluteAddress uint16) *ORA {
	return &ORA{
		baseOperation{
			code:        OpCodeORAAbsolute,
			addressMode: AddrModeAbsolute,
			mnemonic:    OpMnemonicORA,
			args:        util.BreakWordIntoBytes(absoluteAddress),
		},
	}
}

// 0x11: ORA ($NN), Y
func NewORAIndirectY(indirectAddress uint8) *ORA {
	return &ORA{
		baseOperation{
			code:        OpCodeORAIndirectY,
			addressMode: AddrModeIndirectY,
			mnemonic:    OpMnemonicORA,
			args:        []uint8{indirectAddress},
		},
	}
}

// 0x15: ORA $NN
func NewORAZeroX(zeroAddress uint8) *ORA {
	return &ORA{
		baseOperation{
			code:        OpCodeORAZeroX,
			addressMode: AddrModeZeroX,
			mnemonic:    OpMnemonicORA,
			args:        []uint8{zeroAddress},
		},
	}
}

// 0x19: ORA $NNNN, Y
func NewORAAbsoluteY(absoluteAddress uint16) *ORA {
	return &ORA{
		baseOperation{
			code:        OpCodeORAAbsoluteY,
			addressMode: AddrModeAbsoluteY,
			mnemonic:    OpMnemonicORA,
			args:        util.BreakWordIntoBytes(absoluteAddress),
		},
	}
}

// 0x1D: ORA $NNNN, X
func NewORAAbsoluteX(absoluteAddress uint16) *ORA {
	return &ORA{
		baseOperation{
			code:        OpCodeORAAbsoluteX,
			addressMode: AddrModeAbsoluteX,
			mnemonic:    OpMnemonicORA,
			args:        util.BreakWordIntoBytes(absoluteAddress),
		},
	}
}

func NewORABinary(opCode uint8, data io.Reader) (*ORA, error) {
	switch opCode {
	case OpCodeORAIndirectX:
		var addr uint8

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewORAIndirectX(addr), nil
	case OpCodeORAZero:
		var addr uint8

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewORAZero(addr), nil
	case OpCodeORAImmediate:
		var addr uint8

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewORAImmediate(addr), nil
	case OpCodeORAAbsolute:
		var addr uint16

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewORAAbsolute(addr), nil
	case OpCodeORAIndirectY:
		var addr uint8

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewORAIndirectY(addr), nil
	case OpCodeORAZeroX:
		var addr uint8

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewORAZeroX(addr), nil
	case OpCodeORAAbsoluteY:
		var addr uint16

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewORAAbsoluteY(addr), nil
	case OpCodeORAAbsoluteX:
		var addr uint16

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewORAAbsoluteX(addr), nil
	default:
		return nil, InvalidOpCodeError{
			OpCode: opCode,
		}
	}
}

func NewORAFromBytes(addrMode uint8, arg0 uint8, arg1 uint8) (*ORA, error) {
	switch addrMode {
	case AddrModeIndirectX:
		return NewORAIndirectX(arg0), nil
	case AddrModeZero, AddrModeRelative:
		return NewORAZero(arg0), nil
	case AddrModeImmediate:
		return NewORAImmediate(arg0), nil
	case AddrModeAbsolute:
		return NewORAAbsolute(util.JoinBytesInWord([]uint8{arg0, arg1})), nil
	case AddrModeIndirectY:
		return NewORAIndirectY(arg0), nil
	case AddrModeZeroX:
		return NewORAZeroX(arg0), nil
	case AddrModeAbsoluteY:
		return NewORAAbsoluteY(util.JoinBytesInWord([]uint8{arg0, arg1})), nil
	case AddrModeAbsoluteX:
		return NewORAAbsoluteX(util.JoinBytesInWord([]uint8{arg0, arg1})), nil
	default:
		return nil, InvalidAddressModeError{
			AddressMode: addrMode,
		}
	}
}

func (op ORA) Cycles() uint8 {
	switch op.AddressMode() {
	case AddrModeIndirectX:
		return 6
	case AddrModeZero:
		return 3
	case AddrModeImmediate:
		return 2
	case AddrModeAbsolute:
		return 4
	case AddrModeIndirectY:
		return 5
	case AddrModeZeroX:
		return 4
	case AddrModeAbsoluteY:
		return 4
	case AddrModeAbsoluteX:
		return 4
	default:
		log.Printf("cannot calculate number of cycles for %#v: invalid address mode (%v)", op, op.AddressMode())
		return 0
	}
}

func (op ORA) ExecuteIn(env OperationEnvironment) (uint8, error) {
	cycles := op.Cycles()
	_, operand, pageCrossed := env.FetchOperand(op)

	if pageCrossed {
		cycles++
	}

	result := env.GetAccumulator() | operand
	env.SetAccumulator(result)

	env.SetStatusZero(result == 0x00)
	env.SetStatusNegative(result&0x80 != 0x00)

	env.IncrementProgramCounter(op.Size())

	return cycles, nil
}
