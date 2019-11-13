package cpu

import (
	"encoding/binary"
	"io"
	"log"

	"github.com/cd1/nes-emulator/util"
)

const (
	OpMnemonicNOP = "NOP"

	OpCodeNOP = 0xEA

	// unofficial opcodes
	OpCodeUnNOPZero0      = 0x04
	OpCodeUnNOPAbsolute   = 0x0C
	OpCodeUnNOPZeroX0     = 0x14
	OpCodeUnNOPImplied0   = 0x1A
	OpCodeUnNOPAbsoluteX0 = 0x1C
	OpCodeUnNOPZeroX1     = 0x34
	OpCodeUnNOPImplied1   = 0x3A
	OpCodeUnNOPAbsoluteX1 = 0x3C
	OpCodeUnNOPZero1      = 0x44
	OpCodeUnNOPZeroX2     = 0x54
	OpCodeUnNOPImplied2   = 0x5A
	OpCodeUnNOPAbsoluteX2 = 0x5C
	OpCodeUnNOPZero2      = 0x64
	OpCodeUnNOPZeroX3     = 0x74
	OpCodeUnNOPImplied3   = 0x7A
	OpCodeUnNOPAbsoluteX3 = 0x7C
	OpCodeUnNOPImmediate  = 0x80
	OpCodeUnNOPZeroX4     = 0xD4
	OpCodeUnNOPImplied4   = 0xDA
	OpCodeUnNOPAbsoluteX4 = 0xDC
	OpCodeUnNOPZeroX5     = 0xF4
	OpCodeUnNOPImplied5   = 0xFA
	OpCodeUnNOPAbsoluteX5 = 0xFC
)

func IsOpCodeValidNOP(opCode uint8) bool {
	return opCode == OpCodeNOP ||
		opCode == OpCodeUnNOPZero0 ||
		opCode == OpCodeUnNOPAbsolute ||
		opCode == OpCodeUnNOPZeroX0 ||
		opCode == OpCodeUnNOPImplied0 ||
		opCode == OpCodeUnNOPAbsoluteX0 ||
		opCode == OpCodeUnNOPZeroX1 ||
		opCode == OpCodeUnNOPImplied1 ||
		opCode == OpCodeUnNOPAbsoluteX1 ||
		opCode == OpCodeUnNOPZero1 ||
		opCode == OpCodeUnNOPZeroX2 ||
		opCode == OpCodeUnNOPImplied2 ||
		opCode == OpCodeUnNOPAbsoluteX2 ||
		opCode == OpCodeUnNOPZero2 ||
		opCode == OpCodeUnNOPZeroX3 ||
		opCode == OpCodeUnNOPImplied3 ||
		opCode == OpCodeUnNOPAbsoluteX3 ||
		opCode == OpCodeUnNOPImmediate ||
		opCode == OpCodeUnNOPZeroX4 ||
		opCode == OpCodeUnNOPImplied4 ||
		opCode == OpCodeUnNOPAbsoluteX4 ||
		opCode == OpCodeUnNOPZeroX5 ||
		opCode == OpCodeUnNOPImplied5 ||
		opCode == OpCodeUnNOPAbsoluteX5
}

func IsMnemonicValidNOP(mnemonic string) bool {
	return mnemonic == OpMnemonicNOP
}

type NOP struct {
	baseOperation
}

// 0xEA: NOP
func NewNOP() *NOP {
	return &NOP{
		baseOperation{
			code:        OpCodeNOP,
			addressMode: AddrModeImplied,
			mnemonic:    OpMnemonicNOP,
		},
	}
}

// 0x04: NOP $NN
func NewUnNOPZero0(zeroAddress uint8) *NOP {
	return &NOP{
		baseOperation{
			code:        OpCodeUnNOPZero0,
			addressMode: AddrModeZero,
			mnemonic:    OpMnemonicNOP,
			args:        []uint8{zeroAddress},
			unofficial:  true,
		},
	}
}

// 0x0C: NOP $NNNN
func NewUnNOPAbsolute(absoluteAddress uint16) *NOP {
	return &NOP{
		baseOperation{
			code:        OpCodeUnNOPAbsolute,
			addressMode: AddrModeAbsolute,
			mnemonic:    OpMnemonicNOP,
			args:        util.BreakWordIntoBytes(absoluteAddress),
			unofficial:  true,
		},
	}
}

