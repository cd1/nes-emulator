package cpu

import (
	"encoding/binary"
	"io"
	"log"

	"github.com/cd1/nes-emulator/util"
)

const (
	OpMnemonicDCP = "DCP"

	// unofficial opcodes
	OpCodeUnDCPIndirectX = 0xC3
	OpCodeUnDCPZero      = 0xC7
	OpCodeUnDCPAbsolute  = 0xCF
	OpCodeUnDCPIndirectY = 0xD3
	OpCodeUnDCPZeroX     = 0xD7
	OpCodeUnDCPAbsoluteY = 0xDB
	OpCodeUnDCPAbsoluteX = 0xDF
)

func IsOpCodeValidDCP(opCode uint8) bool {
	return opCode == OpCodeUnDCPIndirectX ||
		opCode == OpCodeUnDCPZero ||
		opCode == OpCodeUnDCPAbsolute ||
		opCode == OpCodeUnDCPIndirectY ||
		opCode == OpCodeUnDCPZeroX ||
		opCode == OpCodeUnDCPAbsoluteY ||
		opCode == OpCodeUnDCPAbsoluteX
}

func IsMnemonicValidDCP(mnemonic string) bool {
	return mnemonic == OpMnemonicDCP
}

type DCP struct {
	baseOperation
}

// 0xC3: DCP ($NN, X)
func NewUnDCPIndirectX(indirectAddress uint8) *DCP {
	return &DCP{
		baseOperation{
			code:        OpCodeUnDCPIndirectX,
			addressMode: AddrModeIndirectX,
			mnemonic:    OpMnemonicDCP,
			args:        []uint8{indirectAddress},
			unofficial:  true,
		},
	}
}

// 0xC7: DCP $NN
func NewUnDCPZero(zeroAddress uint8) *DCP {
	return &DCP{
		baseOperation{
			code:        OpCodeUnDCPZero,
			addressMode: AddrModeZero,
			mnemonic:    OpMnemonicDCP,
			args:        []uint8{zeroAddress},
			unofficial:  true,
		},
	}
}

// 0xCF: DCP $NNNN
func NewUnDCPAbsolute(absoluteAddress uint16) *DCP {
	return &DCP{
		baseOperation{
			code:        OpCodeUnDCPAbsolute,
			addressMode: AddrModeAbsolute,
			mnemonic:    OpMnemonicDCP,
			args:        util.BreakWordIntoBytes(absoluteAddress),
			unofficial:  true,
		},
	}
}

// 0xD3: DCP ($NN), Y
func NewUnDCPIndirectY(indirectAddress uint8) *DCP {
	return &DCP{
		baseOperation{
			code:        OpCodeUnDCPIndirectY,
			addressMode: AddrModeIndirectY,
			mnemonic:    OpMnemonicDCP,
			args:        []uint8{indirectAddress},
			unofficial:  true,
		},
	}
}

// 0xD7: DCP $NN, X
func NewUnDCPZeroX(zeroAddress uint8) *DCP {
	return &DCP{
		baseOperation{
			code:        OpCodeUnDCPZeroX,
			addressMode: AddrModeZeroX,
			mnemonic:    OpMnemonicDCP,
			args:        []uint8{zeroAddress},
			unofficial:  true,
		},
	}
}

// 0xDB: DCP $NNNN, Y
func NewUnDCPAbsoluteY(absoluteAddress uint16) *DCP {
	return &DCP{
		baseOperation{
			code:        OpCodeUnDCPAbsoluteY,
			addressMode: AddrModeAbsoluteY,
			mnemonic:    OpMnemonicDCP,
			args:        util.BreakWordIntoBytes(absoluteAddress),
			unofficial:  true,
		},
	}
}

// 0xDF: DCP $NNNN, X
func NewUnDCPAbsoluteX(absoluteAddress uint16) *DCP {
	return &DCP{
		baseOperation{
			code:        OpCodeUnDCPAbsoluteX,
			addressMode: AddrModeAbsoluteX,
			mnemonic:    OpMnemonicDCP,
			args:        util.BreakWordIntoBytes(absoluteAddress),
			unofficial:  true,
		},
	}
}

func NewDCPBinary(opCode uint8, data io.Reader) (*DCP, error) {
	switch opCode {
	case OpCodeUnDCPIndirectX:
		var addr uint8

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewUnDCPIndirectX(addr), nil
	case OpCodeUnDCPZero:
		var addr uint8

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewUnDCPZero(addr), nil
	case OpCodeUnDCPAbsolute:
		var addr uint16

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewUnDCPAbsolute(addr), nil
	case OpCodeUnDCPIndirectY:
		var addr uint8

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewUnDCPIndirectY(addr), nil
	case OpCodeUnDCPZeroX:
		var addr uint8

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewUnDCPZeroX(addr), nil
	case OpCodeUnDCPAbsoluteY:
		var addr uint16

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewUnDCPAbsoluteY(addr), nil
	case OpCodeUnDCPAbsoluteX:
		var addr uint16

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewUnDCPAbsoluteX(addr), nil
	default:
		return nil, InvalidOpCodeError{
			OpCode: opCode,
		}
	}
}

func NewDCPFromBytes(addrMode uint8, arg0 uint8, arg1 uint8) (*DCP, error) {
	switch addrMode {
	case AddrModeIndirectX:
		return NewUnDCPIndirectX(arg0), nil
	case AddrModeZero, AddrModeRelative:
		return NewUnDCPZero(arg0), nil
	case AddrModeAbsolute:
		return NewUnDCPAbsolute(util.JoinBytesInWord([]uint8{arg0, arg1})), nil
	case AddrModeIndirectY:
		return NewUnDCPIndirectY(arg0), nil
	case AddrModeZeroX:
		return NewUnDCPZeroX(arg0), nil
	case AddrModeAbsoluteY:
		return NewUnDCPAbsoluteY(util.JoinBytesInWord([]uint8{arg0, arg1})), nil
	case AddrModeAbsoluteX:
		return NewUnDCPAbsoluteX(util.JoinBytesInWord([]uint8{arg0, arg1})), nil
	default:
		return nil, InvalidAddressModeError{
			AddressMode: addrMode,
		}
	}
}

func (op DCP) Cycles() uint8 {
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

func (op DCP) ExecuteIn(env OperationEnvironment) (uint8, error) {
	cycles := op.Cycles()
	address, operand, _ := env.FetchOperand(op)

	newMemValue := operand - 1
	env.WriteByte(address, newMemValue)

	a := env.GetAccumulator()

	env.SetStatusCarry(a >= newMemValue)
	env.SetStatusZero(a-newMemValue == 0x00)
	env.SetStatusNegative((a-newMemValue)&0x80 != 0x00)

	env.IncrementProgramCounter(op.Size())

	return cycles, nil
}
