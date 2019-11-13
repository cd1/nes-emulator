package cpu

import (
	"io"
	"log"
)

const (
	OpMnemonicBRK = "BRK"

	OpCodeBRK = 0x00

	InterruptVectorAddress = 0xFFFE
)

func IsOpCodeValidBRK(opCode uint8) bool {
	return opCode == OpCodeBRK
}

func IsMnemonicValidBRK(mnemonic string) bool {
	return mnemonic == OpMnemonicBRK
}

type BRK struct {
	baseOperation
}

// 0x00: BRK
func NewBRK() *BRK {
	return &BRK{
		baseOperation{
			code:        OpCodeBRK,
			addressMode: AddrModeImplied,
			mnemonic:    OpMnemonicBRK,
		},
	}
}

func NewBRKBinary(opCode uint8, data io.Reader) (*BRK, error) {
	switch opCode {
	case OpCodeBRK:
		return NewBRK(), nil
	default:
		return nil, InvalidOpCodeError{
			OpCode: opCode,
		}
	}
}

func NewBRKFromBytes(addrMode uint8, arg0 uint8, arg1 uint8) (*BRK, error) {
	switch addrMode {
	case AddrModeImplied:
		return NewBRK(), nil
	default:
		return nil, InvalidAddressModeError{
			AddressMode: addrMode,
		}
	}
}

func (op BRK) Cycles() uint8 {
	switch op.AddressMode() {
	case AddrModeImplied:
		return 7
	default:
		log.Printf("cannot calculate number of cycles for %#v: invalid address mode (%v)", op, op.AddressMode())
		return 0
	}
}

func (op BRK) ExecuteIn(env OperationEnvironment) (uint8, error) {
	cycles := op.Cycles()

	env.SetStatusBreak(true)
	// TODO: put this login in GetStatus? but what about /IRQ and /NMI?
	// TODO: unused?
	env.PushWordToStack(env.GetProgramCounter() + uint16(op.Size()))
	env.PushByteToStack(env.GetStatus())

	env.SetProgramCounter(env.ReadWord(InterruptVectorAddress))

	return cycles, nil
}
