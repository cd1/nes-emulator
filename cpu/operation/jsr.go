package operation

import (
	"encoding/binary"
	"io"

	"github.com/cd1/nes-emulator/util"
)

const (
	OpMnemonicJSR = "JSR"

	OpCodeJSR = 0x20
)

func IsOpCodeValidJSR(opCode uint8) bool {
	return opCode == OpCodeJSR
}

func IsMnemonicValidJSR(mnemonic string) bool {
	return mnemonic == OpMnemonicJSR
}

type JSR struct {
	baseOperation
}

// 0x20: JSR $NNNN
func NewJSR(absoluteAddress uint16) *JSR {
	return &JSR{
		baseOperation{
			code:        OpCodeJSR,
			addressMode: AddrModeAbsolute,
			mnemonic:    OpMnemonicJSR,
			args:        util.BreakWordIntoBytes(absoluteAddress),
		},
	}
}

func NewJSRBinary(opCode uint8, data io.Reader) (*JSR, error) {
	switch opCode {
	case OpCodeJSR:
		var addr uint16

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewJSR(addr), nil
	default:
		return nil, InvalidOpCodeError{
			OpCode: opCode,
		}
	}
}

func NewJSRFromBytes(addrMode uint8, arg0 uint8, arg1 uint8) (*JSR, error) {
	switch addrMode {
	case AddrModeAbsolute:
		return NewJSR(util.JoinBytesInWord([]uint8{arg0, arg1})), nil
	default:
		return nil, InvalidAddressModeError{
			AddressMode: addrMode,
		}
	}
}
