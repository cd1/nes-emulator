package cpu

import (
	"io"
	"log"
)

const (
	OpMnemonicTXS = "TXS"

	OpCodeTXS = 0x9A
)

func IsOpCodeValidTXS(opCode uint8) bool {
	return opCode == OpCodeTXS
}

func IsMnemonicValidTXS(mnemonic string) bool {
	return mnemonic == OpMnemonicTXS
}

type TXS struct {
	baseOperation
}

// 0x9A: TXS
func NewTXS() *TXS {
	return &TXS{
		baseOperation{
			code:        OpCodeTXS,
			addressMode: AddrModeImplied,
			mnemonic:    OpMnemonicTXS,
		},
	}
}

func NewTXSBinary(opCode uint8, data io.Reader) (*TXS, error) {
	switch opCode {
	case OpCodeTXS:
		return NewTXS(), nil
	default:
		return nil, InvalidOpCodeError{
			OpCode: opCode,
		}
	}
}

func NewTXSFromBytes(addrMode uint8, arg0 uint8, arg1 uint8) (*TXS, error) {
	switch addrMode {
	case AddrModeImplied:
		return NewTXS(), nil
	default:
		return nil, InvalidAddressModeError{
			AddressMode: addrMode,
		}
	}
}

func (op TXS) Cycles() uint8 {
	switch op.AddressMode() {
	case AddrModeImplied:
		return 2
	default:
		log.Printf("cannot calculate number of cycles for %#v: invalid address mode (%v)", op, op.AddressMode())
		return 0
	}
}

func (op TXS) ExecuteIn(env OperationEnvironment) (uint8, error) {
	cycles := op.Cycles()

	env.SetStackPointer(env.GetIndexX())

	env.IncrementProgramCounter(op.Size())

	return cycles, nil
}
