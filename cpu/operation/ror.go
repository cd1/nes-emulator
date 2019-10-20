package operation

import (
	"encoding/binary"
	"io"

	"github.com/cd1/nes-emulator/util"
)

const (
	OpMnemonicROR = "ROR"

	OpCodeRORZero        = 0x66
	OpCodeRORAccumulator = 0x6A
	OpCodeRORAbsoluteX   = 0x6E
	OpCodeRORZeroX       = 0x76
	OpCodeRORAbsolute    = 0x7E
)

func IsOpCodeValidROR(opCode uint8) bool {
	return opCode == OpCodeRORZero ||
		opCode == OpCodeRORAccumulator ||
		opCode == OpCodeRORAbsoluteX ||
		opCode == OpCodeRORZeroX ||
		opCode == OpCodeRORAbsolute
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
	case OpCodeRORAbsoluteX:
		var addr uint16

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewRORAbsoluteX(addr), nil
	case OpCodeRORZeroX:
		var addr uint8

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewRORZeroX(addr), nil
	case OpCodeRORAbsolute:
		var addr uint16

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewRORAbsolute(addr), nil
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
	case AddrModeAbsoluteX:
		return NewRORAbsoluteX(util.JoinBytesInWord([]uint8{arg0, arg1})), nil
	case AddrModeZeroX:
		return NewRORZeroX(arg0), nil
	case AddrModeAbsolute:
		return NewRORAbsolute(util.JoinBytesInWord([]uint8{arg0, arg1})), nil
	default:
		return nil, InvalidAddressModeError{
			AddressMode: addrMode,
		}
	}
}