// 0x14: NOP $NN, X
func NewUnNOPZeroX0(zeroAddress uint8) *NOP {
	return &NOP{
		baseOperation{
			code:        OpCodeUnNOPZeroX0,
			addressMode: AddrModeZeroX,
			mnemonic:    OpMnemonicNOP,
			args:        []uint8{zeroAddress},
			unofficial:  true,
		},
	}
}

// 0x1A: NOP
func NewUnNOPImplied0() *NOP {
	return &NOP{
		baseOperation{
			code:        OpCodeUnNOPImplied0,
			addressMode: AddrModeImplied,
			mnemonic:    OpMnemonicNOP,
			unofficial:  true,
		},
	}
}

// 0x1C: NOP $NNNN, X
func NewUnNOPAbsoluteX0(absoluteAddress uint16) *NOP {
	return &NOP{
		baseOperation{
			code:        OpCodeUnNOPAbsoluteX0,
			addressMode: AddrModeAbsoluteX,
			mnemonic:    OpMnemonicNOP,
			args:        util.BreakWordIntoBytes(absoluteAddress),
			unofficial:  true,
		},
	}
}

// 0x34: NOP $NN, X
func NewUnNOPZeroX1(zeroAddress uint8) *NOP {
	return &NOP{
		baseOperation{
			code:        OpCodeUnNOPZeroX1,
			addressMode: AddrModeZeroX,
			mnemonic:    OpMnemonicNOP,
			args:        []uint8{zeroAddress},
			unofficial:  true,
		},
	}
}

// 0x3A: NOP
func NewUnNOPImplied1() *NOP {
	return &NOP{
		baseOperation{
			code:        OpCodeUnNOPImplied1,
			addressMode: AddrModeImplied,
			mnemonic:    OpMnemonicNOP,
			unofficial:  true,
		},
	}
}

// 0x3C: NOP $NNNN, X
func NewUnNOPAbsoluteX1(absoluteAddress uint16) *NOP {
	return &NOP{
		baseOperation{
			code:        OpCodeUnNOPAbsoluteX1,
			addressMode: AddrModeAbsoluteX,
			mnemonic:    OpMnemonicNOP,
			args:        util.BreakWordIntoBytes(absoluteAddress),
			unofficial:  true,
		},
	}
}

// 0x44: NOP $NN
func NewUnNOPZero1(zeroAddress uint8) *NOP {
	return &NOP{
		baseOperation{
			code:        OpCodeUnNOPZero1,
			addressMode: AddrModeZero,
			mnemonic:    OpMnemonicNOP,
			args:        []uint8{zeroAddress},
			unofficial:  true,
		},
	}
}

// 0x54: NOP $NN, X
func NewUnNOPZeroX2(zeroAddress uint8) *NOP {
	return &NOP{
		baseOperation{
			code:        OpCodeUnNOPZeroX2,
			addressMode: AddrModeZeroX,
			mnemonic:    OpMnemonicNOP,
			args:        []uint8{zeroAddress},
			unofficial:  true,
		},
	}
}

// 0x5A: NOP
func NewUnNOPImplied2() *NOP {
	return &NOP{
		baseOperation{
			code:        OpCodeUnNOPImplied2,
			addressMode: AddrModeImplied,
			mnemonic:    OpMnemonicNOP,
			unofficial:  true,
		},
	}
}

// 0x5C: NOP $NNNN, X
func NewUnNOPAbsoluteX2(absoluteAddress uint16) *NOP {
	return &NOP{
		baseOperation{
			code:        OpCodeUnNOPAbsoluteX2,
			addressMode: AddrModeAbsoluteX,
			mnemonic:    OpMnemonicNOP,
			args:        util.BreakWordIntoBytes(absoluteAddress),
			unofficial:  true,
		},
	}
}

// 0x64: NOP $NN
func NewUnNOPZero2(zeroAddress uint8) *NOP {
	return &NOP{
		baseOperation{
			code:        OpCodeUnNOPZero2,
			addressMode: AddrModeZero,
			mnemonic:    OpMnemonicNOP,
			args:        []uint8{zeroAddress},
			unofficial:  true,
		},
	}
}

