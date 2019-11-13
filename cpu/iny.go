package cpu

import (
	"io"
	"log"
)

const (
	OpMnemonicINY = "INY"

	OpCodeINY = 0xC8
)

func IsOpCodeValidINY(opCode uint8) bool {
	return opCode == OpCodeINY
}

func IsMnemonicValidINY(mnemonic string) bool {
	return mnemonic == OpMnemonicINY
}

type INY struct {
	baseOperation
}

// 0xC8: INY
func NewINY() *INY {
	return &INY{
		baseOperation{
			code:        OpCodeINY,
			addressMode: AddrModeImplied,
			mnemonic:    OpMnemonicINY,
		},
	}
}

func NewINYBinary(opCode uint8, data io.Reader) (*INY, error) {
	switch opCode {
	case OpCodeINY:
		return NewINY(), nil
	default:
		return nil, InvalidOpCodeError{
			OpCode: opCode,
		}
	}
}

func NewINYFromBytes(addrMode uint8, arg0 uint8, arg1 uint8) (*INY, error) {
	switch addrMode {
	case AddrModeImplied:
		return NewINY(), nil
	default:
		return nil, InvalidAddressModeError{
			AddressMode: addrMode,
		}
	}
}

func (op INY) Cycles() uint8 {
	switch op.AddressMode() {
	case AddrModeImplied:
		return 2
	default:
		log.Printf("cannot calculate number of cycles for %#v: invalid address mode (%v)", op, op.AddressMode())
		return 0
	}
}

func (op INY) ExecuteIn(env OperationEnvironment) (uint8, error) {
	cycles := op.Cycles()

	newY := env.GetIndexY() + 1
	env.SetIndexY(newY)

	env.SetStatusZero(newY == 0x00)
	env.SetStatusNegative(newY&0x80 != 0x00)

	env.IncrementProgramCounter(op.Size())

	return cycles, nil
}
