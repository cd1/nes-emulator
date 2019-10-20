package operation

import (
	"encoding/binary"
	"io"
)

const (
	OpMnemonicBMI = "BMI"

	OpCodeBMI = 0x30
)

func IsOpCodeValidBMI(opCode uint8) bool {
	return opCode == OpCodeBMI
}

func IsMnemonicValidBMI(mnemonic string) bool {
	return mnemonic == OpMnemonicBMI
}

type BMI struct {
	baseOperation
}

// 0x30: BMI $NN
func NewBMI(relativeAddress uint8) *BMI {
	return &BMI{
		baseOperation{
			code:        OpCodeBMI,
			addressMode: AddrModeRelative,
			mnemonic:    OpMnemonicBMI,
			args:        []uint8{relativeAddress},
		},
	}
}

func NewBMIBinary(opCode uint8, data io.Reader) (*BMI, error) {
	switch opCode {
	case OpCodeBMI:
		var addr uint8

		if err := binary.Read(data, binary.LittleEndian, &addr); err != nil {
			return nil, err
		}

		return NewBMI(addr), nil
	default:
		return nil, InvalidOpCodeError{
			OpCode: opCode,
		}
	}
}

func NewBMIFromBytes(addrMode uint8, arg0 uint8, arg1 uint8) (*BMI, error) {
	switch addrMode {
	case AddrModeZero, AddrModeRelative:
		return NewBMI(arg0), nil
	default:
		return nil, InvalidAddressModeError{
			AddressMode: addrMode,
		}
	}
}
