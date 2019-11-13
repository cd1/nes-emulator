package cpu

import (
	"encoding/binary"
	"io"
	"log"

	"github.com/cd1/nes-emulator/util"
)

const (
	OpMnemonicROR = "ROR"

	OpCodeRORZero        = 0x66
	OpCodeRORAccumulator = 0x6A
	OpCodeRORAbsolute    = 0x6E
	OpCodeRORZeroX       = 0x76
	OpCodeRORAbsoluteX   = 0x7E
)

func IsOpCodeValidROR(opCode uint8) bool {
	return opCode == OpCodeRORZero ||
		opCode == OpCodeRORAccumulator ||
		opCode == OpCodeRORAbsolute ||
		opCode == OpCodeRORZeroX ||
		opCode == OpCodeRORAbsoluteX
}

func IsMnemonicValidROR(mnemonic string) bool {
	return mnemonic == OpMnemonicROR
}

type ROR struct {
	baseOperation
}

// 0x66: ROR $NN
func NewRORZero(zeroAddress uint8) *ROR {
	return &ROR{
		baseOperation{
			code:        OpCodeRORZero,
			addressMode: AddrModeZero,
			mnemonic:    OpMnemonicROR,
			args:        []uint8{zeroAddress},
		},
	}
}

// 0x6A: ROR A
func NewRORAccumulator() *ROR {
	return &ROR{
		baseOperation{
			code:        OpCodeRORAccumulator,
			addressMode: AddrModeAccumulator,
			mnemonic:    OpMnemonicROR,
		},
	}
}

// 0x7E: ROR $NNNN
func NewRORAbsolute(absoluteAddress uint16) *ROR {
	return &ROR{
		baseOperation{
			code:        OpCodeRORAbsolute,
			addressMode: AddrModeAbsolute,
			mnemonic:    OpMnemonicROR,
			args:        util.BreakWordIntoBytes(absoluteAddress),
		},
	}
}

// 0x76: ROR $NN, X
func NewRORZeroX(zeroAddress uint8) *ROR {
	return &ROR{
		baseOperation{
			code:        OpCodeRORZeroX,
			addressMode: AddrModeZeroX,
			mnemonic:    OpMnemonicROR,
			args:        []uint8{zeroAddress},
		},
	}
}

// 0x6E: ROR $NNNN, X
func NewRORAbsoluteX(absoluteAddress uint16) *ROR {
	return &ROR{
		baseOperation{
			code:        OpCodeRORAbsoluteX,
			addressMode: AddrModeAbsoluteX,
			mnemonic:    OpMnemonicROR,
			args:        util.BreakWordIntoBytes(absoluteAddress),
		},
	}
}

func NewRORBinary(opCode uint8, data io.Reader) (*ROR, error) {
	switch opCode {
	case OpCodeRORZero:
		var addr uint8

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewRORZero(addr), nil
	case OpCodeRORAccumulator:
		return NewRORAccumulator(), nil
	case OpCodeRORAbsolute:
		var addr uint16

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewRORAbsolute(addr), nil
	case OpCodeRORZeroX:
		var addr uint8

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewRORZeroX(addr), nil
	case OpCodeRORAbsoluteX:
		var addr uint16

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewRORAbsoluteX(addr), nil
	default:
		return nil, InvalidOpCodeError{
			OpCode: opCode,
		}
	}
}

func NewRORFromBytes(addrMode uint8, arg0 uint8, arg1 uint8) (*ROR, error) {
	switch addrMode {
	case AddrModeZero, AddrModeRelative:
		return NewRORZero(arg0), nil
	case AddrModeAccumulator:
		return NewRORAccumulator(), nil
	case AddrModeAbsolute:
		return NewRORAbsolute(util.JoinBytesInWord([]uint8{arg0, arg1})), nil
	case AddrModeZeroX:
		return NewRORZeroX(arg0), nil
	case AddrModeAbsoluteX:
		return NewRORAbsoluteX(util.JoinBytesInWord([]uint8{arg0, arg1})), nil
	default:
		return nil, InvalidAddressModeError{
			AddressMode: addrMode,
		}
	}
}

func (op ROR) Cycles() uint8 {
	switch op.AddressMode() {
	case AddrModeZero:
		return 5
	case AddrModeAccumulator:
		return 2
	case AddrModeAbsolute:
		return 6
	case AddrModeZeroX:
		return 6
	case AddrModeAbsoluteX:
		return 7
	default:
		log.Printf("cannot calculate number of cycles for %#v: invalid address mode (%v)", op, op.AddressMode())
		return 0
	}
}

func (op ROR) ExecuteIn(env OperationEnvironment) (uint8, error) {
	cycles := op.Cycles()
	address, operand, _ := env.FetchOperand(op)

	result := operand >> 1
	if env.IsStatusCarry() {
		result |= 0x80
	}

	if op.AddressMode() == AddrModeAccumulator {
		env.SetAccumulator(result)
	} else {
		env.WriteByte(address, result)
	}

	env.SetStatusCarry(operand&0x01 != 0x00)
	env.SetStatusZero(result == 0x00)
	env.SetStatusNegative(result&0x80 != 0x00)

	env.IncrementProgramCounter(op.Size())

	return cycles, nil
}
