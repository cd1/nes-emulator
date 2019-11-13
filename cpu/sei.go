package cpu

import (
	"io"
	"log"
)

const (
	OpMnemonicSEI = "SEI"

	OpCodeSEI = 0x78
)

func IsOpCodeValidSEI(opCode uint8) bool {
	return opCode == OpCodeSEI
}

func IsMnemonicValidSEI(mnemonic string) bool {
	return mnemonic == OpMnemonicSEI
}

type SEI struct {
	baseOperation
}

// 0x78: SEI
func NewSEI() *SEI {
	return &SEI{
		baseOperation{
			code:        OpCodeSEI,
			addressMode: AddrModeImplied,
			mnemonic:    OpMnemonicSEI,
		},
	}
}

func NewSEIBinary(opCode uint8, data io.Reader) (*SEI, error) {
	switch opCode {
	case OpCodeSEI:
		return NewSEI(), nil
	default:
		return nil, InvalidOpCodeError{
			OpCode: opCode,
		}
	}
}

func NewSEIFromBytes(addrMode uint8, arg0 uint8, arg1 uint8) (*SEI, error) {
	switch addrMode {
	case AddrModeImplied:
		return NewSEI(), nil
	default:
		return nil, InvalidAddressModeError{
			AddressMode: addrMode,
		}
	}
}

func (op SEI) Cycles() uint8 {
	switch op.AddressMode() {
	case AddrModeImplied:
		return 2
	default:
		log.Printf("cannot calculate number of cycles for %#v: invalid address mode (%v)", op, op.AddressMode())
		return 0
	}
}

func (op SEI) ExecuteIn(env OperationEnvironment) (uint8, error) {
	cycles := op.Cycles()

	env.SetStatusInterrupt(true)

	env.IncrementProgramCounter(op.Size())

	return cycles, nil
}
