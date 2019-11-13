package cpu

import (
	"io"
	"log"
)

const (
	OpMnemonicTAX = "TAX"

	OpCodeTAX = 0xAA
)

func IsOpCodeValidTAX(opCode uint8) bool {
	return opCode == OpCodeTAX
}

func IsMnemonicValidTAX(mnemonic string) bool {
	return mnemonic == OpMnemonicTAX
}

type TAX struct {
	baseOperation
}

// 0xAA: TAX
func NewTAX() *TAX {
	return &TAX{
		baseOperation{
			code:        OpCodeTAX,
			addressMode: AddrModeImplied,
			mnemonic:    OpMnemonicTAX,
		},
	}
}

func NewTAXBinary(opCode uint8, data io.Reader) (*TAX, error) {
	switch opCode {
	case OpCodeTAX:
		return NewTAX(), nil
	default:
		return nil, InvalidOpCodeError{
			OpCode: opCode,
		}
	}
}

func NewTAXFromBytes(addrMode uint8, arg0 uint8, arg1 uint8) (*TAX, error) {
	switch addrMode {
	case AddrModeImplied:
		return NewTAX(), nil
	default:
		return nil, InvalidAddressModeError{
			AddressMode: addrMode,
		}
	}
}

func (op TAX) Cycles() uint8 {
	switch op.AddressMode() {
	case AddrModeImplied:
		return 2
	default:
		log.Printf("cannot calculate number of cycles for %#v: invalid address mode (%v)", op, op.AddressMode())
		return 0
	}
}

func (op TAX) ExecuteIn(env OperationEnvironment) (uint8, error) {
	cycles := op.Cycles()

	a := env.GetAccumulator()
	env.SetIndexX(a)

	env.SetStatusZero(a == 0x00)
	env.SetStatusNegative(a&0x80 != 0x00)

	env.IncrementProgramCounter(op.Size())

	return cycles, nil
}
