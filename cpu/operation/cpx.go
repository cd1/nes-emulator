package operation

import (
	"encoding/binary"
	"io"

	"github.com/cd1/nes-emulator/util"
)

const (
	OpMnemonicCPX = "CPX"

	OpCodeCPXImmediate = 0xE0
	OpCodeCPXZero      = 0xE4
	OpCodeCPXAbsolute  = 0xEC
)

func IsOpCodeValidCPX(opCode uint8) bool {
	return opCode == OpCodeCPXImmediate ||
		opCode == OpCodeCPXZero ||
		opCode == OpCodeCPXAbsolute
}

func IsMnemonicValidCPX(mnemonic string) bool {
	return mnemonic == OpMnemonicCPX
}

type CPX struct {
	baseOperation
}

// 0xE0: CPX #$NN
func NewCPXImmediate(value uint8) *CPX {
	return &CPX{
		baseOperation{
			code:        OpCodeCPXImmediate,
			addressMode: AddrModeImmediate,
			mnemonic:    OpMnemonicCPX,
			args:        []uint8{value},
		},
	}
}

// 0xE4: CPX $NN
func NewCPXZero(zeroAddress uint8) *CPX {
	return &CPX{
		baseOperation{
			code:        OpCodeCPXZero,
			addressMode: AddrModeZero,
			mnemonic:    OpMnemonicCPX,
			args:        []uint8{zeroAddress},
		},
	}
}

// 0xEC: CPX $NNNN
func NewCPXAbsolute(absoluteAddress uint16) *CPX {
	return &CPX{
		baseOperation{
			code:        OpCodeCPXAbsolute,
			addressMode: AddrModeAbsolute,
			mnemonic:    OpMnemonicCPX,
			args:        util.BreakWordIntoBytes(absoluteAddress),
		},
	}
}

func NewCPXBinary(opCode uint8, data io.Reader) (*CPX, error) {
	switch opCode {
	case OpCodeCPXImmediate:
		var addr uint8

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewCPXImmediate(addr), nil
	case OpCodeCPXZero:
		var addr uint8

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewCPXZero(addr), nil
	case OpCodeCPXAbsolute:
		var addr uint16

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewCPXAbsolute(addr), nil
	default:
		return nil, InvalidOpCodeError{
			OpCode: opCode,
		}
	}
}

func NewCPXFromBytes(addrMode uint8, arg0 uint8, arg1 uint8) (*CPX, error) {
	switch addrMode {
	case AddrModeImmediate:
		return NewCPXImmediate(arg0), nil
	case AddrModeZero, AddrModeRelative:
		return NewCPXZero(arg0), nil
	case AddrModeAbsolute:
		return NewCPXAbsolute(util.JoinBytesInWord([]uint8{arg0, arg1})), nil
	default:
		return nil, InvalidAddressModeError{
			AddressMode: addrMode,
		}
	}
}
