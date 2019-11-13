package cpu

import (
	"encoding/binary"
	"io"
	"log"
)

const (
	OpMnemonicBCS = "BCS"

	OpCodeBCS = 0xB0
)

func IsOpCodeValidBCS(opCode uint8) bool {
	return opCode == OpCodeBCS
}

func IsMnemonicValidBCS(mnemonic string) bool {
	return mnemonic == OpMnemonicBCS
}

type BCS struct {
	baseOperation
}

// 0xB0: BCS $NN
func NewBCS(relativeAddress uint8) *BCS {
	return &BCS{
		baseOperation{
			code:        OpCodeBCS,
			addressMode: AddrModeRelative,
			mnemonic:    OpMnemonicBCS,
			args:        []uint8{relativeAddress},
		},
	}
}

func NewBCSBinary(opCode uint8, data io.Reader) (*BCS, error) {
	switch opCode {
	case OpCodeBCS:
		var addr uint8

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewBCS(addr), nil
	default:
		return nil, InvalidOpCodeError{
			OpCode: opCode,
		}
	}
}

func NewBCSFromBytes(addrMode uint8, arg0 uint8, arg1 uint8) (*BCS, error) {
	switch addrMode {
	case AddrModeZero, AddrModeRelative:
		return NewBCS(arg0), nil
	default:
		return nil, InvalidAddressModeError{
			AddressMode: addrMode,
		}
	}
}

func (op BCS) Cycles() uint8 {
	switch op.AddressMode() {
	case AddrModeRelative:
		return 2
	default:
		log.Printf("cannot calculate number of cycles for %#v: invalid address mode (%v)", op, op.AddressMode())
		return 0
	}
}

func (op BCS) ExecuteIn(env OperationEnvironment) (uint8, error) {
	cycles := op.Cycles()
	_, operand, pageCrossed := env.FetchOperand(op)

	pcIncrement := op.Size()

	if env.IsStatusCarry() {
		pcIncrement += operand
		cycles++

		if pageCrossed {
			cycles++
		}
	}

	env.IncrementProgramCounter(pcIncrement)

	return cycles, nil
}
