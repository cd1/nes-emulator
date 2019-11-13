package cpu

import (
	"encoding/binary"
	"io"
	"log"

	"github.com/cd1/nes-emulator/util"
)

const (
	OpMnemonicJMP = "JMP"

	OpCodeJMPAbsolute = 0x4C
	OpCodeJMPIndirect = 0x6C
)

func IsOpCodeValidJMP(opCode uint8) bool {
	return opCode == OpCodeJMPAbsolute ||
		opCode == OpCodeJMPIndirect
}

func IsMnemonicValidJMP(mnemonic string) bool {
	return mnemonic == OpMnemonicJMP
}

type JMP struct {
	baseOperation
}

// 0x4C: JMP $NNNN
func NewJMPAbsolute(absoluteAddress uint16) *JMP {
	return &JMP{
		baseOperation{
			code:        OpCodeJMPAbsolute,
			addressMode: AddrModeAbsolute,
			mnemonic:    OpMnemonicJMP,
			args:        util.BreakWordIntoBytes(absoluteAddress),
		},
	}
}

// 0x6C: JMP ($NNNN)
func NewJMPIndirect(indirectAddress uint16) *JMP {
	return &JMP{
		baseOperation{
			code:        OpCodeJMPIndirect,
			addressMode: AddrModeIndirect,
			mnemonic:    OpMnemonicJMP,
			args:        util.BreakWordIntoBytes(indirectAddress),
		},
	}
}

func NewJMPBinary(opCode uint8, data io.Reader) (*JMP, error) {
	switch opCode {
	case OpCodeJMPAbsolute:
		var addr uint16

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewJMPAbsolute(addr), nil
	case OpCodeJMPIndirect:
		var addr uint16

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewJMPIndirect(addr), nil
	default:
		return nil, InvalidOpCodeError{
			OpCode: opCode,
		}
	}
}

func NewJMPFromBytes(addrMode uint8, arg0 uint8, arg1 uint8) (*JMP, error) {
	switch addrMode {
	case AddrModeAbsolute:
		return NewJMPAbsolute(util.JoinBytesInWord([]uint8{arg0, arg1})), nil
	case AddrModeIndirect:
		return NewJMPIndirect(util.JoinBytesInWord([]uint8{arg0, arg1})), nil
	default:
		return nil, InvalidAddressModeError{
			AddressMode: addrMode,
		}
	}
}
func (op JMP) Cycles() uint8 {
	switch op.AddressMode() {
	case AddrModeAbsolute:
		return 3
	case AddrModeIndirect:
		return 5
	default:
		log.Printf("cannot calculate number of cycles for %#v: invalid address mode (%v)", op, op.AddressMode())
		return 0
	}
}

func (op JMP) ExecuteIn(env OperationEnvironment) (uint8, error) {
	cycles := op.Cycles()
	address, _, _ := env.FetchOperand(op)

	env.SetProgramCounter(address)

	return cycles, nil
}
