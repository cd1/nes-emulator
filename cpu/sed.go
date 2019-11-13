package cpu

import (
	"io"
	"log"
)

const (
	OpMnemonicSED = "SED"

	OpCodeSED = 0xF8
)

func IsOpCodeValidSED(opCode uint8) bool {
	return opCode == OpCodeSED
}

func IsMnemonicValidSED(mnemonic string) bool {
	return mnemonic == OpMnemonicSED
}

type SED struct {
	baseOperation
}

// 0xF8: SED
func NewSED() *SED {
	return &SED{
		baseOperation{
			code:        OpCodeSED,
			addressMode: AddrModeImplied,
			mnemonic:    OpMnemonicSED,
		},
	}
}

func NewSEDBinary(opCode uint8, data io.Reader) (*SED, error) {
	switch opCode {
	case OpCodeSED:
		return NewSED(), nil
	default:
		return nil, InvalidOpCodeError{
			OpCode: opCode,
		}
	}
}

func NewSEDFromBytes(addrMode uint8, arg0 uint8, arg1 uint8) (*SED, error) {
	switch addrMode {
	case AddrModeImplied:
		return NewSED(), nil
	default:
		return nil, InvalidAddressModeError{
			AddressMode: addrMode,
		}
	}
}

func (op SED) Cycles() uint8 {
	switch op.AddressMode() {
	case AddrModeImplied:
		return 2
	default:
		log.Printf("cannot calculate number of cycles for %#v: invalid address mode (%v)", op, op.AddressMode())
		return 0
	}
}

func (op SED) ExecuteIn(env OperationEnvironment) (uint8, error) {
	cycles := op.Cycles()

	env.SetStatusDecimal(true)

	env.IncrementProgramCounter(op.Size())

	return cycles, nil
}
