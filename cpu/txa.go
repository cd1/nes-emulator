package cpu

import (
	"io"
	"log"
)

const (
	OpMnemonicTXA = "TXA"

	OpCodeTXA = 0x8A
)

func IsOpCodeValidTXA(opCode uint8) bool {
	return opCode == OpCodeTXA
}

func IsMnemonicValidTXA(mnemonic string) bool {
	return mnemonic == OpMnemonicTXA
}

type TXA struct {
	baseOperation
}

// 0x8A: TXA
func NewTXA() *TXA {
	return &TXA{
		baseOperation{
			code:        OpCodeTXA,
			addressMode: AddrModeImplied,
			mnemonic:    OpMnemonicTXA,
		},
	}
}

func NewTXABinary(opCode uint8, data io.Reader) (*TXA, error) {
	switch opCode {
	case OpCodeTXA:
		return NewTXA(), nil
	default:
		return nil, InvalidOpCodeError{
			OpCode: opCode,
		}
	}
}

func NewTXAFromBytes(addrMode uint8, arg0 uint8, arg1 uint8) (*TXA, error) {
	switch addrMode {
	case AddrModeImplied:
		return NewTXA(), nil
	default:
		return nil, InvalidAddressModeError{
			AddressMode: addrMode,
		}
	}
}

func (op TXA) Cycles() uint8 {
	switch op.AddressMode() {
	case AddrModeImplied:
		return 2
	default:
		log.Printf("cannot calculate number of cycles for %#v: invalid address mode (%v)", op, op.AddressMode())
		return 0
	}
}

func (op TXA) ExecuteIn(env OperationEnvironment) (uint8, error) {
	cycles := op.Cycles()

	x := env.GetIndexX()
	env.SetAccumulator(x)

	env.SetStatusZero(x == 0x00)
	env.SetStatusNegative(x&0x80 != 0x00)

	env.IncrementProgramCounter(op.Size())

	return cycles, nil
}
