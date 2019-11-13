package cpu

import (
	"io"
	"log"
)

const (
	OpMnemonicTYA = "TYA"

	OpCodeTYA = 0x98
)

func IsOpCodeValidTYA(opCode uint8) bool {
	return opCode == OpCodeTYA
}

func IsMnemonicValidTYA(mnemonic string) bool {
	return mnemonic == OpMnemonicTYA
}

type TYA struct {
	baseOperation
}

// 0x98: TYA
func NewTYA() *TYA {
	return &TYA{
		baseOperation{
			code:        OpCodeTYA,
			addressMode: AddrModeImplied,
			mnemonic:    OpMnemonicTYA,
		},
	}
}

func NewTYABinary(opCode uint8, data io.Reader) (*TYA, error) {
	switch opCode {
	case OpCodeTYA:
		return NewTYA(), nil
	default:
		return nil, InvalidOpCodeError{
			OpCode: opCode,
		}
	}
}

func NewTYAFromBytes(addrMode uint8, arg0 uint8, arg1 uint8) (*TYA, error) {
	switch addrMode {
	case AddrModeImplied:
		return NewTYA(), nil
	default:
		return nil, InvalidAddressModeError{
			AddressMode: addrMode,
		}
	}
}

func (op TYA) Cycles() uint8 {
	switch op.AddressMode() {
	case AddrModeImplied:
		return 2
	default:
		log.Printf("cannot calculate number of cycles for %#v: invalid address mode (%v)", op, op.AddressMode())
		return 0
	}
}

func (op TYA) ExecuteIn(env OperationEnvironment) (uint8, error) {
	cycles := op.Cycles()

	y := env.GetIndexY()
	env.SetAccumulator(y)

	env.SetStatusZero(y == 0x00)
	env.SetStatusNegative(y&0x80 != 0x00)

	env.IncrementProgramCounter(op.Size())

	return cycles, nil
}
