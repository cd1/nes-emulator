package operation

import (
	"encoding/binary"
	"io"

	"github.com/cd1/nes-emulator/util"
)

// ADC  Add Memory to Accumulator with Carry
//
// A + M + C -> A, C                N Z C I D V
//                                 + + + - - +
//
// addressing    assembler    opc  bytes  cyles
// --------------------------------------------
// immidiate     ADC #oper     69    2     2
// zeropage      ADC oper      65    2     3
// zeropage,X    ADC oper,X    75    2     4
// absolute      ADC oper      6D    3     4
// absolute,X    ADC oper,X    7D    3     4*
// absolute,Y    ADC oper,Y    79    3     4*
// (indirect,X)  ADC (oper,X)  61    2     6
// (indirect),Y  ADC (oper),Y  71    2     5*

const (
	// Mnemonic for ADC operations
	OpMnemonicADC = "ADC"

	// OpCodes for ADC operations
	OpCodeADCIndirectX = 0x61
	OpCodeADCZero      = 0x65
	OpCodeADCImmediate = 0x69
	OpCodeADCAbsolute  = 0x6D
	OpCodeADCIndirectY = 0x71
	OpCodeADCZeroX     = 0x75
	OpCodeADCAbsoluteY = 0x79
	OpCodeADCAbsoluteX = 0x7D
)

// IsOpCodeValidADC checks if a specific opCode is a valid code for
// an ADC operation.
func IsOpCodeValidADC(opCode uint8) bool {
	return opCode == OpCodeADCIndirectX ||
		opCode == OpCodeADCZero ||
		opCode == OpCodeADCImmediate ||
		opCode == OpCodeADCAbsolute ||
		opCode == OpCodeADCIndirectY ||
		opCode == OpCodeADCZeroX ||
		opCode == OpCodeADCAbsoluteY ||
		opCode == OpCodeADCAbsoluteX
}

func IsMnemonicValidADC(mnemonic string) bool {
	return mnemonic == OpMnemonicADC
}

// ADC defines an ADC operation.
type ADC struct {
	baseOperation
}

// 0x61: ADC ($NN, X)
func NewADCIndirectX(indirectAddress uint8) *ADC {
	return &ADC{
		baseOperation{
			code:        OpCodeADCIndirectX,
			addressMode: AddrModeIndirectX,
			mnemonic:    OpMnemonicADC,
			args:        []uint8{indirectAddress},
		},
	}
}

// 0x65: ADC $NN
func NewADCZero(zeroAddress uint8) *ADC {
	return &ADC{
		baseOperation{
			code:        OpCodeADCZero,
			addressMode: AddrModeZero,
			mnemonic:    OpMnemonicADC,
			args:        []uint8{zeroAddress},
		},
	}
}

// 0x69: ADC #$NN
func NewADCImmediate(value uint8) *ADC {
	return &ADC{
		baseOperation{
			code:        OpCodeADCImmediate,
			addressMode: AddrModeImmediate,
			mnemonic:    OpMnemonicADC,
			args:        []uint8{value},
		},
	}
}

// 0x6D: ADC $NNNN
func NewADCAbsolute(absoluteAddress uint16) *ADC {
	return &ADC{
		baseOperation{
			code:        OpCodeADCAbsolute,
			addressMode: AddrModeAbsolute,
			mnemonic:    OpMnemonicADC,
			args:        util.BreakWordIntoBytes(absoluteAddress),
		},
	}
}

// 0x71: ADC ($NN), Y
func NewADCIndirectY(indirectAddress uint8) *ADC {
	return &ADC{
		baseOperation{
			code:        OpCodeADCIndirectY,
			addressMode: AddrModeIndirectY,
			mnemonic:    OpMnemonicADC,
			args:        []uint8{indirectAddress},
		},
	}
}

// 0x75: ADC $NN, X
func NewADCZeroX(zeroAddress uint8) *ADC {
	return &ADC{
		baseOperation{
			code:        OpCodeADCZeroX,
			addressMode: AddrModeZeroX,
			mnemonic:    OpMnemonicADC,
			args:        []uint8{zeroAddress},
		},
	}
}

// 0x79: ADC $NNNN, Y
func NewADCAbsoluteY(absoluteAddress uint16) *ADC {
	return &ADC{
		baseOperation{
			code:        OpCodeADCAbsoluteY,
			addressMode: AddrModeAbsoluteY,
			mnemonic:    OpMnemonicADC,
			args:        util.BreakWordIntoBytes(absoluteAddress),
		},
	}
}

// 0x7D: ADC $NNNN, X
func NewADCAbsoluteX(absoluteAddress uint16) *ADC {
	return &ADC{
		baseOperation{
			code:        OpCodeADCAbsoluteX,
			addressMode: AddrModeAbsoluteX,
			mnemonic:    OpMnemonicADC,
			args:        util.BreakWordIntoBytes(absoluteAddress),
		},
	}
}

// NewADCBinary creates a new ADC operation according to
// the given opCode. The rest of the parameters, if needed, will be read from
// data. If opCode isn't a valid ADC code or if there's an error reading from
// data, an error is returned.
func NewADCBinary(opCode uint8, data io.Reader) (*ADC, error) {
	switch opCode {
	case OpCodeADCIndirectX:
		var addr uint8

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewADCIndirectX(addr), nil
	case OpCodeADCZero:
		var addr uint8

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewADCZero(addr), nil
	case OpCodeADCImmediate:
		var addr uint8

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewADCImmediate(addr), nil
	case OpCodeADCAbsolute:
		var addr uint16

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewADCAbsolute(addr), nil
	case OpCodeADCIndirectY:
		var addr uint8

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewADCIndirectY(addr), nil
	case OpCodeADCZeroX:
		var addr uint8

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewADCZeroX(addr), nil
	case OpCodeADCAbsoluteY:
		var addr uint16

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewADCAbsoluteY(addr), nil
	case OpCodeADCAbsoluteX:
		var addr uint16

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewADCAbsoluteX(addr), nil
	default:
		return nil, InvalidOpCodeError{
			OpCode: opCode,
		}
	}
}

func NewADCFromBytes(addrMode uint8, arg0 uint8, arg1 uint8) (*ADC, error) {
	switch addrMode {
	case AddrModeIndirectX:
		return NewADCIndirectX(arg0), nil
	case AddrModeZero, AddrModeRelative:
		return NewADCZero(arg0), nil
	case AddrModeImmediate:
		return NewADCImmediate(arg0), nil
	case AddrModeAbsolute:
		return NewADCAbsolute(util.JoinBytesInWord([]uint8{arg0, arg1})), nil
	case AddrModeIndirectY:
		return NewADCIndirectY(arg0), nil
	case AddrModeZeroX:
		return NewADCZeroX(arg0), nil
	case AddrModeAbsoluteY:
		return NewADCAbsoluteY(util.JoinBytesInWord([]uint8{arg0, arg1})), nil
	case AddrModeAbsoluteX:
		return NewADCAbsoluteX(util.JoinBytesInWord([]uint8{arg0, arg1})), nil
	default:
		return nil, InvalidAddressModeError{
			AddressMode: addrMode,
		}
	}
}
