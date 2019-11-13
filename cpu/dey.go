package cpu

import (
	"io"
	"log"
)

const (
	OpMnemonicDEY = "DEY"

	OpCodeDEY = 0x88
)

func IsOpCodeValidDEY(opCode uint8) bool {
	return opCode == OpCodeDEY
}

func IsMnemonicValidDEY(mnemonic string) bool {
	return mnemonic == OpMnemonicDEY
}

type DEY struct {
	baseOperation
}

// 0x88: DEY
func NewDEY() *DEY {
	return &DEY{
		baseOperation{
			code:        OpCodeDEY,
			addressMode: AddrModeImplied,
			mnemonic:    OpMnemonicDEY,
		},
	}
}

func NewDEYBinary(opCode uint8, data io.Reader) (*DEY, error) {
	switch opCode {
	case OpCodeDEY:
		return NewDEY(), nil
	default:
		return nil, InvalidOpCodeError{
			OpCode: opCode,
		}
	}
}

func NewDEYFromBytes(addrMode uint8, arg0 uint8, arg1 uint8) (*DEY, error) {
	switch addrMode {
	case AddrModeImplied:
		return NewDEY(), nil
	default:
		return nil, InvalidAddressModeError{
			AddressMode: addrMode,
		}
	}
}

func (op DEY) Cycles() uint8 {
	switch op.AddressMode() {
	case AddrModeImplied:
		return 2
	default:
		log.Printf("cannot calculate number of cycles for %#v: invalid address mode (%v)", op, op.AddressMode())
		return 0
	}
}

func (op DEY) ExecuteIn(env OperationEnvironment) (uint8, error) {
	cycles := op.Cycles()

	newY := env.GetIndexY() - 1
	env.SetIndexY(newY)

	env.SetStatusZero(newY == 0x00)
	env.SetStatusNegative(newY&0x80 != 0x00)

	env.IncrementProgramCounter(op.Size())

	return cycles, nil
}
