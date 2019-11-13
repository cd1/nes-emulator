package cpu

import (
	"encoding/binary"
	"io"
	"log"

	"github.com/cd1/nes-emulator/util"
)

const (
	OpMnemonicSTA = "STA"

	OpCodeSTAIndirectX = 0x81
	OpCodeSTAZero      = 0x85
	OpCodeSTAAbsolute  = 0x8D
	OpCodeSTAIndirectY = 0x91
	OpCodeSTAZeroX     = 0x95
	OpCodeSTAAbsoluteY = 0x99
	OpCodeSTAAbsoluteX = 0x9D
)

func IsOpCodeValidSTA(opCode uint8) bool {
	return opCode == OpCodeSTAIndirectX ||
		opCode == OpCodeSTAZero ||
		opCode == OpCodeSTAAbsolute ||
		opCode == OpCodeSTAIndirectY ||
		opCode == OpCodeSTAZeroX ||
		opCode == OpCodeSTAAbsoluteY ||
		opCode == OpCodeSTAAbsoluteX
}

func IsMnemonicValidSTA(mnemonic string) bool {
	return mnemonic == OpMnemonicSTA
}

type STA struct {
	baseOperation
}

// 0x81: STA ($NN, X)
func NewSTAIndirectX(indirectAddress uint8) *STA {
	return &STA{
		baseOperation{
			code:        OpCodeSTAIndirectX,
			addressMode: AddrModeIndirectX,
			mnemonic:    OpMnemonicSTA,
			args:        []uint8{indirectAddress},
		},
	}
}

// 0x85: STA $NN
func NewSTAZero(zeroAddress uint8) *STA {
	return &STA{
		baseOperation{
			code:        OpCodeSTAZero,
			addressMode: AddrModeZero,
			mnemonic:    OpMnemonicSTA,
			args:        []uint8{zeroAddress},
		},
	}
}

// 0x8D: STA $NNNN
func NewSTAAbsolute(absoluteAddress uint16) *STA {
	return &STA{
		baseOperation{
			code:        OpCodeSTAAbsolute,
			addressMode: AddrModeAbsolute,
			mnemonic:    OpMnemonicSTA,
			args:        util.BreakWordIntoBytes(absoluteAddress),
		},
	}
}

// 0x91: STA ($NN), Y
func NewSTAIndirectY(indirectAddress uint8) *STA {
	return &STA{
		baseOperation{
			code:        OpCodeSTAIndirectY,
			addressMode: AddrModeIndirectY,
			mnemonic:    OpMnemonicSTA,
			args:        []uint8{indirectAddress},
		},
	}
}

// 0x95: STA $NN, X
func NewSTAZeroX(zeroAddress uint8) *STA {
	return &STA{
		baseOperation{
			code:        OpCodeSTAZeroX,
			addressMode: AddrModeZeroX,
			mnemonic:    OpMnemonicSTA,
			args:        []uint8{zeroAddress},
		},
	}
}

// 0x99: STA $NNNN, Y
func NewSTAAbsoluteY(absoluteAddress uint16) *STA {
	return &STA{
		baseOperation{
			code:        OpCodeSTAAbsoluteY,
			addressMode: AddrModeAbsoluteY,
			mnemonic:    OpMnemonicSTA,
			args:        util.BreakWordIntoBytes(absoluteAddress),
		},
	}
}

// 0x9D: STA $NNNN, X
func NewSTAAbsoluteX(absoluteAddress uint16) *STA {
	return &STA{
		baseOperation{
			code:        OpCodeSTAAbsoluteX,
			addressMode: AddrModeAbsoluteX,
			mnemonic:    OpMnemonicSTA,
			args:        util.BreakWordIntoBytes(absoluteAddress),
		},
	}
}

func NewSTABinary(opCode uint8, data io.Reader) (*STA, error) {
	switch opCode {
	case OpCodeSTAIndirectX:
		var addr uint8

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewSTAIndirectX(addr), nil
	case OpCodeSTAZero:
		var addr uint8

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewSTAZero(addr), nil
	case OpCodeSTAAbsolute:
		var addr uint16

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewSTAAbsolute(addr), nil
	case OpCodeSTAIndirectY:
		var addr uint8

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewSTAIndirectY(addr), nil
	case OpCodeSTAZeroX:
		var addr uint8

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewSTAZeroX(addr), nil
	case OpCodeSTAAbsoluteY:
		var addr uint16

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewSTAAbsoluteY(addr), nil
	case OpCodeSTAAbsoluteX:
		var addr uint16

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewSTAAbsoluteX(addr), nil
	default:
		return nil, InvalidOpCodeError{
			OpCode: opCode,
		}
	}
}

func NewSTAFromBytes(addrMode uint8, arg0 uint8, arg1 uint8) (*STA, error) {
	switch addrMode {
	case AddrModeIndirectX:
		return NewSTAIndirectX(arg0), nil
	case AddrModeZero, AddrModeRelative:
		return NewSTAZero(arg0), nil
	case AddrModeAbsolute:
		return NewSTAAbsolute(util.JoinBytesInWord([]uint8{arg0, arg1})), nil
	case AddrModeIndirectY:
		return NewSTAIndirectY(arg0), nil
	case AddrModeZeroX:
		return NewSTAZeroX(arg0), nil
	case AddrModeAbsoluteY:
		return NewSTAAbsoluteY(util.JoinBytesInWord([]uint8{arg0, arg1})), nil
	case AddrModeAbsoluteX:
		return NewSTAAbsoluteX(util.JoinBytesInWord([]uint8{arg0, arg1})), nil
	default:
		return nil, InvalidAddressModeError{
			AddressMode: addrMode,
		}
	}
}

func (op STA) Cycles() uint8 {
	switch op.AddressMode() {
	case AddrModeIndirectX:
		return 6
	case AddrModeZero:
		return 3
	case AddrModeAbsolute:
		return 4
	case AddrModeIndirectY:
		return 6
	case AddrModeZeroX:
		return 4
	case AddrModeAbsoluteY:
		return 5
	case AddrModeAbsoluteX:
		return 5
	default:
		log.Printf("cannot calculate number of cycles for %#v: invalid address mode (%v)", op, op.AddressMode())
		return 0
	}
}
func (op STA) ExecuteIn(env OperationEnvironment) (uint8, error) {
	cycles := op.Cycles()
	address, _, _ := env.FetchOperand(op)

	env.WriteByte(address, env.GetAccumulator())

	env.IncrementProgramCounter(op.Size())

	return cycles, nil
}
