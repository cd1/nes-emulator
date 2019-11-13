package cpu

import (
	"encoding/binary"
	"io"
	"log"
)

const (
	OpMnemonicBVC = "BVC"

	OpCodeBVC = 0x50
)

func IsOpCodeValidBVC(opCode uint8) bool {
	return opCode == OpCodeBVC
}

func IsMnemonicValidBVC(mnemonic string) bool {
	return mnemonic == OpMnemonicBVC
}

type BVC struct {
	baseOperation
}

// 0x50: BVC $NN
func NewBVC(relativeAddress uint8) *BVC {
	return &BVC{
		baseOperation{
			code:        OpCodeBVC,
			addressMode: AddrModeRelative,
			mnemonic:    OpMnemonicBVC,
			args:        []uint8{relativeAddress},
		},
	}
}

func NewBVCBinary(opCode uint8, data io.Reader) (*BVC, error) {
	switch opCode {
	case OpCodeBVC:
		var addr uint8

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewBVC(addr), nil
	default:
		return nil, InvalidOpCodeError{
			OpCode: opCode,
		}
	}
}

func NewBVCFromBytes(addrMode uint8, arg0 uint8, arg1 uint8) (*BVC, error) {
	switch addrMode {
	case AddrModeZero, AddrModeRelative:
		return NewBVC(arg0), nil
	default:
		return nil, InvalidAddressModeError{
			AddressMode: addrMode,
		}
	}
}

func (op BVC) Cycles() uint8 {
	switch op.AddressMode() {
	case AddrModeRelative:
		return 2
	default:
		log.Printf("cannot calculate number of cycles for %#v: invalid address mode (%v)", op, op.AddressMode())
		return 0
	}
}

func (op BVC) ExecuteIn(env OperationEnvironment) (uint8, error) {
	cycles := op.Cycles()
	_, operand, pageCrossed := env.FetchOperand(op)

	pcIncrement := op.Size()

	if !env.IsStatusOverflow() {
		pcIncrement += operand
		cycles++

		if pageCrossed {
			cycles++
		}
	}

	env.IncrementProgramCounter(pcIncrement)

	return cycles, nil
}
