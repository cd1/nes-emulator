package cpu

import (
	"encoding/binary"
	"io"
	"log"

	"github.com/cd1/nes-emulator/util"
)

const (
	OpMnemonicSRE = "SRE"

	// unofficial opcodes
	OpCodeUnSREIndirectX = 0x43
	OpCodeUnSREZero      = 0x47
	OpCodeUnSREAbsolute  = 0x4F
	OpCodeUnSREIndirectY = 0x53
	OpCodeUnSREZeroX     = 0x57
	OpCodeUnSREAbsoluteY = 0x5B
	OpCodeUnSREAbsoluteX = 0x5F
)

func IsOpCodeValidSRE(opCode uint8) bool {
	return opCode == OpCodeUnSREIndirectX ||
		opCode == OpCodeUnSREZero ||
		opCode == OpCodeUnSREAbsolute ||
		opCode == OpCodeUnSREIndirectY ||
		opCode == OpCodeUnSREZeroX ||
		opCode == OpCodeUnSREAbsoluteY ||
		opCode == OpCodeUnSREAbsoluteX
}

func IsMnemonicValidSRE(mnemonic string) bool {
	return mnemonic == OpMnemonicSRE
}

type SRE struct {
	baseOperation
}

// 0x43: SRE $NN, X
func NewUnSREIndirectX(indirectAddress uint8) *SRE {
	return &SRE{
		baseOperation{
			code:        OpCodeUnSREIndirectX,
			addressMode: AddrModeIndirectX,
			mnemonic:    OpMnemonicSRE,
			args:        []uint8{indirectAddress},
			unofficial:  true,
		},
	}
}

// 0x47: SRE $NN
func NewUnSREZero(zeroAddress uint8) *SRE {
	return &SRE{
		baseOperation{
			code:        OpCodeUnSREZero,
			addressMode: AddrModeZero,
			mnemonic:    OpMnemonicSRE,
			args:        []uint8{zeroAddress},
			unofficial:  true,
		},
	}
}

// 0x4F: SRE $NNNN
func NewUnSREAbsolute(absoluteAddress uint16) *SRE {
	return &SRE{
		baseOperation{
			code:        OpCodeUnSREAbsolute,
			addressMode: AddrModeAbsolute,
			mnemonic:    OpMnemonicSRE,
			args:        util.BreakWordIntoBytes(absoluteAddress),
			unofficial:  true,
		},
	}
}

// 0x53: SRE ($NN), Y
func NewUnSREIndirectY(indirectAddress uint8) *SRE {
	return &SRE{
		baseOperation{
			code:        OpCodeUnSREIndirectY,
			addressMode: AddrModeIndirectY,
			mnemonic:    OpMnemonicSRE,
			args:        []uint8{indirectAddress},
			unofficial:  true,
		},
	}
}

// 0x57: SRE $NN, X
func NewUnSREZeroX(zeroAddress uint8) *SRE {
	return &SRE{
		baseOperation{
			code:        OpCodeUnSREZeroX,
			addressMode: AddrModeZeroX,
			mnemonic:    OpMnemonicSRE,
			args:        []uint8{zeroAddress},
			unofficial:  true,
		},
	}
}

// 0x5B: SRE $NNNN, Y
func NewUnSREAbsoluteY(absoluteAddress uint16) *SRE {
	return &SRE{
		baseOperation{
			code:        OpCodeUnSREAbsoluteY,
			addressMode: AddrModeAbsoluteY,
			mnemonic:    OpMnemonicSRE,
			args:        util.BreakWordIntoBytes(absoluteAddress),
			unofficial:  true,
		},
	}
}

// 0x5F: SRE $NNNN, X
func NewUnSREAbsoluteX(absoluteAddress uint16) *SRE {
	return &SRE{
		baseOperation{
			code:        OpCodeUnSREAbsoluteX,
			addressMode: AddrModeAbsoluteX,
			mnemonic:    OpMnemonicSRE,
			args:        util.BreakWordIntoBytes(absoluteAddress),
			unofficial:  true,
		},
	}
}

func NewSREBinary(opCode uint8, data io.Reader) (*SRE, error) {
	switch opCode {
	case OpCodeUnSREIndirectX:
		var addr uint8

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewUnSREIndirectX(addr), nil
	case OpCodeUnSREZero:
		var addr uint8

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewUnSREZero(addr), nil
	case OpCodeUnSREAbsolute:
		var addr uint16

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewUnSREAbsolute(addr), nil
	case OpCodeUnSREIndirectY:
		var addr uint8

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewUnSREIndirectY(addr), nil
	case OpCodeUnSREZeroX:
		var addr uint8

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewUnSREZeroX(addr), nil
	case OpCodeUnSREAbsoluteY:
		var addr uint16

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewUnSREAbsoluteY(addr), nil
	case OpCodeUnSREAbsoluteX:
		var addr uint16

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewUnSREAbsoluteX(addr), nil
	default:
		return nil, InvalidOpCodeError{
			OpCode: opCode,
		}
	}
}

func NewSREFromBytes(addrMode uint8, arg0 uint8, arg1 uint8) (*SRE, error) {
	switch addrMode {
	case AddrModeIndirectX:
		return NewUnSREIndirectX(arg0), nil
	case AddrModeZero, AddrModeRelative:
		return NewUnSREZero(arg0), nil
	case AddrModeAbsolute:
		return NewUnSREAbsolute(util.JoinBytesInWord([]uint8{arg0, arg1})), nil
	case AddrModeIndirectY:
		return NewUnSREIndirectY(arg0), nil
	case AddrModeZeroX:
		return NewUnSREZeroX(arg0), nil
	case AddrModeAbsoluteY:
		return NewUnSREAbsoluteY(util.JoinBytesInWord([]uint8{arg0, arg1})), nil
	case AddrModeAbsoluteX:
		return NewUnSREAbsoluteX(util.JoinBytesInWord([]uint8{arg0, arg1})), nil
	default:
		return nil, InvalidAddressModeError{
			AddressMode: addrMode,
		}
	}
}

func (op SRE) Cycles() uint8 {
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

func (op SRE) ExecuteIn(env OperationEnvironment) (uint8, error) {
	cycles := op.Cycles()
	address, operand, _ := env.FetchOperand(op)

	result := operand >> 1
	env.WriteByte(address, result)

	result ^= env.GetAccumulator()
	env.SetAccumulator(result)

	env.SetStatusCarry(operand&0x01 != 0x00)
	env.SetStatusZero(result == 0x00)
	env.SetStatusNegative(result&0x80 != 0x00)

	env.IncrementProgramCounter(op.Size())

	return cycles, nil
}
