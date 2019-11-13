package cpu

import (
	"encoding/binary"
	"io"
	"log"

	"github.com/cd1/nes-emulator/util"
)

const (
	OpMnemonicSLO = "SLO"

	// unofficial opcodes
	OpCodeUnSLOIndirectX = 0x03
	OpCodeUnSLOZero      = 0x07
	OpCodeUnSLOAbsolute  = 0x0F
	OpCodeUnSLOIndirectY = 0x13
	OpCodeUnSLOZeroX     = 0x17
	OpCodeUnSLOAbsoluteY = 0x1B
	OpCodeUnSLOAbsoluteX = 0x1F
)

func IsOpCodeValidSLO(opCode uint8) bool {
	return opCode == OpCodeUnSLOIndirectX ||
		opCode == OpCodeUnSLOZero ||
		opCode == OpCodeUnSLOAbsolute ||
		opCode == OpCodeUnSLOIndirectY ||
		opCode == OpCodeUnSLOZeroX ||
		opCode == OpCodeUnSLOAbsoluteY ||
		opCode == OpCodeUnSLOAbsoluteX
}

func IsMnemonicValidSLO(mnemonic string) bool {
	return mnemonic == OpMnemonicSLO
}

type SLO struct {
	baseOperation
}

// 0x03: SLO ($NN, X)
func NewUnSLOIndirectX(indirectAddress uint8) *SLO {
	return &SLO{
		baseOperation{
			code:        OpCodeUnSLOIndirectX,
			addressMode: AddrModeIndirectX,
			mnemonic:    OpMnemonicSLO,
			args:        []uint8{indirectAddress},
			unofficial:  true,
		},
	}
}

// 0x07: SLO $NN
func NewUnSLOZero(zeroAddress uint8) *SLO {
	return &SLO{
		baseOperation{
			code:        OpCodeUnSLOZero,
			addressMode: AddrModeZero,
			mnemonic:    OpMnemonicSLO,
			args:        []uint8{zeroAddress},
			unofficial:  true,
		},
	}
}

// 0x0F: SLO $NNNN
func NewUnSLOAbsolute(absoluteAddress uint16) *SLO {
	return &SLO{
		baseOperation{
			code:        OpCodeUnSLOAbsolute,
			addressMode: AddrModeAbsolute,
			mnemonic:    OpMnemonicSLO,
			args:        util.BreakWordIntoBytes(absoluteAddress),
			unofficial:  true,
		},
	}
}

// 0x13: SLO ($NN), Y
func NewUnSLOIndirectY(indirectAddress uint8) *SLO {
	return &SLO{
		baseOperation{
			code:        OpCodeUnSLOIndirectY,
			addressMode: AddrModeIndirectY,
			mnemonic:    OpMnemonicSLO,
			args:        []uint8{indirectAddress},
			unofficial:  true,
		},
	}
}

// 0x17: SLO $NN
func NewUnSLOZeroX(zeroAddress uint8) *SLO {
	return &SLO{
		baseOperation{
			code:        OpCodeUnSLOZeroX,
			addressMode: AddrModeZeroX,
			mnemonic:    OpMnemonicSLO,
			args:        []uint8{zeroAddress},
			unofficial:  true,
		},
	}
}

// 0x1B: SLO $NNNN, Y
func NewUnSLOAbsoluteY(absoluteAddress uint16) *SLO {
	return &SLO{
		baseOperation{
			code:        OpCodeUnSLOAbsoluteY,
			addressMode: AddrModeAbsoluteY,
			mnemonic:    OpMnemonicSLO,
			args:        util.BreakWordIntoBytes(absoluteAddress),
			unofficial:  true,
		},
	}
}

// 0x1F: SLO $NNNN, X
func NewUnSLOAbsoluteX(absoluteAddress uint16) *SLO {
	return &SLO{
		baseOperation{
			code:        OpCodeUnSLOAbsoluteX,
			addressMode: AddrModeAbsoluteX,
			mnemonic:    OpMnemonicSLO,
			args:        util.BreakWordIntoBytes(absoluteAddress),
			unofficial:  true,
		},
	}
}

func NewSLOBinary(opCode uint8, data io.Reader) (*SLO, error) {
	switch opCode {
	case OpCodeUnSLOIndirectX:
		var addr uint8

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewUnSLOIndirectX(addr), nil
	case OpCodeUnSLOZero:
		var addr uint8

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewUnSLOZero(addr), nil
	case OpCodeUnSLOAbsolute:
		var addr uint16

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewUnSLOAbsolute(addr), nil
	case OpCodeUnSLOIndirectY:
		var addr uint8

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewUnSLOIndirectY(addr), nil
	case OpCodeUnSLOZeroX:
		var addr uint8

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewUnSLOZeroX(addr), nil
	case OpCodeUnSLOAbsoluteY:
		var addr uint16

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewUnSLOAbsoluteY(addr), nil
	case OpCodeUnSLOAbsoluteX:
		var addr uint16

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewUnSLOAbsoluteX(addr), nil
	default:
		return nil, InvalidOpCodeError{
			OpCode: opCode,
		}
	}
}

func NewSLOFromBytes(addrMode uint8, arg0 uint8, arg1 uint8) (*SLO, error) {
	switch addrMode {
	case AddrModeIndirectX:
		return NewUnSLOIndirectX(arg0), nil
	case AddrModeZero, AddrModeRelative:
		return NewUnSLOZero(arg0), nil
	case AddrModeAbsolute:
		return NewUnSLOAbsolute(util.JoinBytesInWord([]uint8{arg0, arg1})), nil
	case AddrModeIndirectY:
		return NewUnSLOIndirectY(arg0), nil
	case AddrModeZeroX:
		return NewUnSLOZeroX(arg0), nil
	case AddrModeAbsoluteY:
		return NewUnSLOAbsoluteY(util.JoinBytesInWord([]uint8{arg0, arg1})), nil
	case AddrModeAbsoluteX:
		return NewUnSLOAbsoluteX(util.JoinBytesInWord([]uint8{arg0, arg1})), nil
	default:
		return nil, InvalidAddressModeError{
			AddressMode: addrMode,
		}
	}
}

func (op SLO) Cycles() uint8 {
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

func (op SLO) ExecuteIn(env OperationEnvironment) (uint8, error) {
	cycles := op.Cycles()
	address, operand, _ := env.FetchOperand(op)

	result := operand << 1
	env.WriteByte(address, result)

	result |= env.GetAccumulator()
	env.SetAccumulator(result)

	env.SetStatusCarry(operand&0x80 != 0x00)
	env.SetStatusZero(result == 0x00)
	env.SetStatusNegative(result&0x80 != 0x00)

	env.IncrementProgramCounter(op.Size())

	return cycles, nil
}
