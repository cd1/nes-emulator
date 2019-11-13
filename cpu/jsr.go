package cpu

import (
	"encoding/binary"
	"io"
	"log"

	"github.com/cd1/nes-emulator/util"
)

const (
	OpMnemonicJSR = "JSR"

	OpCodeJSR = 0x20
)

func IsOpCodeValidJSR(opCode uint8) bool {
	return opCode == OpCodeJSR
}

func IsMnemonicValidJSR(mnemonic string) bool {
	return mnemonic == OpMnemonicJSR
}

type JSR struct {
	baseOperation
}

// 0x20: JSR $NNNN
func NewJSR(absoluteAddress uint16) *JSR {
	return &JSR{
		baseOperation{
			code:        OpCodeJSR,
			addressMode: AddrModeAbsolute,
			mnemonic:    OpMnemonicJSR,
			args:        util.BreakWordIntoBytes(absoluteAddress),
		},
	}
}

func NewJSRBinary(opCode uint8, data io.Reader) (*JSR, error) {
	switch opCode {
	case OpCodeJSR:
		var addr uint16

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewJSR(addr), nil
	default:
		return nil, InvalidOpCodeError{
			OpCode: opCode,
		}
	}
}

func NewJSRFromBytes(addrMode uint8, arg0 uint8, arg1 uint8) (*JSR, error) {
	switch addrMode {
	case AddrModeAbsolute:
		return NewJSR(util.JoinBytesInWord([]uint8{arg0, arg1})), nil
	default:
		return nil, InvalidAddressModeError{
			AddressMode: addrMode,
		}
	}
}

func (op JSR) Cycles() uint8 {
	switch op.AddressMode() {
	case AddrModeAbsolute:
		return 6
	default:
		log.Printf("cannot calculate number of cycles for %#v: invalid address mode (%v)", op, op.AddressMode())
		return 0
	}
}

func (op JSR) ExecuteIn(env OperationEnvironment) (uint8, error) {
	cycles := op.Cycles()
	address, _, _ := env.FetchOperand(op)

	env.PushWordToStack(env.GetProgramCounter() + uint16(op.Size()) - 1)
	env.SetProgramCounter(address)

	return cycles, nil
}
