package cpu

import (
	"io"
	"log"
)

const (
	OpMnemonicCLC = "CLC"

	OpCodeCLC = 0x18
)

func IsOpCodeValidCLC(opCode uint8) bool {
	return opCode == OpCodeCLC
}

func IsMnemonicValidCLC(mnemonic string) bool {
	return mnemonic == OpMnemonicCLC
}

type CLC struct {
	baseOperation
}

// 0x18: CLC
func NewCLC() *CLC {
	return &CLC{
		baseOperation{
			code:        OpCodeCLC,
			addressMode: AddrModeImplied,
			mnemonic:    OpMnemonicCLC,
		},
	}
}

func NewCLCBinary(opCode uint8, data io.Reader) (*CLC, error) {
	switch opCode {
	case OpCodeCLC:
		return NewCLC(), nil
	default:
		return nil, InvalidOpCodeError{
			OpCode: opCode,
		}
	}
}

func NewCLCFromBytes(addrMode uint8, arg0 uint8, arg1 uint8) (*CLC, error) {
	switch addrMode {
	case AddrModeImplied:
		return NewCLC(), nil
	default:
		return nil, InvalidAddressModeError{
			AddressMode: addrMode,
		}
	}
}

func (op CLC) Cycles() uint8 {
	switch op.AddressMode() {
	case AddrModeImplied:
		return 2
	default:
		log.Printf("cannot calculate number of cycles for %#v: invalid address mode (%v)", op, op.AddressMode())
		return 0
	}
}

func (op CLC) ExecuteIn(env OperationEnvironment) (uint8, error) {
	cycles := op.Cycles()

	env.SetStatusCarry(false)

	env.IncrementProgramCounter(op.Size())

	return cycles, nil
}
