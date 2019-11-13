package cpu

import (
	"io"
	"log"
)

const (
	OpMnemonicPLA = "PLA"

	OpCodePLA = 0x68
)

func IsOpCodeValidPLA(opCode uint8) bool {
	return opCode == OpCodePLA
}

func IsMnemonicValidPLA(mnemonic string) bool {
	return mnemonic == OpMnemonicPLA
}

type PLA struct {
	baseOperation
}

// 0x68: PLA
func NewPLA() *PLA {
	return &PLA{
		baseOperation{
			code:        OpCodePLA,
			addressMode: AddrModeImplied,
			mnemonic:    OpMnemonicPLA,
		},
	}
}

func NewPLABinary(opCode uint8, data io.Reader) (*PLA, error) {
	switch opCode {
	case OpCodePLA:
		return NewPLA(), nil
	default:
		return nil, InvalidOpCodeError{
			OpCode: opCode,
		}
	}
}

func NewPLAFromBytes(addrMode uint8, arg0 uint8, arg1 uint8) (*PLA, error) {
	switch addrMode {
	case AddrModeImplied:
		return NewPLA(), nil
	default:
		return nil, InvalidAddressModeError{
			AddressMode: addrMode,
		}
	}
}

func (op PLA) Cycles() uint8 {
	switch op.AddressMode() {
	case AddrModeImplied:
		return 4
	default:
		log.Printf("cannot calculate number of cycles for %#v: invalid address mode (%v)", op, op.AddressMode())
		return 0
	}
}

func (op PLA) ExecuteIn(env OperationEnvironment) (uint8, error) {
	cycles := op.Cycles()

	newA := env.PullByteFromStack()
	env.SetAccumulator(newA)

	env.SetStatusZero(newA == 0x00)
	env.SetStatusNegative(newA&0x80 != 0x00)

	env.IncrementProgramCounter(op.Size())

	return cycles, nil
}
