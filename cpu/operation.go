package cpu

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
	unofficial  bool
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

func (op baseOperation) IsUnofficial() bool {
	return op.unofficial
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

func (op baseOperation) StringWithEnv(env OperationEnvironment) string {
	var mnemonic string

	if op.IsUnofficial() {
		mnemonic = "*" + op.Mnemonic()
	} else {
		mnemonic = " " + op.Mnemonic()
	}

	var address uint16
	var operand uint8

	hasEnv := (env != nil)

	if hasEnv {
		address, operand, _ = env.FetchOperand(op)
	}

	switch op.AddressMode() {
	case AddrModeAbsolute:
		// JMP and JSR use absolute address but they don't need to read the content in that location
		if hasEnv && !IsOpCodeValidJMP(op.Code()) && !IsOpCodeValidJSR(op.Code()) {
			return fmt.Sprintf("%v $%04X = %02X", mnemonic, op.WordArg(), operand)
		}
		return fmt.Sprintf("%v $%04X", mnemonic, op.WordArg())
	case AddrModeAbsoluteX:
		if hasEnv {
			return fmt.Sprintf("%v $%04X,X @ %04X = %02X", mnemonic, op.WordArg(), address, operand)
		}
		return fmt.Sprintf("%v $%04X,X", mnemonic, op.WordArg())
	case AddrModeAbsoluteY:
		if hasEnv {
			return fmt.Sprintf("%v $%04X,Y @ %04X = %02X", mnemonic, op.WordArg(), address, operand)
		}
		return fmt.Sprintf("%v $%04X,Y", mnemonic, op.WordArg())
	case AddrModeAccumulator:
		return fmt.Sprintf("%v A", mnemonic)
	case AddrModeImmediate:
		return fmt.Sprintf("%v #$%02X", mnemonic, op.ByteArg())
	case AddrModeImplied:
		return mnemonic
	case AddrModeIndirect:
		if hasEnv {
			return fmt.Sprintf("%v ($%04X) = %04X", mnemonic, op.WordArg(), address)
		}
		return fmt.Sprintf("%v ($%04X)", mnemonic, op.WordArg())
	case AddrModeIndirectX:
		if hasEnv {
			return fmt.Sprintf("%v ($%02X,X) @ %02X = %04X = %02X", mnemonic, op.ByteArg(), uint16(op.ByteArg()+env.GetIndexX()), address, operand)
		}
		return fmt.Sprintf("%v ($%02X,X)", mnemonic, op.ByteArg())
	case AddrModeIndirectY:
		if hasEnv {
			return fmt.Sprintf("%v ($%02X),Y = %04X @ %04X = %02X", mnemonic, op.ByteArg(), env.ReadWordSamePage(uint16(op.ByteArg())), env.ReadWordSamePage(uint16(op.ByteArg()))+uint16(env.GetIndexY()), operand)
		}
		return fmt.Sprintf("%v ($%02X),Y", mnemonic, op.ByteArg())
	case AddrModeRelative:
		if hasEnv {
			return fmt.Sprintf("%v $%04X", mnemonic, env.GetProgramCounter()+uint16(int8(op.ByteArg()+op.Size())))
		}
		return fmt.Sprintf("%v $%02X", mnemonic, op.ByteArg())
	case AddrModeZero:
		if hasEnv {
			return fmt.Sprintf("%v $%02X = %02X", mnemonic, op.ByteArg(), operand)
		}
		return fmt.Sprintf("%v $%02X", mnemonic, op.ByteArg())
	case AddrModeZeroX:
		if hasEnv {
			return fmt.Sprintf("%v $%02X,X @ %02X = %02X", mnemonic, op.ByteArg(), address, operand)
		}
		return fmt.Sprintf("%v $%02X,X", mnemonic, op.ByteArg())
	case AddrModeZeroY:
		if hasEnv {
			return fmt.Sprintf("%v $%02X,Y @ %02X = %02X", mnemonic, op.ByteArg(), address, operand)
		}
		return fmt.Sprintf("%v $%02X,Y", mnemonic, op.ByteArg())
	default:
		return fmt.Sprintf("[%#v ; unexpected address mode]", op)
	}
}

func (op baseOperation) String() string {
	return op.StringWithEnv(nil)
}

func (op baseOperation) WordArg() uint16 {
	return util.JoinBytesInWord(op.args)
}

func (op baseOperation) ExecuteIn(env OperationEnvironment) (uint8, error) {
	return 0, nil
}

type Operation interface {
	Code() uint8
	AddressMode() uint8
	Mnemonic() string
	Size() uint8
	ByteArg() uint8
	WordArg() uint16
	StringWithEnv(OperationEnvironment) string
	ExecuteIn(OperationEnvironment) (uint8, error)
}

type OperationEnvironment interface {
	ReadByte(uint16) uint8
	ReadWord(uint16) uint16
	ReadWordSamePage(uint16) uint16
	WriteByte(uint16, uint8)
	WriteWord(uint16, uint16)

	IsStatusCarry() bool
	SetStatusCarry(bool)
	IsStatusZero() bool
	SetStatusZero(bool)
	IsStatusInterrupt() bool
	SetStatusInterrupt(bool)
	IsStatusDecimal() bool
	SetStatusDecimal(bool)
	IsStatusBreak() bool
	SetStatusBreak(bool)
	IsStatusUnused() bool
	SetStatusUnused(bool)
	IsStatusOverflow() bool
	SetStatusOverflow(bool)
	IsStatusNegative() bool
	SetStatusNegative(bool)

	GetAccumulator() uint8
	SetAccumulator(uint8)
	GetIndexX() uint8
	SetIndexX(uint8)
	GetIndexY() uint8
	SetIndexY(uint8)
	GetStackPointer() uint8
	SetStackPointer(uint8)
	GetProgramCounter() uint16
	SetProgramCounter(uint16)
	IncrementProgramCounter(uint8)
	GetStatus() uint8
	SetStatus(uint8)

	PushByteToStack(uint8)
	PushWordToStack(uint16)
	PullByteFromStack() uint8
	PullWordFromStack() uint16

	FetchOperand(Operation) (uint16, uint8, bool)
}
