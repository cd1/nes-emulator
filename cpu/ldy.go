package cpu

import (
	"encoding/binary"
	"io"
	"log"

	"github.com/cd1/nes-emulator/util"
)

const (
	OpMnemonicLDY = "LDY"

	OpCodeLDYImmediate = 0xA0
	OpCodeLDYZero      = 0xA4
	OpCodeLDYAbsolute  = 0xAC
	OpCodeLDYZeroX     = 0xB4
	OpCodeLDYAbsoluteX = 0xBC
)

func IsOpCodeValidLDY(opCode uint8) bool {
	return opCode == OpCodeLDYImmediate ||
		opCode == OpCodeLDYZero ||
		opCode == OpCodeLDYAbsolute ||
		opCode == OpCodeLDYZeroX ||
		opCode == OpCodeLDYAbsoluteX
}

func IsMnemonicValidLDY(mnemonic string) bool {
	return mnemonic == OpMnemonicLDY
}

type LDY struct {
	baseOperation
}

func NewLDYImmediate(value uint8) *LDY {
	return &LDY{
		baseOperation{
			code:        OpCodeLDYImmediate,
			addressMode: AddrModeImmediate,
			mnemonic:    OpMnemonicLDY,
			args:        []uint8{value},
		},
	}
}

func NewLDYZero(zeroAddress uint8) *LDY {
	return &LDY{
		baseOperation{
			code:        OpCodeLDYZero,
			addressMode: AddrModeZero,
			mnemonic:    OpMnemonicLDY,
			args:        []uint8{zeroAddress},
		},
	}
}

func NewLDYAbsolute(absoluteAddress uint16) *LDY {
	return &LDY{
		baseOperation{
			code:        OpCodeLDYAbsolute,
			addressMode: AddrModeAbsolute,
			mnemonic:    OpMnemonicLDY,
			args:        util.BreakWordIntoBytes(absoluteAddress),
		},
	}
}

func NewLDYZeroX(zeroAddress uint8) *LDY {
	return &LDY{
		baseOperation{
			code:        OpCodeLDYZeroX,
			addressMode: AddrModeZeroX,
			mnemonic:    OpMnemonicLDY,
			args:        []uint8{zeroAddress},
		},
	}
}

func NewLDYAbsoluteX(absoluteAddress uint16) *LDY {
	return &LDY{
		baseOperation{
			code:        OpCodeLDYAbsoluteX,
			addressMode: AddrModeAbsoluteX,
			mnemonic:    OpMnemonicLDY,
			args:        util.BreakWordIntoBytes(absoluteAddress),
		},
	}
}

func NewLDYBinary(opCode uint8, data io.Reader) (*LDY, error) {
	switch opCode {
	case OpCodeLDYImmediate:
		var addr uint8

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewLDYImmediate(addr), nil
	case OpCodeLDYZero:
		var addr uint8

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewLDYZero(addr), nil
	case OpCodeLDYAbsolute:
		var addr uint16

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewLDYAbsolute(addr), nil
	case OpCodeLDYZeroX:
		var addr uint8

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewLDYZeroX(addr), nil
	case OpCodeLDYAbsoluteX:
		var addr uint16

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewLDYAbsoluteX(addr), nil
	default:
		return nil, InvalidOpCodeError{opCode}
	}
}

func NewLDYFromBytes(addrMode uint8, arg0 uint8, arg1 uint8) (*LDY, error) {
	switch addrMode {
	case AddrModeImmediate:
		return NewLDYImmediate(arg0), nil
	case AddrModeZero, AddrModeRelative:
		return NewLDYZero(arg0), nil
	case AddrModeAbsolute:
		return NewLDYAbsolute(util.JoinBytesInWord([]uint8{arg0, arg1})), nil
	case AddrModeZeroX:
		return NewLDYZeroX(arg0), nil
	case AddrModeAbsoluteX:
		return NewLDYAbsoluteX(util.JoinBytesInWord([]uint8{arg0, arg1})), nil
	default:
		return nil, InvalidAddressModeError{
			AddressMode: addrMode,
		}
	}
}

func (op LDY) Cycles() uint8 {
	switch op.AddressMode() {
	case AddrModeZero:
		return 3
	case AddrModeImmediate:
		return 2
	case AddrModeAbsolute:
		return 4
	case AddrModeZeroX:
		return 4
	case AddrModeAbsoluteX:
		return 4
	default:
		log.Printf("cannot calculate number of cycles for %#v: invalid address mode (%v)", op, op.AddressMode())
		return 0
	}
}
func (op LDY) ExecuteIn(env OperationEnvironment) (uint8, error) {
	cycles := op.Cycles()
	_, operand, pageCrossed := env.FetchOperand(op)

	if pageCrossed {
		cycles++
	}

	env.SetIndexY(operand)

	env.SetStatusZero(operand == 0x00)
	env.SetStatusNegative(operand&0x80 != 0)

	env.IncrementProgramCounter(op.Size())

	return cycles, nil
}
