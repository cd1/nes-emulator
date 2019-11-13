package cpu

import (
	"encoding/binary"
	"io"
	"log"
)

const (
	OpMnemonicBVS = "BVS"

	OpCodeBVS = 0x70
)

func IsOpCodeValidBVS(opCode uint8) bool {
	return opCode == OpCodeBVS
}

func IsMnemonicValidBVS(mnemonic string) bool {
	return mnemonic == OpMnemonicBVS
}

type BVS struct {
	baseOperation
}

// 0x70: BVS $NN
func NewBVS(relativeAddress uint8) *BVS {
	return &BVS{
		baseOperation{
			code:        OpCodeBVS,
			addressMode: AddrModeRelative,
			mnemonic:    OpMnemonicBVS,
			args:        []uint8{relativeAddress},
		},
	}
}

func NewBVSBinary(opCode uint8, data io.Reader) (*BVS, error) {
	switch opCode {
	case OpCodeBVS:
		var addr uint8

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewBVS(addr), nil
	default:
		return nil, InvalidOpCodeError{
			OpCode: opCode,
		}
	}
}

func NewBVSFromBytes(addrMode uint8, arg0 uint8, arg1 uint8) (*BVS, error) {
	switch addrMode {
	case AddrModeZero, AddrModeRelative:
		return NewBVS(arg0), nil
	default:
		return nil, InvalidAddressModeError{
			AddressMode: addrMode,
		}
	}
}

func (op BVS) Cycles() uint8 {
	switch op.AddressMode() {
	case AddrModeRelative:
		return 2
	default:
		log.Printf("cannot calculate number of cycles for %#v: invalid address mode (%v)", op, op.AddressMode())
		return 0
	}
}

func (op BVS) ExecuteIn(env OperationEnvironment) (uint8, error) {
	cycles := op.Cycles()
	_, operand, pageCrossed := env.FetchOperand(op)

	pcIncrement := op.Size()

	if env.IsStatusOverflow() {
		pcIncrement += operand
		cycles++

		if pageCrossed {
			cycles++
		}
	}

	env.IncrementProgramCounter(pcIncrement)

	return cycles, nil
}
