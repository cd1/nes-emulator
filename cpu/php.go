package cpu

import (
	"io"
	"log"
)

const (
	OpMnemonicPHP = "PHP"

	OpCodePHP = 0x08
)

func IsOpCodeValidPHP(opCode uint8) bool {
	return opCode == OpCodePHP
}

func IsMnemonicValidPHP(mnemonic string) bool {
	return mnemonic == OpMnemonicPHP
}

type PHP struct {
	baseOperation
}

// 0x08: PHP
func NewPHP() *PHP {
	return &PHP{
		baseOperation{
			code:        OpCodePHP,
			addressMode: AddrModeImplied,
			mnemonic:    OpMnemonicPHP,
		},
	}
}

func NewPHPBinary(opCode uint8, data io.Reader) (*PHP, error) {
	switch opCode {
	case OpCodePHP:
		return NewPHP(), nil
	default:
		return nil, InvalidOpCodeError{
			OpCode: opCode,
		}
	}
}

func NewPHPFromBytes(addrMode uint8, arg0 uint8, arg1 uint8) (*PHP, error) {
	switch addrMode {
	case AddrModeImplied:
		return NewPHP(), nil
	default:
		return nil, InvalidAddressModeError{
			AddressMode: addrMode,
		}
	}
}

func (op PHP) Cycles() uint8 {
	switch op.AddressMode() {
	case AddrModeImplied:
		return 3
	default:
		log.Printf("cannot calculate number of cycles for %#v: invalid address mode (%v)", op, op.AddressMode())
		return 0
	}
}

func (op PHP) ExecuteIn(env OperationEnvironment) (uint8, error) {
	cycles := op.Cycles()

	// TODO: put this login in GetStatus? but what about /IRQ and /NMI?
	hasBreak := env.IsStatusBreak()
	if !hasBreak {
		env.SetStatusBreak(true)
	}
	env.PushByteToStack(env.GetStatus())
	if !hasBreak {
		env.SetStatusBreak(false)
	}

	env.IncrementProgramCounter(op.Size())

	return cycles, nil
}
