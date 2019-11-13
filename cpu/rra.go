package cpu

import (
	"encoding/binary"
	"io"
	"log"

	"github.com/cd1/nes-emulator/util"
)

const (
	OpMnemonicRRA = "RRA"

	// unofficial opcodes
	OpCodeUnRRAIndirectX = 0x63
	OpCodeUnRRAZero      = 0x67
	OpCodeUnRRAAbsolute  = 0x6F
	OpCodeUnRRAIndirectY = 0x73
	OpCodeUnRRAZeroX     = 0x77
	OpCodeUnRRAAbsoluteY = 0x7B
	OpCodeUnRRAAbsoluteX = 0x7F
)

func IsOpCodeValidRRA(opCode uint8) bool {
	return opCode == OpCodeUnRRAIndirectX ||
		opCode == OpCodeUnRRAZero ||
		opCode == OpCodeUnRRAAbsolute ||
		opCode == OpCodeUnRRAIndirectY ||
		opCode == OpCodeUnRRAZeroX ||
		opCode == OpCodeUnRRAAbsoluteY ||
		opCode == OpCodeUnRRAAbsoluteX
}

func IsMnemonicValidRRA(mnemonic string) bool {
	return mnemonic == OpMnemonicRRA
}

type RRA struct {
	baseOperation
}

// 0x63: RRA ($NN, X)
func NewUnRRAIndirectX(indirectAddress uint8) *RRA {
	return &RRA{
		baseOperation{
			code:        OpCodeUnRRAIndirectX,
			addressMode: AddrModeIndirectX,
			mnemonic:    OpMnemonicRRA,
			args:        []uint8{indirectAddress},
			unofficial:  true,
		},
	}
}

// 0x67: RRA $NN
func NewUnRRAZero(zeroAddress uint8) *RRA {
	return &RRA{
		baseOperation{
			code:        OpCodeUnRRAZero,
			addressMode: AddrModeZero,
			mnemonic:    OpMnemonicRRA,
			args:        []uint8{zeroAddress},
			unofficial:  true,
		},
	}
}

// 0x6F: RRA $NNNN
func NewUnRRAAbsolute(absoluteAddress uint16) *RRA {
	return &RRA{
		baseOperation{
			code:        OpCodeUnRRAAbsolute,
			addressMode: AddrModeAbsolute,
			mnemonic:    OpMnemonicRRA,
			args:        util.BreakWordIntoBytes(absoluteAddress),
			unofficial:  true,
		},
	}
}

// 0x73: RRA ($NN), Y
func NewUnRRAIndirectY(indirectAddress uint8) *RRA {
	return &RRA{
		baseOperation{
			code:        OpCodeUnRRAIndirectY,
			addressMode: AddrModeIndirectY,
			mnemonic:    OpMnemonicRRA,
			args:        []uint8{indirectAddress},
			unofficial:  true,
		},
	}
}

// 0x77: RRA $NN, X
func NewUnRRAZeroX(zeroAddress uint8) *RRA {
	return &RRA{
		baseOperation{
			code:        OpCodeUnRRAZeroX,
			addressMode: AddrModeZeroX,
			mnemonic:    OpMnemonicRRA,
			args:        []uint8{zeroAddress},
			unofficial:  true,
		},
	}
}

// 0x7B: RRA $NNNN, Y
func NewUnRRAAbsoluteY(absoluteAddress uint16) *RRA {
	return &RRA{
		baseOperation{
			code:        OpCodeUnRRAAbsoluteY,
			addressMode: AddrModeAbsoluteY,
			mnemonic:    OpMnemonicRRA,
			args:        util.BreakWordIntoBytes(absoluteAddress),
			unofficial:  true,
		},
	}
}

// 0x7F: RRA $NNNN, X
func NewUnRRAAbsoluteX(absoluteAddress uint16) *RRA {
	return &RRA{
		baseOperation{
			code:        OpCodeUnRRAAbsoluteX,
			addressMode: AddrModeAbsoluteX,
			mnemonic:    OpMnemonicRRA,
			args:        util.BreakWordIntoBytes(absoluteAddress),
			unofficial:  true,
		},
	}
}

