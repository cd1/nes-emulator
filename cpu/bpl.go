package cpu

import (
	"encoding/binary"
	"io"
	"log"
)

const (
	OpMnemonicBPL = "BPL"

	OpCodeBPL = 0x10
)

func IsOpCodeValidBPL(opCode uint8) bool {
	return opCode == OpCodeBPL
}

func IsMnemonicValidBPL(mnemonic string) bool {
	return mnemonic == OpMnemonicBPL
}

type BPL struct {
	baseOperation
}

// 0x10: BPL $NN
func NewBPL(relativeAddress uint8) *BPL {
	return &BPL{
		baseOperation{
			code:        OpCodeBPL,
			addressMode: AddrModeRelative,
			mnemonic:    OpMnemonicBPL,
			args:        []uint8{relativeAddress},
		},
	}
}

func NewBPLBinary(opCode uint8, data io.Reader) (*BPL, error) {
	switch opCode {
	case OpCodeBPL:
		var addr uint8

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewBPL(addr), nil
	default:
		return nil, InvalidOpCodeError{
			OpCode: opCode,
		}
	}
}

func NewBPLFromBytes(addrMode uint8, arg0 uint8, arg1 uint8) (*BPL, error) {
	switch addrMode {
	case AddrModeZero, AddrModeRelative:
		return NewBPL(arg0), nil
	default:
		return nil, InvalidAddressModeError{
			AddressMode: addrMode,
		}
	}
}

func (op BPL) Cycles() uint8 {
	switch op.AddressMode() {
	case AddrModeRelative:
		return 2
	default:
		log.Printf("cannot calculate number of cycles for %#v: invalid address mode (%v)", op, op.AddressMode())
		return 0
	}
}

func (op BPL) ExecuteIn(env OperationEnvironment) (uint8, error) {
	cycles := op.Cycles()
	_, operand, pageCrossed := env.FetchOperand(op)

	pcIncrement := op.Size()

	if !env.IsStatusNegative() {
		pcIncrement += operand
		cycles++

		if pageCrossed {
			cycles++
		}
	}

	env.IncrementProgramCounter(pcIncrement)

	return cycles, nil
}
