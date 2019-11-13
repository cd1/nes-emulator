package cpu

import (
	"io"
	"log"
)

const (
	OpMnemonicRTI = "RTI"

	OpCodeRTI = 0x40
)

func IsOpCodeValidRTI(opCode uint8) bool {
	return opCode == OpCodeRTI
}

func IsMnemonicValidRTI(mnemonic string) bool {
	return mnemonic == OpMnemonicRTI
}

type RTI struct {
	baseOperation
}

// 0x40: RTI
func NewRTI() *RTI {
	return &RTI{
		baseOperation{
			code:        OpCodeRTI,
			addressMode: AddrModeImplied,
			mnemonic:    OpMnemonicRTI,
		},
	}
}

func NewRTIBinary(opCode uint8, data io.Reader) (*RTI, error) {
	switch opCode {
	case OpCodeRTI:
		return NewRTI(), nil
	default:
		return nil, InvalidOpCodeError{
			OpCode: opCode,
		}
	}
}

func NewRTIFromBytes(addrMode uint8, arg0 uint8, arg1 uint8) (*RTI, error) {
	switch addrMode {
	case AddrModeImplied:
		return NewRTI(), nil
	default:
		return nil, InvalidAddressModeError{
			AddressMode: addrMode,
		}
	}
}

func (op RTI) Cycles() uint8 {
	switch op.AddressMode() {
	case AddrModeImplied:
		return 6
	default:
		log.Printf("cannot calculate number of cycles for %#v: invalid address mode (%v)", op, op.AddressMode())
		return 0
	}
}

func (op RTI) ExecuteIn(env OperationEnvironment) (uint8, error) {
	cycles := op.Cycles()

	env.SetStatus(env.PullByteFromStack())
	env.SetProgramCounter(env.PullWordFromStack())

	return cycles, nil
}
