package cpu

import (
	"encoding/binary"
	"io"
	"log"

	"github.com/cd1/nes-emulator/util"
)

const (
	OpMnemonicEOR = "EOR"

	OpCodeEORIndirectX = 0x41
	OpCodeEORZero      = 0x45
	OpCodeEORImmediate = 0x49
	OpCodeEORAbsolute  = 0x4D
	OpCodeEORIndirectY = 0x51
	OpCodeEORZeroX     = 0x55
	OpCodeEORAbsoluteY = 0x59
	OpCodeEORAbsoluteX = 0x5D
)

func IsOpCodeValidEOR(opCode uint8) bool {
	return opCode == OpCodeEORIndirectX ||
		opCode == OpCodeEORZero ||
		opCode == OpCodeEORImmediate ||
		opCode == OpCodeEORAbsolute ||
		opCode == OpCodeEORIndirectY ||
		opCode == OpCodeEORZeroX ||
		opCode == OpCodeEORAbsoluteY ||
		opCode == OpCodeEORAbsoluteX
}

func IsMnemonicValidEOR(mnemonic string) bool {
	return mnemonic == OpMnemonicEOR
}

type EOR struct {
	baseOperation
}

// 0x41: EOR $NN, X
func NewEORIndirectX(indirectAddress uint8) *EOR {
	return &EOR{
		baseOperation{
			code:        OpCodeEORIndirectX,
			addressMode: AddrModeIndirectX,
			mnemonic:    OpMnemonicEOR,
			args:        []uint8{indirectAddress},
		},
	}
}

// 0x45: EOR $NN
func NewEORZero(zeroAddress uint8) *EOR {
	return &EOR{
		baseOperation{
			code:        OpCodeEORZero,
			addressMode: AddrModeZero,
			mnemonic:    OpMnemonicEOR,
			args:        []uint8{zeroAddress},
		},
	}
}

// 0x49: EOR #$NN
func NewEORImmediate(value uint8) *EOR {
	return &EOR{
		baseOperation{
			code:        OpCodeEORImmediate,
			addressMode: AddrModeImmediate,
			mnemonic:    OpMnemonicEOR,
			args:        []uint8{value},
		},
	}
}

// 0x4D: EOR $NNNN
func NewEORAbsolute(absoluteAddress uint16) *EOR {
	return &EOR{
		baseOperation{
			code:        OpCodeEORAbsolute,
			addressMode: AddrModeAbsolute,
			mnemonic:    OpMnemonicEOR,
			args:        util.BreakWordIntoBytes(absoluteAddress),
		},
	}
}

// 0x51: EOR ($NN), Y
func NewEORIndirectY(indirectAddress uint8) *EOR {
	return &EOR{
		baseOperation{
			code:        OpCodeEORIndirectY,
			addressMode: AddrModeIndirectY,
			mnemonic:    OpMnemonicEOR,
			args:        []uint8{indirectAddress},
		},
	}
}

// 0x55: EOR $NN, X
func NewEORZeroX(zeroAddress uint8) *EOR {
	return &EOR{
		baseOperation{
			code:        OpCodeEORZeroX,
			addressMode: AddrModeZeroX,
			mnemonic:    OpMnemonicEOR,
			args:        []uint8{zeroAddress},
		},
	}
}

// 0x59: EOR $NNNN, Y
func NewEORAbsoluteY(absoluteAddress uint16) *EOR {
	return &EOR{
		baseOperation{
			code:        OpCodeEORAbsoluteY,
			addressMode: AddrModeAbsoluteY,
			mnemonic:    OpMnemonicEOR,
			args:        util.BreakWordIntoBytes(absoluteAddress),
		},
	}
}

// 0x5D: EOR $NNNN, X
func NewEORAbsoluteX(absoluteAddress uint16) *EOR {
	return &EOR{
		baseOperation{
			code:        OpCodeEORAbsoluteX,
			addressMode: AddrModeAbsoluteX,
			mnemonic:    OpMnemonicEOR,
			args:        util.BreakWordIntoBytes(absoluteAddress),
		},
	}
}

func NewEORBinary(opCode uint8, data io.Reader) (*EOR, error) {
	switch opCode {
	case OpCodeEORIndirectX:
		var addr uint8

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewEORIndirectX(addr), nil
	case OpCodeEORZero:
		var addr uint8

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewEORZero(addr), nil
	case OpCodeEORImmediate:
		var addr uint8

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewEORImmediate(addr), nil
	case OpCodeEORAbsolute:
		var addr uint16

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewEORAbsolute(addr), nil
	case OpCodeEORIndirectY:
		var addr uint8

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewEORIndirectY(addr), nil
	case OpCodeEORZeroX:
		var addr uint8

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewEORZeroX(addr), nil
	case OpCodeEORAbsoluteY:
		var addr uint16

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewEORAbsoluteY(addr), nil
	case OpCodeEORAbsoluteX:
		var addr uint16

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewEORAbsoluteX(addr), nil
	default:
		return nil, InvalidOpCodeError{
			OpCode: opCode,
		}
	}
}

func NewEORFromBytes(addrMode uint8, arg0 uint8, arg1 uint8) (*EOR, error) {
	switch addrMode {
	case AddrModeIndirectX:
		return NewEORIndirectX(arg0), nil
	case AddrModeZero, AddrModeRelative:
		return NewEORZero(arg0), nil
	case AddrModeImmediate:
		return NewEORImmediate(arg0), nil
	case AddrModeAbsolute:
		return NewEORAbsolute(util.JoinBytesInWord([]uint8{arg0, arg1})), nil
	case AddrModeIndirectY:
		return NewEORIndirectY(arg0), nil
	case AddrModeZeroX:
		return NewEORZeroX(arg0), nil
	case AddrModeAbsoluteY:
		return NewEORAbsoluteY(util.JoinBytesInWord([]uint8{arg0, arg1})), nil
	case AddrModeAbsoluteX:
		return NewEORAbsoluteX(util.JoinBytesInWord([]uint8{arg0, arg1})), nil
	default:
		return nil, InvalidAddressModeError{
			AddressMode: addrMode,
		}
	}
}

func (op EOR) Cycles() uint8 {
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

func (op EOR) ExecuteIn(env OperationEnvironment) (uint8, error) {
	cycles := op.Cycles()
	_, operand, pageCrossed := env.FetchOperand(op)

	if pageCrossed {
		cycles++
	}

	result := env.GetAccumulator() ^ operand
	env.SetAccumulator(result)

	env.SetStatusZero(result == 0x00)
	env.SetStatusNegative(result&0x80 != 0x00)

	env.IncrementProgramCounter(op.Size())

	return cycles, nil
}