// 0x74: NOP $NN, X
func NewUnNOPZeroX3(zeroAddress uint8) *NOP {
	return &NOP{
		baseOperation{
			code:        OpCodeUnNOPZeroX3,
			addressMode: AddrModeZeroX,
			mnemonic:    OpMnemonicNOP,
			args:        []uint8{zeroAddress},
			unofficial:  true,
		},
	}
}

// 0x7A: NOP
func NewUnNOPImplied3() *NOP {
	return &NOP{
		baseOperation{
			code:        OpCodeUnNOPImplied3,
			addressMode: AddrModeImplied,
			mnemonic:    OpMnemonicNOP,
			unofficial:  true,
		},
	}
}

// 0x7C: NOP $NNNN, X
func NewUnNOPAbsoluteX3(absoluteAddress uint16) *NOP {
	return &NOP{
		baseOperation{
			code:        OpCodeUnNOPAbsoluteX3,
			addressMode: AddrModeAbsoluteX,
			mnemonic:    OpMnemonicNOP,
			args:        util.BreakWordIntoBytes(absoluteAddress),
			unofficial:  true,
		},
	}
}

// 0x80: NOP
func NewUnNOPImmediate(value uint8) *NOP {
	return &NOP{
		baseOperation{
			code:        OpCodeUnNOPImmediate,
			addressMode: AddrModeImmediate,
			mnemonic:    OpMnemonicNOP,
			args:        []uint8{value},
			unofficial:  true,
		},
	}
}

// 0xD4: NOP $NN, X
func NewUnNOPZeroX4(zeroAddress uint8) *NOP {
	return &NOP{
		baseOperation{
			code:        OpCodeUnNOPZeroX4,
			addressMode: AddrModeZeroX,
			mnemonic:    OpMnemonicNOP,
			args:        []uint8{zeroAddress},
			unofficial:  true,
		},
	}
}

// 0xDA: NOP
func NewUnNOPImplied4() *NOP {
	return &NOP{
		baseOperation{
			code:        OpCodeUnNOPImplied4,
			addressMode: AddrModeImplied,
			mnemonic:    OpMnemonicNOP,
			unofficial:  true,
		},
	}
}

// 0xDC: NOP $NNNN, X
func NewUnNOPAbsoluteX4(absoluteAddress uint16) *NOP {
	return &NOP{
		baseOperation{
			code:        OpCodeUnNOPAbsoluteX4,
			addressMode: AddrModeAbsoluteX,
			mnemonic:    OpMnemonicNOP,
			args:        util.BreakWordIntoBytes(absoluteAddress),
			unofficial:  true,
		},
	}
}

// 0xF4: NOP $NN, X
func NewUnNOPZeroX5(zeroAddress uint8) *NOP {
	return &NOP{
		baseOperation{
			code:        OpCodeUnNOPZeroX5,
			addressMode: AddrModeZeroX,
			mnemonic:    OpMnemonicNOP,
			args:        []uint8{zeroAddress},
			unofficial:  true,
		},
	}
}

// 0xFA: NOP
func NewUnNOPImplied5() *NOP {
	return &NOP{
		baseOperation{
			code:        OpCodeUnNOPImplied5,
			addressMode: AddrModeImplied,
			mnemonic:    OpMnemonicNOP,
			unofficial:  true,
		},
	}
}

// 0xFC: NOP $NNNN, X
func NewUnNOPAbsoluteX5(absoluteAddress uint16) *NOP {
	return &NOP{
		baseOperation{
			code:        OpCodeUnNOPAbsoluteX5,
			addressMode: AddrModeAbsoluteX,
			mnemonic:    OpMnemonicNOP,
			args:        util.BreakWordIntoBytes(absoluteAddress),
			unofficial:  true,
		},
	}
}

