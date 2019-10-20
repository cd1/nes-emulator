package operation

import (
	"fmt"

	"github.com/cd1/nes-emulator/util"
)

const (
	AddrModeAccumulator = iota
	AddrModeImmediate
	AddrModeImplied
	AddrModeRelative
	AddrModeAbsolute
	AddrModeZero
	AddrModeIndirect
	AddrModeAbsoluteX
	AddrModeAbsoluteY
	AddrModeZeroX
	AddrModeZeroY
	AddrModeIndirectX
	AddrModeIndirectY
)

func AddressModeString(addressMode uint8) string {
	switch addressMode {
	case AddrModeAccumulator:
		return "accumulator"
	case AddrModeImmediate:
		return "immediate"
	case AddrModeImplied:
		return "implied"
	default:
		return fmt.Sprintf("[address mode = %v]", addressMode)
	}
}

type baseOperation struct {
	code        uint8
	addressMode uint8
	mnemonic    string
	args        []uint8
}

func (op baseOperation) AddressMode() uint8 {
	return op.addressMode
}

func (op baseOperation) ByteArg() uint8 {
	return op.args[0]
}

func (op baseOperation) Code() uint8 {
	return op.code
}

func (op baseOperation) Mnemonic() string {
	return op.mnemonic
}

func (op baseOperation) Size() uint8 {
	switch op.AddressMode() {
	case AddrModeAccumulator, AddrModeImplied:
		return 1
	case AddrModeImmediate, AddrModeRelative, AddrModeZero, AddrModeZeroX, AddrModeZeroY, AddrModeIndirectX, AddrModeIndirectY:
		return 2
	case AddrModeAbsolute, AddrModeIndirect, AddrModeAbsoluteX, AddrModeAbsoluteY:
		return 3
	default:
		return 1
	}
}

func (op baseOperation) String() string {
	switch op.AddressMode() {
	case AddrModeAccumulator:
		return fmt.Sprintf("%v A", op.Mnemonic())
	case AddrModeImmediate:
		return fmt.Sprintf("%v #$%02X", op.Mnemonic(), op.ByteArg())
	case AddrModeImplied:
		return op.Mnemonic()
	case AddrModeRelative, AddrModeZero:
		return fmt.Sprintf("%v $%02X", op.Mnemonic(), op.ByteArg())
	case AddrModeAbsolute:
		return fmt.Sprintf("%v $%04X", op.Mnemonic(), op.WordArg())
	case AddrModeIndirect:
		return fmt.Sprintf("%v ($%04X)", op.Mnemonic(), op.WordArg())
	case AddrModeAbsoluteX:
		return fmt.Sprintf("%v $%04X, X", op.Mnemonic(), op.WordArg())
	case AddrModeAbsoluteY:
		return fmt.Sprintf("%v $%04X, Y", op.Mnemonic(), op.WordArg())
	case AddrModeZeroX:
		return fmt.Sprintf("%v $%02X, X", op.Mnemonic(), op.ByteArg())
	case AddrModeZeroY:
		return fmt.Sprintf("%v $%02X, Y", op.Mnemonic(), op.ByteArg())
	case AddrModeIndirectX:
		return fmt.Sprintf("%v ($%02X, X)", op.Mnemonic(), op.ByteArg())
	case AddrModeIndirectY:
		return fmt.Sprintf("%v ($%02X), Y", op.Mnemonic(), op.ByteArg())
	default:
		return fmt.Sprintf("[%#v ; unexpected address mode]", op)
	}
}

func (op baseOperation) WordArg() uint16 {
	return util.JoinBytesInWord(op.args)
}

type Operation interface {
	Code() uint8
	AddressMode() uint8
	Mnemonic() string
	Size() uint8
	ByteArg() uint8
	WordArg() uint16
	// Execute()
}
