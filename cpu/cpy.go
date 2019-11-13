package cpu

import (
	"encoding/binary"
	"io"
	"log"

	"github.com/cd1/nes-emulator/util"
)

const (
	OpMnemonicCPY = "CPY"

	OpCodeCPYImmediate = 0xC0
	OpCodeCPYZero      = 0xC4
	OpCodeCPYAbsolute  = 0xCC
)

func IsOpCodeValidCPY(opCode uint8) bool {
	return opCode == OpCodeCPYImmediate ||
		opCode == OpCodeCPYZero ||
		opCode == OpCodeCPYAbsolute
}

func IsMnemonicValidCPY(mnemonic string) bool {
	return mnemonic == OpMnemonicCPY
}

type CPY struct {
	baseOperation
}

// 0xC0: CPY #$NN
func NewCPYImmediate(value uint8) *CPY {
	return &CPY{
		baseOperation{
			code:        OpCodeCPYImmediate,
			addressMode: AddrModeImmediate,
			mnemonic:    OpMnemonicCPY,
			args:        []uint8{value},
		},
	}
}

// 0xC4: CPY $NN
func NewCPYZero(zeroAddress uint8) *CPY {
	return &CPY{
		baseOperation{
			code:        OpCodeCPYZero,
			addressMode: AddrModeZero,
			mnemonic:    OpMnemonicCPY,
			args:        []uint8{zeroAddress},
		},
	}
}

// 0xCC: CPY $NNNN
func NewCPYAbsolute(absoluteAddress uint16) *CPY {
	return &CPY{
		baseOperation{
			code:        OpCodeCPYAbsolute,
			addressMode: AddrModeAbsolute,
			mnemonic:    OpMnemonicCPY,
			args:        util.BreakWordIntoBytes(absoluteAddress),
		},
	}
}

func NewCPYBinary(opCode uint8, data io.Reader) (*CPY, error) {
	switch opCode {
	case OpCodeCPYImmediate:
		var addr uint8

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewCPYImmediate(addr), nil
	case OpCodeCPYZero:
		var addr uint8

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewCPYZero(addr), nil
	case OpCodeCPYAbsolute:
		var addr uint16

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewCPYAbsolute(addr), nil
	default:
		return nil, InvalidOpCodeError{
			OpCode: opCode,
		}
	}
}

func NewCPYFromBytes(addrMode uint8, arg0 uint8, arg1 uint8) (*CPY, error) {
	switch addrMode {
	case AddrModeImmediate:
		return NewCPYImmediate(arg0), nil
	case AddrModeZero, AddrModeRelative:
		return NewCPYZero(arg0), nil
	case AddrModeAbsolute:
		return NewCPYAbsolute(util.JoinBytesInWord([]uint8{arg0, arg1})), nil
	default:
		return nil, InvalidAddressModeError{
			AddressMode: addrMode,
		}
	}
}

func (op CPY) Cycles() uint8 {
	switch op.AddressMode() {
	case AddrModeZero:
		return 3
	case AddrModeImmediate:
		return 2
	case AddrModeAbsolute:
		return 4
	default:
		log.Printf("cannot calculate number of cycles for %#v: invalid address mode (%v)", op, op.AddressMode())
		return 0
	}
}

func (op CPY) ExecuteIn(env OperationEnvironment) (uint8, error) {
	cycles := op.Cycles()
	_, operand, _ := env.FetchOperand(op)

	y := env.GetIndexY()
	env.SetStatusCarry(y >= operand)
	env.SetStatusZero(y-operand == 0x00)
	env.SetStatusNegative((y-operand)&0x80 != 0x00)

	env.IncrementProgramCounter(op.Size())

	return cycles, nil
}
