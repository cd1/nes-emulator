package cpu

import (
	"encoding/binary"
	"io"
	"log"

	"github.com/cd1/nes-emulator/util"
)

const (
	OpMnemonicRLA = "RLA"

	// unofficial opcodes
	OpCodeUnRLAIndirectX = 0x23
	OpCodeUnRLAZero      = 0x27
	OpCodeUnRLAAbsolute  = 0x2F
	OpCodeUnRLAIndirectY = 0x33
	OpCodeUnRLAZeroX     = 0x37
	OpCodeUnRLAAbsoluteY = 0x3B
	OpCodeUnRLAAbsoluteX = 0x3F
)

func IsOpCodeValidRLA(opCode uint8) bool {
	return opCode == OpCodeUnRLAIndirectX ||
		opCode == OpCodeUnRLAZero ||
		opCode == OpCodeUnRLAAbsolute ||
		opCode == OpCodeUnRLAIndirectY ||
		opCode == OpCodeUnRLAZeroX ||
		opCode == OpCodeUnRLAAbsoluteY ||
		opCode == OpCodeUnRLAAbsoluteX
}

func IsMnemonicValidRLA(mnemonic string) bool {
	return mnemonic == OpMnemonicRLA
}

type RLA struct {
	baseOperation
}

func NewUnRLAIndirectX(indirectAddress uint8) *RLA {
	return &RLA{
		baseOperation{
			code:        OpCodeUnRLAIndirectX,
			addressMode: AddrModeIndirectX,
			mnemonic:    OpMnemonicRLA,
			args:        []uint8{indirectAddress},
			unofficial:  true,
		},
	}
}

func NewUnRLAZero(zeroAddress uint8) *RLA {
	return &RLA{
		baseOperation{
			code:        OpCodeUnRLAZero,
			addressMode: AddrModeZero,
			mnemonic:    OpMnemonicRLA,
			args:        []uint8{zeroAddress},
			unofficial:  true,
		},
	}
}

func NewUnRLAAbsolute(absoluteAddress uint16) *RLA {
	return &RLA{
		baseOperation{
			code:        OpCodeUnRLAAbsolute,
			addressMode: AddrModeAbsolute,
			mnemonic:    OpMnemonicRLA,
			args:        util.BreakWordIntoBytes(absoluteAddress),
			unofficial:  true,
		},
	}
}

func NewUnRLAIndirectY(indirectAddress uint8) *RLA {
	return &RLA{
		baseOperation{
			code:        OpCodeUnRLAIndirectY,
			addressMode: AddrModeIndirectY,
			mnemonic:    OpMnemonicRLA,
			args:        []uint8{indirectAddress},
			unofficial:  true,
		},
	}
}

func NewUnRLAZeroX(zeroAddress uint8) *RLA {
	return &RLA{
		baseOperation{
			code:        OpCodeUnRLAZeroX,
			addressMode: AddrModeZeroX,
			mnemonic:    OpMnemonicRLA,
			args:        []uint8{zeroAddress},
			unofficial:  true,
		},
	}
}

func NewUnRLAAbsoluteY(absoluteAddress uint16) *RLA {
	return &RLA{
		baseOperation{
			code:        OpCodeUnRLAAbsoluteY,
			addressMode: AddrModeAbsoluteY,
			mnemonic:    OpMnemonicRLA,
			args:        util.BreakWordIntoBytes(absoluteAddress),
			unofficial:  true,
		},
	}
}

func NewUnRLAAbsoluteX(absoluteAddress uint16) *RLA {
	return &RLA{
		baseOperation{
			code:        OpCodeUnRLAAbsoluteX,
			addressMode: AddrModeAbsoluteX,
			mnemonic:    OpMnemonicRLA,
			args:        util.BreakWordIntoBytes(absoluteAddress),
			unofficial:  true,
		},
	}
}

func NewRLABinary(opCode uint8, data io.Reader) (*RLA, error) {
	switch opCode {
	case OpCodeUnRLAIndirectX:
		var addr uint8

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewUnRLAIndirectX(addr), nil
	case OpCodeUnRLAZero:
		var addr uint8

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewUnRLAZero(addr), nil
	case OpCodeUnRLAAbsolute:
		var addr uint16

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewUnRLAAbsolute(addr), nil
	case OpCodeUnRLAIndirectY:
		var addr uint8

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewUnRLAIndirectY(addr), nil
	case OpCodeUnRLAZeroX:
		var addr uint8

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewUnRLAZeroX(addr), nil
	case OpCodeUnRLAAbsoluteY:
		var addr uint16

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewUnRLAAbsoluteY(addr), nil
	case OpCodeUnRLAAbsoluteX:
		var addr uint16

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewUnRLAAbsoluteX(addr), nil
	default:
		return nil, InvalidOpCodeError{opCode}
	}
}

func NewRLAFromBytes(addrMode uint8, arg0 uint8, arg1 uint8) (*RLA, error) {
	switch addrMode {
	case AddrModeIndirectX:
		return NewUnRLAIndirectX(arg0), nil
	case AddrModeZero, AddrModeRelative:
		return NewUnRLAZero(arg0), nil
	case AddrModeAbsolute:
		return NewUnRLAAbsolute(util.JoinBytesInWord([]uint8{arg0, arg1})), nil
	case AddrModeIndirectY:
		return NewUnRLAIndirectY(arg0), nil
	case AddrModeZeroX:
		return NewUnRLAZeroX(arg0), nil
	case AddrModeAbsoluteY:
		return NewUnRLAAbsoluteY(util.JoinBytesInWord([]uint8{arg0, arg1})), nil
	case AddrModeAbsoluteX:
		return NewUnRLAAbsoluteX(util.JoinBytesInWord([]uint8{arg0, arg1})), nil
	default:
		return nil, InvalidAddressModeError{
			AddressMode: addrMode,
		}
	}
}

func (op RLA) Cycles() uint8 {
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

func (op RLA) ExecuteIn(env OperationEnvironment) (uint8, error) {
	cycles := op.Cycles()
	address, operand, _ := env.FetchOperand(op)

	result := operand << 1
	if env.IsStatusCarry() {
		result |= 0x01
	}
	env.WriteByte(address, result)

	result &= env.GetAccumulator()
	env.SetAccumulator(result)

	env.SetStatusCarry(operand&0x80 != 0x00)
	env.SetStatusZero(result == 0x00)
	env.SetStatusNegative(result&0x80 != 0x00)

	env.IncrementProgramCounter(op.Size())

	return cycles, nil
}