func NewRRABinary(opCode uint8, data io.Reader) (*RRA, error) {
	switch opCode {
	case OpCodeUnRRAIndirectX:
		var addr uint8

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewUnRRAIndirectX(addr), nil
	case OpCodeUnRRAZero:
		var addr uint8

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewUnRRAZero(addr), nil
	case OpCodeUnRRAAbsolute:
		var addr uint16

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewUnRRAAbsolute(addr), nil
	case OpCodeUnRRAIndirectY:
		var addr uint8

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewUnRRAIndirectY(addr), nil
	case OpCodeUnRRAZeroX:
		var addr uint8

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewUnRRAZeroX(addr), nil
	case OpCodeUnRRAAbsoluteY:
		var addr uint16

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewUnRRAAbsoluteY(addr), nil
	case OpCodeUnRRAAbsoluteX:
		var addr uint16

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewUnRRAAbsoluteX(addr), nil
	default:
		return nil, InvalidOpCodeError{
			OpCode: opCode,
		}
	}
}

func NewRRAFromBytes(addrMode uint8, arg0 uint8, arg1 uint8) (*RRA, error) {
	switch addrMode {
	case AddrModeIndirectX:
		return NewUnRRAIndirectX(arg0), nil
	case AddrModeZero, AddrModeRelative:
		return NewUnRRAZero(arg0), nil
	case AddrModeAbsolute:
		return NewUnRRAAbsolute(util.JoinBytesInWord([]uint8{arg0, arg1})), nil
	case AddrModeIndirectY:
		return NewUnRRAIndirectY(arg0), nil
	case AddrModeZeroX:
		return NewUnRRAZeroX(arg0), nil
	case AddrModeAbsoluteY:
		return NewUnRRAAbsoluteY(util.JoinBytesInWord([]uint8{arg0, arg1})), nil
	case AddrModeAbsoluteX:
		return NewUnRRAAbsoluteX(util.JoinBytesInWord([]uint8{arg0, arg1})), nil
	default:
		return nil, InvalidAddressModeError{
			AddressMode: addrMode,
		}
	}
}

func (op RRA) Cycles() uint8 {
	switch op.AddressMode() {
	case AddrModeIndirectX:
		return 8
	case AddrModeZero:
		return 5
	case AddrModeAbsolute:
		return 6
	case AddrModeIndirectY:
		return 8
	case AddrModeZeroX:
		return 6
	case AddrModeAbsoluteY:
		return 7
	case AddrModeAbsoluteX:
		return 7
	default:
		log.Printf("cannot calculate number of cycles for %#v: invalid address mode (%v)", op, op.AddressMode())
		return 0
	}
}

func (op RRA) ExecuteIn(env OperationEnvironment) (uint8, error) {
	cycles := op.Cycles()
	address, operand, _ := env.FetchOperand(op)

	result := operand >> 1
	if env.IsStatusCarry() {
		result |= 0x80
	}
	env.WriteByte(address, result)

	tmpCarry := (operand&0x01 != 0x00)

	oldA := env.GetAccumulator()
	wordResult := uint16(oldA) + uint16(result)
	if tmpCarry {
		wordResult++
	}

	newA := uint8(wordResult)
	env.SetAccumulator(newA)
	env.SetStatusCarry(wordResult&0x0100 != 0x00)
	env.SetStatusZero(newA == 0x00)
	env.SetStatusOverflow((oldA&0x80 == 0x00 && result&0x80 == 0 && wordResult&0x80 != 0x00) ||
		(oldA&0x80 != 0x00 && result&0x80 != 0x00 && wordResult&0x80 == 0x00))
	env.SetStatusNegative(newA&0x80 != 0x00)

	env.IncrementProgramCounter(op.Size())

	return cycles, nil
}
