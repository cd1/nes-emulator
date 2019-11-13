package cpu

import (
	"io"
	"log"
)

const (
	OpMnemonicRTS = "RTS"

	OpCodeRTS = 0x60
)

func IsOpCodeValidRTS(opCode uint8) bool {
	return opCode == OpCodeRTS
}

func IsMnemonicValidRTS(mnemonic string) bool {
	return mnemonic == OpMnemonicRTS
}

type RTS struct {
	baseOperation
}

// 0x60: RTS
func NewRTS() *RTS {
	return &RTS{
		baseOperation{
			code:        OpCodeRTS,
			addressMode: AddrModeImplied,
			mnemonic:    OpMnemonicRTS,
		},
	}
}

func NewRTSBinary(opCode uint8, data io.Reader) (*RTS, error) {
	switch opCode {
	case OpCodeRTS:
		return NewRTS(), nil
	default:
		return nil, InvalidOpCodeError{
			OpCode: opCode,
		}
	}
}

func NewRTSFromBytes(addrMode uint8, arg0 uint8, arg1 uint8) (*RTS, error) {
	switch addrMode {
	case AddrModeImplied:
		return NewRTS(), nil
	default:
		return nil, InvalidAddressModeError{
			AddressMode: addrMode,
		}
	}
}

func (op RTS) Cycles() uint8 {
	switch op.AddressMode() {
	case AddrModeImplied:
		return 6
	default:
		log.Printf("cannot calculate number of cycles for %#v: invalid address mode (%v)", op, op.AddressMode())
		return 0
	}
}

func (op RTS) ExecuteIn(env OperationEnvironment) (uint8, error) {
	cycles := op.Cycles()

	env.SetProgramCounter(env.PullWordFromStack() + 1)

	return cycles, nil
}
