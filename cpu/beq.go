package cpu

import (
	"encoding/binary"
	"io"
	"log"
)

const (
	OpMnemonicBEQ = "BEQ"

	OpCodeBEQ = 0xF0
)

func IsOpCodeValidBEQ(opCode uint8) bool {
	return opCode == OpCodeBEQ
}

func IsMnemonicValidBEQ(mnemonic string) bool {
	return mnemonic == OpMnemonicBEQ
}

type BEQ struct {
	baseOperation
}

// 0xF0: BEQ $NN
func NewBEQ(relativeAddress uint8) *BEQ {
	return &BEQ{
		baseOperation{
			code:        OpCodeBEQ,
			addressMode: AddrModeRelative,
			mnemonic:    OpMnemonicBEQ,
			args:        []uint8{relativeAddress},
		},
	}
}

func NewBEQBinary(opCode uint8, data io.Reader) (*BEQ, error) {
	switch opCode {
	case OpCodeBEQ:
		var addr uint8

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewBEQ(addr), nil
	default:
		return nil, InvalidOpCodeError{
			OpCode: opCode,
		}
	}
}

func NewBEQFromBytes(addrMode uint8, arg0 uint8, arg1 uint8) (*BEQ, error) {
	switch addrMode {
	case AddrModeZero, AddrModeRelative:
		return NewBEQ(arg0), nil
	default:
		return nil, InvalidAddressModeError{
			AddressMode: addrMode,
		}
	}
}

func (op BEQ) Cycles() uint8 {
	switch op.AddressMode() {
	case AddrModeRelative:
		return 2
	default:
		log.Printf("cannot calculate number of cycles for %#v: invalid address mode (%v)", op, op.AddressMode())
		return 0
	}
}

func (op BEQ) ExecuteIn(env OperationEnvironment) (uint8, error) {
	cycles := op.Cycles()
	_, operand, pageCrossed := env.FetchOperand(op)

	pcIncrement := op.Size()

	if env.IsStatusZero() {
		pcIncrement += operand
		cycles++

		if pageCrossed {
			cycles++
		}
	}

	env.IncrementProgramCounter(pcIncrement)

	return cycles, nil
}
