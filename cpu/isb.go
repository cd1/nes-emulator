package cpu

import (
	"encoding/binary"
	"io"
	"log"

	"github.com/cd1/nes-emulator/util"
)

const (
	OpMnemonicISB = "ISB"

	// unofficial opcodes
	OpCodeUnISBIndirectX = 0xE3
	OpCodeUnISBZero      = 0xE7
	OpCodeUnISBAbsolute  = 0xEF
	OpCodeUnISBIndirectY = 0xF3
	OpCodeUnISBZeroX     = 0xF7
	OpCodeUnISBAbsoluteY = 0xFB
	OpCodeUnISBAbsoluteX = 0xFF
)

func IsOpCodeValidISB(opCode uint8) bool {
	return opCode == OpCodeUnISBIndirectX ||
		opCode == OpCodeUnISBZero ||
		opCode == OpCodeUnISBAbsolute ||
		opCode == OpCodeUnISBIndirectY ||
		opCode == OpCodeUnISBZeroX ||
		opCode == OpCodeUnISBAbsoluteY ||
		opCode == OpCodeUnISBAbsoluteX
}

func IsMnemonicValidISB(mnemonic string) bool {
	return mnemonic == OpMnemonicISB
}

type ISB struct {
	baseOperation
}

// 0xE3: ISB ($NN, X)
func NewUnISBIndirectX(indirectAddress uint8) *ISB {
	return &ISB{
		baseOperation{
			code:        OpCodeUnISBIndirectX,
			addressMode: AddrModeIndirectX,
			mnemonic:    OpMnemonicISB,
			args:        []uint8{indirectAddress},
			unofficial:  true,
		},
	}
}

// 0xE7: ISB $NN
func NewUnISBZero(zeroAddress uint8) *ISB {
	return &ISB{
		baseOperation{
			code:        OpCodeUnISBZero,
			addressMode: AddrModeZero,
			mnemonic:    OpMnemonicISB,
			args:        []uint8{zeroAddress},
			unofficial:  true,
		},
	}
}

// 0xEF: ISB $NNNN
func NewUnISBAbsolute(absoluteAddress uint16) *ISB {
	return &ISB{
		baseOperation{
			code:        OpCodeUnISBAbsolute,
			addressMode: AddrModeAbsolute,
			mnemonic:    OpMnemonicISB,
			args:        util.BreakWordIntoBytes(absoluteAddress),
			unofficial:  true,
		},
	}
}

// 0xF3: ISB ($NN), Y
func NewUnISBIndirectY(indirectAddress uint8) *ISB {
	return &ISB{
		baseOperation{
			code:        OpCodeUnISBIndirectY,
			addressMode: AddrModeIndirectY,
			mnemonic:    OpMnemonicISB,
			args:        []uint8{indirectAddress},
			unofficial:  true,
		},
	}
}

// 0xF7: ISB $NN, X
func NewUnISBZeroX(zeroAddress uint8) *ISB {
	return &ISB{
		baseOperation{
			code:        OpCodeUnISBZeroX,
			addressMode: AddrModeZeroX,
			mnemonic:    OpMnemonicISB,
			args:        []uint8{zeroAddress},
			unofficial:  true,
		},
	}
}

// 0xFB: ISB $NNNN, Y
func NewUnISBAbsoluteY(absoluteAddress uint16) *ISB {
	return &ISB{
		baseOperation{
			code:        OpCodeUnISBAbsoluteY,
			addressMode: AddrModeAbsoluteY,
			mnemonic:    OpMnemonicISB,
			args:        util.BreakWordIntoBytes(absoluteAddress),
			unofficial:  true,
		},
	}
}

// 0xFF: ISB $NNNN, X
func NewUnISBAbsoluteX(absoluteAddress uint16) *ISB {
	return &ISB{
		baseOperation{
			code:        OpCodeUnISBAbsoluteX,
			addressMode: AddrModeAbsoluteX,
			mnemonic:    OpMnemonicISB,
			args:        util.BreakWordIntoBytes(absoluteAddress),
			unofficial:  true,
		},
	}
}

func NewISBBinary(opCode uint8, data io.Reader) (*ISB, error) {
	switch opCode {
	case OpCodeUnISBIndirectX:
		var addr uint8

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewUnISBIndirectX(addr), nil
	case OpCodeUnISBZero:
		var addr uint8

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewUnISBZero(addr), nil
	case OpCodeUnISBAbsolute:
		var addr uint16

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewUnISBAbsolute(addr), nil
	case OpCodeUnISBIndirectY:
		var addr uint8

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewUnISBIndirectY(addr), nil
	case OpCodeUnISBZeroX:
		var addr uint8

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewUnISBZeroX(addr), nil
	case OpCodeUnISBAbsoluteY:
		var addr uint16

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewUnISBAbsoluteY(addr), nil
	case OpCodeUnISBAbsoluteX:
		var addr uint16

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewUnISBAbsoluteX(addr), nil
	default:
		return nil, InvalidOpCodeError{
			OpCode: opCode,
		}
	}
}

func NewISBFromBytes(addrMode uint8, arg0 uint8, arg1 uint8) (*ISB, error) {
	switch addrMode {
	case AddrModeIndirectX:
		return NewUnISBIndirectX(arg0), nil
	case AddrModeZero, AddrModeRelative:
		return NewUnISBZero(arg0), nil
	case AddrModeAbsolute:
		return NewUnISBAbsolute(util.JoinBytesInWord([]uint8{arg0, arg1})), nil
	case AddrModeIndirectY:
		return NewUnISBIndirectY(arg0), nil
	case AddrModeZeroX:
		return NewUnISBZeroX(arg0), nil
	case AddrModeAbsoluteY:
		return NewUnISBAbsoluteY(util.JoinBytesInWord([]uint8{arg0, arg1})), nil
	case AddrModeAbsoluteX:
		return NewUnISBAbsoluteX(util.JoinBytesInWord([]uint8{arg0, arg1})), nil
	default:
		return nil, InvalidAddressModeError{
			AddressMode: addrMode,
		}
	}
}

func (op ISB) Cycles() uint8 {
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

func (op ISB) ExecuteIn(env OperationEnvironment) (uint8, error) {
	cycles := op.Cycles()
	address, operand, _ := env.FetchOperand(op)

	newMemValue := operand + 1
	env.WriteByte(address, newMemValue)

	oldA := env.GetAccumulator()
	result := oldA - newMemValue
	signedResult := int16(int8(oldA)) - int16(int8(newMemValue))
	if !env.IsStatusCarry() {
		result--
		signedResult--
	}

	env.SetAccumulator(result)
	env.SetStatusCarry(result <= oldA)
	env.SetStatusZero(result == 0x00)
	env.SetStatusOverflow(signedResult > 127 || signedResult < -128)
	env.SetStatusNegative(result&0x80 != 0x00)

	env.IncrementProgramCounter(op.Size())

	return cycles, nil
}
