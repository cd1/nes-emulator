package cpu

import (
	"encoding/binary"
	"io"
	"log"

	"github.com/cd1/nes-emulator/util"
)

const (
	OpMnemonicDEC = "DEC"

	OpCodeDECZero      = 0xC6
	OpCodeDECAbsolute  = 0xCE
	OpCodeDECZeroX     = 0xD6
	OpCodeDECAbsoluteX = 0xDE
)

func IsOpCodeValidDEC(opCode uint8) bool {
	return opCode == OpCodeDECZero ||
		opCode == OpCodeDECAbsolute ||
		opCode == OpCodeDECZeroX ||
		opCode == OpCodeDECAbsoluteX
}

func IsMnemonicValidDEC(mnemonic string) bool {
	return mnemonic == OpMnemonicDEC
}

type DEC struct {
	baseOperation
}

// 0xC6: DEC $NN
func NewDECZero(zeroAddress uint8) *DEC {
	return &DEC{
		baseOperation{
			code:        OpCodeDECZero,
			addressMode: AddrModeZero,
			mnemonic:    OpMnemonicDEC,
			args:        []uint8{zeroAddress},
		},
	}
}

// 0xCE: DEC $NNNN
func NewDECAbsolute(absoluteAddress uint16) *DEC {
	return &DEC{
		baseOperation{
			code:        OpCodeDECAbsolute,
			addressMode: AddrModeAbsolute,
			mnemonic:    OpMnemonicDEC,
			args:        util.BreakWordIntoBytes(absoluteAddress),
		},
	}
}

// 0xD6: DEC $NN, X
func NewDECZeroX(zeroAddress uint8) *DEC {
	return &DEC{
		baseOperation{
			code:        OpCodeDECZeroX,
			addressMode: AddrModeZeroX,
			mnemonic:    OpMnemonicDEC,
			args:        []uint8{zeroAddress},
		},
	}
}

// 0xDE: DEC $NNNN, X
func NewDECAbsoluteX(absoluteAddress uint16) *DEC {
	return &DEC{
		baseOperation{
			code:        OpCodeDECAbsoluteX,
			addressMode: AddrModeAbsoluteX,
			mnemonic:    OpMnemonicDEC,
			args:        util.BreakWordIntoBytes(absoluteAddress),
		},
	}
}

func NewDECBinary(opCode uint8, data io.Reader) (*DEC, error) {
	switch opCode {
	case OpCodeDECZero:
		var addr uint8

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewDECZero(addr), nil
	case OpCodeDECAbsolute:
		var addr uint16

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewDECAbsolute(addr), nil
	case OpCodeDECZeroX:
		var addr uint8

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewDECZeroX(addr), nil
	case OpCodeDECAbsoluteX:
		var addr uint16

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewDECAbsoluteX(addr), nil
	default:
		return nil, InvalidOpCodeError{
			OpCode: opCode,
		}
	}
}

func NewDECFromBytes(addrMode uint8, arg0 uint8, arg1 uint8) (*DEC, error) {
	switch addrMode {
	case AddrModeZero, AddrModeRelative:
		return NewDECZero(arg0), nil
	case AddrModeAbsolute:
		return NewDECAbsolute(util.JoinBytesInWord([]uint8{arg0, arg1})), nil
	case AddrModeZeroX:
		return NewDECZeroX(arg0), nil
	case AddrModeAbsoluteX:
		return NewDECAbsoluteX(util.JoinBytesInWord([]uint8{arg0, arg1})), nil
	default:
		return nil, InvalidAddressModeError{
			AddressMode: addrMode,
		}
	}
}

func (op DEC) Cycles() uint8 {
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

func (op DEC) ExecuteIn(env OperationEnvironment) (uint8, error) {
	cycles := op.Cycles()
	address, operand, _ := env.FetchOperand(op)

	newMemValue := operand - 1
	env.WriteByte(address, newMemValue)

	env.SetStatusZero(newMemValue == 0x00)
	env.SetStatusNegative(newMemValue&0x80 != 0x00)

	env.IncrementProgramCounter(op.Size())

	return cycles, nil
}
