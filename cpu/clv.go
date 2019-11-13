package cpu

import (
	"io"
	"log"
)

const (
	OpMnemonicCLV = "CLV"

	OpCodeCLV = 0xB8
)

func IsOpCodeValidCLV(opCode uint8) bool {
	return opCode == OpCodeCLV
}

func IsMnemonicValidCLV(mnemonic string) bool {
	return mnemonic == OpMnemonicCLV
}

type CLV struct {
	baseOperation
}

// 0xB8: CLV
func NewCLV() *CLV {
	return &CLV{
		baseOperation{
			code:        OpCodeCLV,
			addressMode: AddrModeImplied,
			mnemonic:    OpMnemonicCLV,
		},
	}
}

func NewCLVBinary(opCode uint8, data io.Reader) (*CLV, error) {
	switch opCode {
	case OpCodeCLV:
		return NewCLV(), nil
	default:
		return nil, InvalidOpCodeError{
			OpCode: opCode,
		}
	}
}

func NewCLVFromBytes(addrMode uint8, arg0 uint8, arg1 uint8) (*CLV, error) {
	switch addrMode {
	case AddrModeImplied:
		return NewCLV(), nil
	default:
		return nil, InvalidAddressModeError{
			AddressMode: addrMode,
		}
	}
}

func (op CLV) Cycles() uint8 {
	switch op.AddressMode() {
	case AddrModeImplied:
		return 2
	default:
		log.Printf("cannot calculate number of cycles for %#v: invalid address mode (%v)", op, op.AddressMode())
		return 0
	}
}

func (op CLV) ExecuteIn(env OperationEnvironment) (uint8, error) {
	cycles := op.Cycles()

	env.SetStatusOverflow(false)

	env.IncrementProgramCounter(op.Size())

	return cycles, nil
}