func NewNOPBinary(opCode uint8, data io.Reader) (*NOP, error) {
	switch opCode {
	case OpCodeNOP:
		return NewNOP(), nil
	case OpCodeUnNOPZero0:
		var addr uint8

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewUnNOPZero0(addr), nil
	case OpCodeUnNOPAbsolute:
		var addr uint16

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewUnNOPAbsolute(addr), nil
	case OpCodeUnNOPZeroX0:
		var addr uint8

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewUnNOPZeroX0(addr), nil
	case OpCodeUnNOPImplied0:
		return NewUnNOPImplied0(), nil
	case OpCodeUnNOPAbsoluteX0:
		var addr uint16

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewUnNOPAbsoluteX0(addr), nil
	case OpCodeUnNOPZeroX1:
		var addr uint8

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewUnNOPZeroX1(addr), nil
	case OpCodeUnNOPImplied1:
		return NewUnNOPImplied1(), nil
	case OpCodeUnNOPAbsoluteX1:
		var addr uint16

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewUnNOPAbsoluteX1(addr), nil
	case OpCodeUnNOPZero1:
		var addr uint8

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewUnNOPZero1(addr), nil
	case OpCodeUnNOPZeroX2:
		var addr uint8

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewUnNOPZeroX2(addr), nil
	case OpCodeUnNOPImplied2:
		return NewUnNOPImplied2(), nil
	case OpCodeUnNOPAbsoluteX2:
		var addr uint16

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewUnNOPAbsoluteX2(addr), nil
	case OpCodeUnNOPZero2:
		var addr uint8

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewUnNOPZero2(addr), nil
	case OpCodeUnNOPZeroX3:
		var addr uint8

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewUnNOPZeroX3(addr), nil
	case OpCodeUnNOPImplied3:
		return NewUnNOPImplied3(), nil
	case OpCodeUnNOPAbsoluteX3:
		var addr uint16

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewUnNOPAbsoluteX3(addr), nil
	case OpCodeUnNOPImmediate:
		var addr uint8

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewUnNOPImmediate(addr), nil
	case OpCodeUnNOPZeroX4:
		var addr uint8

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewUnNOPZeroX4(addr), nil
	case OpCodeUnNOPImplied4:
		return NewUnNOPImplied4(), nil
	case OpCodeUnNOPAbsoluteX4:
		var addr uint16

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewUnNOPAbsoluteX4(addr), nil
	case OpCodeUnNOPZeroX5:
		var addr uint8

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewUnNOPZeroX5(addr), nil
	case OpCodeUnNOPImplied5:
		return NewUnNOPImplied5(), nil
	case OpCodeUnNOPAbsoluteX5:
		var addr uint16

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewUnNOPAbsoluteX5(addr), nil
	default:
		return nil, InvalidOpCodeError{
			OpCode: opCode,
		}
	}
}

func NewNOPFromBytes(addrMode uint8, arg0 uint8, arg1 uint8) (*NOP, error) {
	switch addrMode {
	case AddrModeImplied:
		return NewNOP(), nil
	case AddrModeZero, AddrModeRelative:
		return NewUnNOPZero0(arg0), nil
	case AddrModeAbsolute:
		return NewUnNOPAbsolute(util.JoinBytesInWord([]uint8{arg0, arg1})), nil
	case AddrModeZeroX:
		return NewUnNOPZeroX0(arg0), nil
	case AddrModeAbsoluteX:
		return NewUnNOPAbsoluteX0(util.JoinBytesInWord([]uint8{arg0, arg1})), nil
	case AddrModeImmediate:
		return NewUnNOPImmediate(arg0), nil
	default:
		return nil, InvalidAddressModeError{
			AddressMode: addrMode,
		}
	}
}

func (op NOP) Cycles() uint8 {
	switch op.AddressMode() {
	case AddrModeImplied:
		return 2
	case AddrModeZero:
		return 3
	case AddrModeAbsolute:
		return 4
	case AddrModeZeroX:
		return 4
	case AddrModeAbsoluteX:
		return 4
	case AddrModeImmediate:
		return 2
	default:
		log.Printf("cannot calculate number of cycles for %#v: invalid address mode (%v)", op, op.AddressMode())
		return 0
	}
}

func (op NOP) ExecuteIn(env OperationEnvironment) (uint8, error) {
	cycles := op.Cycles()
	_, _, pageCrossed := env.FetchOperand(op)

	if pageCrossed {
		cycles++
	}

	env.IncrementProgramCounter(op.Size())

	return cycles, nil
}
