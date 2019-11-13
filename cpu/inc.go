package cpu

import (
	"encoding/binary"
	"io"
	"log"

	"github.com/cd1/nes-emulator/util"
)

const (
	OpMnemonicINC = "INC"

	OpCodeINCZero      = 0xE6
	OpCodeINCAbsolute  = 0xEE
	OpCodeINCZeroX     = 0xF6
	OpCodeINCAbsoluteX = 0xFE
)

func IsOpCodeValidINC(opCode uint8) bool {
	return opCode == OpCodeINCZero ||
		opCode == OpCodeINCAbsolute ||
		opCode == OpCodeINCZeroX ||
		opCode == OpCodeINCAbsoluteX
}

func IsMnemonicValidINC(mnemonic string) bool {
	return mnemonic == OpMnemonicINC
}

type INC struct {
	baseOperation
}

// 0xE6: INC $NN
func NewINCZero(zeroAddress uint8) *INC {
	return &INC{
		baseOperation{
			code:        OpCodeINCZero,
			addressMode: AddrModeZero,
			mnemonic:    OpMnemonicINC,
			args:        []uint8{zeroAddress},
		},
	}
}

// 0xEE: INC $NNNN
func NewINCAbsolute(absoluteAddress uint16) *INC {
	return &INC{
		baseOperation{
			code:        OpCodeINCAbsolute,
			addressMode: AddrModeAbsolute,
			mnemonic:    OpMnemonicINC,
			args:        util.BreakWordIntoBytes(absoluteAddress),
		},
	}
}

// 0xF6: INC $NN, X
func NewINCZeroX(zeroAddress uint8) *INC {
	return &INC{
		baseOperation{
			code:        OpCodeINCZeroX,
			addressMode: AddrModeZeroX,
			mnemonic:    OpMnemonicINC,
			args:        []uint8{zeroAddress},
		},
	}
}

// 0xFE: INC $NNNN, X
func NewINCAbsoluteX(absoluteAddress uint16) *INC {
	return &INC{
		baseOperation{
			code:        OpCodeINCAbsoluteX,
			addressMode: AddrModeAbsoluteX,
			mnemonic:    OpMnemonicINC,
			args:        util.BreakWordIntoBytes(absoluteAddress),
		},
	}
}

func NewINCBinary(opCode uint8, data io.Reader) (*INC, error) {
	switch opCode {
	case OpCodeINCZero:
		var addr uint8

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewINCZero(addr), nil
	case OpCodeINCAbsolute:
		var addr uint16

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewINCAbsolute(addr), nil
	case OpCodeINCZeroX:
		var addr uint8

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewINCZeroX(addr), nil
	case OpCodeINCAbsoluteX:
		var addr uint16

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewINCAbsoluteX(addr), nil
	default:
		return nil, InvalidOpCodeError{
			OpCode: opCode,
		}
	}
}

func NewINCFromBytes(addrMode uint8, arg0 uint8, arg1 uint8) (*INC, error) {
	switch addrMode {
	case AddrModeZero, AddrModeRelative:
		return NewINCZero(arg0), nil
	case AddrModeAbsolute:
		return NewINCAbsolute(util.JoinBytesInWord([]uint8{arg0, arg1})), nil
	case AddrModeZeroX:
		return NewINCZeroX(arg0), nil
	case AddrModeAbsoluteX:
		return NewINCAbsoluteX(util.JoinBytesInWord([]uint8{arg0, arg1})), nil
	default:
		return nil, InvalidAddressModeError{
			AddressMode: addrMode,
		}
	}
}

func (op INC) Cycles() uint8 {
	switch op.AddressMode() {
	case AddrModeZero:
		return 5
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

func (op INC) ExecuteIn(env OperationEnvironment) (uint8, error) {
	cycles := op.Cycles()
	address, operand, _ := env.FetchOperand(op)

	newMemValue := operand + 1
	env.WriteByte(address, newMemValue)

	env.SetStatusZero(newMemValue == 0x00)
	env.SetStatusNegative(newMemValue&0x80 != 0x00)

	env.IncrementProgramCounter(op.Size())

	return cycles, nil
}
