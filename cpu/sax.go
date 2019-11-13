package cpu

import (
	"encoding/binary"
	"io"
	"log"

	"github.com/cd1/nes-emulator/util"
)

const (
	OpMnemonicSAX = "SAX"

	// unofficial opcodes
	OpCodeUnSAXIndirectX = 0x83
	OpCodeUnSAXZero      = 0x87
	OpCodeUnSAXAbsolute  = 0x8F
	OpCodeUnSAXZeroY     = 0x97
)

func IsOpCodeValidSAX(opCode uint8) bool {
	return opCode == OpCodeUnSAXIndirectX ||
		opCode == OpCodeUnSAXZero ||
		opCode == OpCodeUnSAXAbsolute ||
		opCode == OpCodeUnSAXZeroY
}

func IsMnemonicValidSAX(mnemonic string) bool {
	return mnemonic == OpMnemonicSAX
}

type SAX struct {
	baseOperation
}

// 0x83: SAX ($NN, X)
func NewUnSAXIndirectX(indirectAddress uint8) *SAX {
	return &SAX{
		baseOperation{
			code:        OpCodeUnSAXIndirectX,
			addressMode: AddrModeIndirectX,
			mnemonic:    OpMnemonicSAX,
			args:        []uint8{indirectAddress},
			unofficial:  true,
		},
	}
}

// 0x87: SAX $NN
func NewUnSAXZero(zeroAddress uint8) *SAX {
	return &SAX{
		baseOperation{
			code:        OpCodeUnSAXZero,
			addressMode: AddrModeZero,
			mnemonic:    OpMnemonicSAX,
			args:        []uint8{zeroAddress},
			unofficial:  true,
		},
	}
}

// 0x8F: SAX $NNNN
func NewUnSAXAbsolute(absoluteAddress uint16) *SAX {
	return &SAX{
		baseOperation{
			code:        OpCodeUnSAXAbsolute,
			addressMode: AddrModeAbsolute,
			mnemonic:    OpMnemonicSAX,
			args:        util.BreakWordIntoBytes(absoluteAddress),
			unofficial:  true,
		},
	}
}

// 0x97: SAX $NN, Y
func NewUnSAXZeroY(zeroAddress uint8) *SAX {
	return &SAX{
		baseOperation{
			code:        OpCodeUnSAXZeroY,
			addressMode: AddrModeZeroY,
			mnemonic:    OpMnemonicSAX,
			args:        []uint8{zeroAddress},
			unofficial:  true,
		},
	}
}

func NewSAXBinary(opCode uint8, data io.Reader) (*SAX, error) {
	switch opCode {
	case OpCodeUnSAXIndirectX:
		var addr uint8

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewUnSAXIndirectX(addr), nil
	case OpCodeUnSAXZero:
		var addr uint8

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewUnSAXZero(addr), nil
	case OpCodeUnSAXAbsolute:
		var addr uint16

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewUnSAXAbsolute(addr), nil
	case OpCodeUnSAXZeroY:
		var addr uint8

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewUnSAXZeroY(addr), nil
	default:
		return nil, InvalidOpCodeError{opCode}
	}
}

func NewSAXFromBytes(addrMode uint8, arg0 uint8, arg1 uint8) (*SAX, error) {
	switch addrMode {
	case AddrModeIndirectX:
		return NewUnSAXIndirectX(arg0), nil
	case AddrModeZero, AddrModeRelative:
		return NewUnSAXZero(arg0), nil
	case AddrModeAbsolute:
		return NewUnSAXAbsolute(util.JoinBytesInWord([]uint8{arg0, arg1})), nil
	case AddrModeZeroY:
		return NewUnSAXZeroY(arg0), nil
	default:
		return nil, InvalidAddressModeError{
			AddressMode: addrMode,
		}
	}
}

func (op SAX) Cycles() uint8 {
	switch op.AddressMode() {
	case AddrModeIndirectX:
		return 6
	case AddrModeZero:
		return 3
	case AddrModeAbsolute:
		return 4
	case AddrModeZeroY:
		return 4
	default:
		log.Printf("cannot calculate number of cycles for %#v: invalid address mode (%v)", op, op.AddressMode())
		return 0
	}
}

func (op SAX) ExecuteIn(env OperationEnvironment) (uint8, error) {
	cycles := op.Cycles()
	address, _, pageCrossed := env.FetchOperand(op)

	if pageCrossed {
		cycles++
	}

	result := env.GetAccumulator() & env.GetIndexX()
	env.WriteByte(address, result)

	env.IncrementProgramCounter(op.Size())

	return cycles, nil
}
