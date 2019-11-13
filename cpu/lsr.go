package cpu

import (
	"encoding/binary"
	"io"
	"log"

	"github.com/cd1/nes-emulator/util"
)

const (
	OpMnemonicLSR = "LSR"

	OpCodeLSRZero        = 0x46
	OpCodeLSRAccumulator = 0x4A
	OpCodeLSRAbsolute    = 0x4E
	OpCodeLSRZeroX       = 0x56
	OpCodeLSRAbsoluteX   = 0x5E
)

func IsOpCodeValidLSR(opCode uint8) bool {
	return opCode == OpCodeLSRZero ||
		opCode == OpCodeLSRAccumulator ||
		opCode == OpCodeLSRAbsolute ||
		opCode == OpCodeLSRZeroX ||
		opCode == OpCodeLSRAbsoluteX
}

func IsMnemonicValidLSR(mnemonic string) bool {
	return mnemonic == OpMnemonicLSR
}

type LSR struct {
	baseOperation
}

// 0x46: LSR A
func NewLSRAccumulator() *LSR {
	return &LSR{
		baseOperation{
			code:        OpCodeLSRAccumulator,
			addressMode: AddrModeAccumulator,
			mnemonic:    OpMnemonicLSR,
		},
	}
}

// 0x4A: LSR $NN
func NewLSRZero(zeroAddress uint8) *LSR {
	return &LSR{
		baseOperation{
			code:        OpCodeLSRZero,
			addressMode: AddrModeZero,
			mnemonic:    OpMnemonicLSR,
			args:        []uint8{zeroAddress},
		},
	}
}

// 0x4E: LSR $NN, X
func NewLSRZeroX(zeroAddress uint8) *LSR {
	return &LSR{
		baseOperation{
			code:        OpCodeLSRZeroX,
			addressMode: AddrModeZeroX,
			mnemonic:    OpMnemonicLSR,
			args:        []uint8{zeroAddress},
		},
	}
}

// 0x56: LSR $NNNN
func NewLSRAbsolute(absoluteAddress uint16) *LSR {
	return &LSR{
		baseOperation{
			code:        OpCodeLSRAbsolute,
			addressMode: AddrModeAbsolute,
			mnemonic:    OpMnemonicLSR,
			args:        util.BreakWordIntoBytes(absoluteAddress),
		},
	}
}

// 0x5E: LSR $NNNN, X
func NewLSRAbsoluteX(absoluteAddress uint16) *LSR {
	return &LSR{
		baseOperation{
			code:        OpCodeLSRAbsoluteX,
			addressMode: AddrModeAbsoluteX,
			mnemonic:    OpMnemonicLSR,
			args:        util.BreakWordIntoBytes(absoluteAddress),
		},
	}
}

func NewLSRBinary(opCode uint8, data io.Reader) (*LSR, error) {
	switch opCode {
	case OpCodeLSRZero:
		var addr uint8

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewLSRZero(addr), nil
	case OpCodeLSRAccumulator:
		return NewLSRAccumulator(), nil
	case OpCodeLSRAbsolute:
		var addr uint16

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewLSRAbsolute(addr), nil
	case OpCodeLSRZeroX:
		var addr uint8

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewLSRZeroX(addr), nil
	case OpCodeLSRAbsoluteX:
		var addr uint16

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewLSRAbsoluteX(addr), nil
	default:
		return nil, InvalidOpCodeError{
			OpCode: opCode,
		}
	}
}

func NewLSRFromBytes(addrMode uint8, arg0 uint8, arg1 uint8) (*LSR, error) {
	switch addrMode {
	case AddrModeZero, AddrModeRelative:
		return NewLSRZero(arg0), nil
	case AddrModeAccumulator:
		return NewLSRAccumulator(), nil
	case AddrModeAbsolute:
		return NewLSRAbsolute(util.JoinBytesInWord([]uint8{arg0, arg1})), nil
	case AddrModeZeroX:
		return NewLSRZeroX(arg0), nil
	case AddrModeAbsoluteX:
		return NewLSRAbsoluteX(util.JoinBytesInWord([]uint8{arg0, arg1})), nil
	default:
		return nil, InvalidAddressModeError{
			AddressMode: addrMode,
		}
	}
}

func (op LSR) Cycles() uint8 {
	switch op.AddressMode() {
	case AddrModeZero:
		return 5
	case AddrModeAccumulator:
		return 2
	case AddrModeAbsolute:
		return 6
	case AddrModeZeroX:
		return 6
	case AddrModeAbsoluteX:
		return 7
	default:
		log.Printf("cannot calculate number of cycles for %#v: invalid address mode (%v)", op, op.AddressMode())
		return 0
	}
}

func (op LSR) ExecuteIn(env OperationEnvironment) (uint8, error) {
	cycles := op.Cycles()
	address, operand, _ := env.FetchOperand(op)

	result := operand >> 1

	if op.AddressMode() == AddrModeAccumulator {
		env.SetAccumulator(result)
	} else {
		env.WriteByte(address, result)
	}

	env.SetStatusCarry(operand&0x01 != 0x00)
	env.SetStatusZero(result == 0x00)
	env.SetStatusNegative(result&0x80 != 0x00)

	env.IncrementProgramCounter(op.Size())

	return cycles, nil
}
