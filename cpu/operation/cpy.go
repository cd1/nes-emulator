package operation

import (
	"encoding/binary"
	"io"

	"github.com/cd1/nes-emulator/util"
)

const (
	OpMnemonicCPY = "CPY"

	OpCodeCPYImmediate = 0xC0
	OpCodeCPYZero      = 0xC4
	OpCodeCPYAbsolute  = 0xCC
)

func IsOpCodeValidCPY(opCode uint8) bool {
	return opCode == OpCodeCPYImmediate ||
		opCode == OpCodeCPYZero ||
		opCode == OpCodeCPYAbsolute
}

func IsMnemonicValidCPY(mnemonic string) bool {
	return mnemonic == OpMnemonicCPY
}

type CPY struct {
	baseOperation
}

// 0xC0: CPY #$NN
func NewCPYImmediate(value uint8) *CPY {
	return &CPY{
		baseOperation{
			code:        OpCodeCPYImmediate,
			addressMode: AddrModeImmediate,
			mnemonic:    OpMnemonicCPY,
			args:        []uint8{value},
		},
	}
}

// 0xC4: CPY $NN
func NewCPYZero(zeroAddress uint8) *CPY {
	return &CPY{
		baseOperation{
			code:        OpCodeCPYZero,
			addressMode: AddrModeZero,
			mnemonic:    OpMnemonicCPY,
			args:        []uint8{zeroAddress},
		},
	}
}

// 0xCC: CPY $NNNN
func NewCPYAbsolute(absoluteAddress uint16) *CPY {
	return &CPY{
		baseOperation{
			code:        OpCodeCPYAbsolute,
			addressMode: AddrModeAbsolute,
			mnemonic:    OpMnemonicCPY,
			args:        util.BreakWordIntoBytes(absoluteAddress),
		},
	}
}

func NewCPYBinary(opCode uint8, data io.Reader) (*CPY, error) {
	switch opCode {
	case OpCodeCPYImmediate:
		var addr uint8

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewCPYImmediate(addr), nil
	case OpCodeCPYZero:
		var addr uint8

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewCPYZero(addr), nil
	case OpCodeCPYAbsolute:
		var addr uint16

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewCPYAbsolute(addr), nil
	default:
		return nil, InvalidOpCodeError{
			OpCode: opCode,
		}
	}
}

func NewCPYFromBytes(addrMode uint8, arg0 uint8, arg1 uint8) (*CPY, error) {
	switch addrMode {
	case AddrModeImmediate:
		return NewCPYImmediate(arg0), nil
	case AddrModeZero, AddrModeRelative:
		return NewCPYZero(arg0), nil
	case AddrModeAbsolute:
		return NewCPYAbsolute(util.JoinBytesInWord([]uint8{arg0, arg1})), nil
	default:
		return nil, InvalidAddressModeError{
			AddressMode: addrMode,
		}
	}
}
