package cpu

import (
	"io"
	"log"
)

const (
	OpMnemonicDEX = "DEX"

	OpCodeDEX = 0xCA
)

func IsOpCodeValidDEX(opCode uint8) bool {
	return opCode == OpCodeDEX
}

func IsMnemonicValidDEX(mnemonic string) bool {
	return mnemonic == OpMnemonicDEX
}

type DEX struct {
	baseOperation
}

// 0xCA: DEX
func NewDEX() *DEX {
	return &DEX{
		baseOperation{
			code:        OpCodeDEX,
			addressMode: AddrModeImplied,
			mnemonic:    OpMnemonicDEX,
		},
	}
}

func NewDEXBinary(opCode uint8, data io.Reader) (*DEX, error) {
	switch opCode {
	case OpCodeDEX:
		return NewDEX(), nil
	default:
		return nil, InvalidOpCodeError{
			OpCode: opCode,
		}
	}
}

func NewDEXFromBytes(addrMode uint8, arg0 uint8, arg1 uint8) (*DEX, error) {
	switch addrMode {
	case AddrModeImplied:
		return NewDEX(), nil
	default:
		return nil, InvalidAddressModeError{
			AddressMode: addrMode,
		}
	}
}

func (op DEX) Cycles() uint8 {
	switch op.AddressMode() {
	case AddrModeImplied:
		return 2
	default:
		log.Printf("cannot calculate number of cycles for %#v: invalid address mode (%v)", op, op.AddressMode())
		return 0
	}
}

func (op DEX) ExecuteIn(env OperationEnvironment) (uint8, error) {
	cycles := op.Cycles()

	newX := env.GetIndexX() - 1
	env.SetIndexX(newX)

	env.SetStatusZero(newX == 0x00)
	env.SetStatusNegative(newX&0x80 != 0x00)

	env.IncrementProgramCounter(op.Size())

	return cycles, nil
}
