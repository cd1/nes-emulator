package cpu

import (
	"encoding/binary"
	"io"
	"log"

	"github.com/cd1/nes-emulator/util"
)

const (
	OpMnemonicSBC = "SBC"

	OpCodeSBCIndirectX = 0xE1
	OpCodeSBCZero      = 0xE5
	OpCodeSBCImmediate = 0xE9
	OpCodeSBCAbsolute  = 0xED
	OpCodeSBCIndirectY = 0xF1
	OpCodeSBCZeroX     = 0xF5
	OpCodeSBCAbsoluteY = 0xF9
	OpCodeSBCAbsoluteX = 0xFD
)

func IsOpCodeValidSBC(opCode uint8) bool {
	return opCode == OpCodeSBCIndirectX ||
		opCode == OpCodeSBCZero ||
		opCode == OpCodeSBCImmediate ||
		opCode == OpCodeSBCAbsolute ||
		opCode == OpCodeSBCIndirectY ||
		opCode == OpCodeSBCZeroX ||
		opCode == OpCodeSBCAbsoluteY ||
		opCode == OpCodeSBCAbsoluteX
}

func IsMnemonicValidSBC(mnemonic string) bool {
	return mnemonic == OpMnemonicSBC
}

type SBC struct {
	baseOperation
}

// 0xE1: SBC ($NN, X)
func NewSBCIndirectX(indirectAddress uint8) *SBC {
	return &SBC{
		baseOperation{
			code:        OpCodeSBCIndirectX,
			addressMode: AddrModeIndirectX,
			mnemonic:    OpMnemonicSBC,
			args:        []uint8{indirectAddress},
		},
	}
}

// 0xE5: SBC $NN
func NewSBCZero(zeroAddress uint8) *SBC {
	return &SBC{
		baseOperation{
			code:        OpCodeSBCZero,
			addressMode: AddrModeZero,
			mnemonic:    OpMnemonicSBC,
			args:        []uint8{zeroAddress},
		},
	}
}

// 0xE9: SBC #$NN
func NewSBCImmediate(value uint8) *SBC {
	return &SBC{
		baseOperation{
			code:        OpCodeSBCImmediate,
			addressMode: AddrModeImmediate,
			mnemonic:    OpMnemonicSBC,
			args:        []uint8{value},
		},
	}
}

// 0xED: SBC $NNNN
func NewSBCAbsolute(absoluteAddress uint16) *SBC {
	return &SBC{
		baseOperation{
			code:        OpCodeSBCAbsolute,
			addressMode: AddrModeAbsolute,
			mnemonic:    OpMnemonicSBC,
			args:        util.BreakWordIntoBytes(absoluteAddress),
		},
	}
}

// 0xF1: SBC ($NN), Y
func NewSBCIndirectY(indirectAddress uint8) *SBC {
	return &SBC{
		baseOperation{
			code:        OpCodeSBCIndirectY,
			addressMode: AddrModeIndirectY,
			mnemonic:    OpMnemonicSBC,
			args:        []uint8{indirectAddress},
		},
	}
}

// 0xF5: SBC $NN, X
func NewSBCZeroX(zeroAddress uint8) *SBC {
	return &SBC{
		baseOperation{
			code:        OpCodeSBCZeroX,
			addressMode: AddrModeZeroX,
			mnemonic:    OpMnemonicSBC,
			args:        []uint8{zeroAddress},
		},
	}
}

// 0xF9: SBC $NNNN, Y
func NewSBCAbsoluteY(absoluteAddress uint16) *SBC {
	return &SBC{
		baseOperation{
			code:        OpCodeSBCAbsoluteY,
			addressMode: AddrModeAbsoluteY,
			mnemonic:    OpMnemonicSBC,
			args:        util.BreakWordIntoBytes(absoluteAddress),
		},
	}
}

// 0xFD: SBC $NNNN, X
func NewSBCAbsoluteX(absoluteAddress uint16) *SBC {
	return &SBC{
		baseOperation{
			code:        OpCodeSBCAbsoluteX,
			addressMode: AddrModeAbsoluteX,
			mnemonic:    OpMnemonicSBC,
			args:        util.BreakWordIntoBytes(absoluteAddress),
		},
	}
}

func NewSBCBinary(opCode uint8, data io.Reader) (*SBC, error) {
	switch opCode {
	case OpCodeSBCIndirectX:
		var addr uint8

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewSBCIndirectX(addr), nil
	case OpCodeSBCZero:
		var addr uint8

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewSBCZero(addr), nil
	case OpCodeSBCImmediate:
		var addr uint8

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewSBCImmediate(addr), nil
	case OpCodeSBCAbsolute:
		var addr uint16

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewSBCAbsolute(addr), nil
	case OpCodeSBCIndirectY:
		var addr uint8

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewSBCIndirectY(addr), nil
	case OpCodeSBCZeroX:
		var addr uint8

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewSBCZeroX(addr), nil
	case OpCodeSBCAbsoluteY:
		var addr uint16

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewSBCAbsoluteY(addr), nil
	case OpCodeSBCAbsoluteX:
		var addr uint16

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewSBCAbsoluteX(addr), nil
	default:
		return nil, InvalidOpCodeError{
			OpCode: opCode,
		}
	}
}

func NewSBCFromBytes(addrMode uint8, arg0 uint8, arg1 uint8) (*SBC, error) {
	switch addrMode {
	case AddrModeIndirectX:
		return NewSBCIndirectX(arg0), nil
	case AddrModeZero, AddrModeRelative:
		return NewSBCZero(arg0), nil
	case AddrModeImmediate:
		return NewSBCImmediate(arg0), nil
	case AddrModeAbsolute:
		return NewSBCAbsolute(util.JoinBytesInWord([]uint8{arg0, arg1})), nil
	case AddrModeIndirectY:
		return NewSBCIndirectY(arg0), nil
	case AddrModeZeroX:
		return NewSBCZeroX(arg0), nil
	case AddrModeAbsoluteY:
		return NewSBCAbsoluteY(util.JoinBytesInWord([]uint8{arg0, arg1})), nil
	case AddrModeAbsoluteX:
		return NewSBCAbsoluteX(util.JoinBytesInWord([]uint8{arg0, arg1})), nil
	default:
		return nil, InvalidAddressModeError{
			AddressMode: addrMode,
		}
	}
}

func (op SBC) Cycles() uint8 {
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

func (op SBC) ExecuteIn(env OperationEnvironment) (uint8, error) {
	cycles := op.Cycles()
	_, operand, pageCrossed := env.FetchOperand(op)

	if pageCrossed {
		cycles++
	}

	oldA := env.GetAccumulator()
	result := oldA - operand
	signedResult := int16(int8(oldA)) - int16(int8(operand))
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
