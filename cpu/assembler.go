package cpu

import (
	"bufio"
	"encoding/binary"
	"io"
	"strconv"
	"strings"

	"github.com/cd1/nes-emulator/cpu/operation"
	"github.com/cd1/nes-emulator/util"
)

func Assemble(r io.Reader, w io.Writer) error {
	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		opLine := scanner.Text()

		op, err := convertTextToOperation(opLine)
		if err != nil {
			return err
		}

		if err := binary.Write(w, binary.LittleEndian, op.Code()); err != nil {
			return err
		}

		switch op.Size() {
		case 2:
			if err := binary.Write(w, binary.LittleEndian, op.ByteArg()); err != nil {
				return err
			}
		case 3:
			if err := binary.Write(w, binary.LittleEndian, op.WordArg()); err != nil {
				return err
			}
		}
	}
	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

func convertTextToOperation(line string) (operation.Operation, error) {
	mnemonic, addrMode, arg0, arg1, err := breakdownText(line)
	if err != nil {
		return nil, err
	}

	var op operation.Operation

	switch {
	case operation.IsMnemonicValidADC(mnemonic):
		op, err = operation.NewADCFromBytes(addrMode, arg0, arg1)
	case operation.IsMnemonicValidAND(mnemonic):
		op, err = operation.NewANDFromBytes(addrMode, arg0, arg1)
	case operation.IsMnemonicValidASL(mnemonic):
		op, err = operation.NewASLFromBytes(addrMode, arg0, arg1)
	case operation.IsMnemonicValidBCC(mnemonic):
		op, err = operation.NewBCCFromBytes(addrMode, arg0, arg1)
	case operation.IsMnemonicValidBCS(mnemonic):
		op, err = operation.NewBCSFromBytes(addrMode, arg0, arg1)
	case operation.IsMnemonicValidBEQ(mnemonic):
		op, err = operation.NewBEQFromBytes(addrMode, arg0, arg1)
	case operation.IsMnemonicValidBIT(mnemonic):
		op, err = operation.NewBITFromBytes(addrMode, arg0, arg1)
	case operation.IsMnemonicValidBMI(mnemonic):
		op, err = operation.NewBMIFromBytes(addrMode, arg0, arg1)
	case operation.IsMnemonicValidBNE(mnemonic):
		op, err = operation.NewBNEFromBytes(addrMode, arg0, arg1)
	case operation.IsMnemonicValidBPL(mnemonic):
		op, err = operation.NewBPLFromBytes(addrMode, arg0, arg1)
	case operation.IsMnemonicValidBRK(mnemonic):
		op, err = operation.NewBRKFromBytes(addrMode, arg0, arg1)
	case operation.IsMnemonicValidBVC(mnemonic):
		op, err = operation.NewBVCFromBytes(addrMode, arg0, arg1)
	case operation.IsMnemonicValidBVS(mnemonic):
		op, err = operation.NewBVSFromBytes(addrMode, arg0, arg1)
	case operation.IsMnemonicValidCLC(mnemonic):
		op, err = operation.NewCLCFromBytes(addrMode, arg0, arg1)
	case operation.IsMnemonicValidCLD(mnemonic):
		op, err = operation.NewCLDFromBytes(addrMode, arg0, arg1)
	case operation.IsMnemonicValidCLI(mnemonic):
		op, err = operation.NewCLIFromBytes(addrMode, arg0, arg1)
	case operation.IsMnemonicValidCLV(mnemonic):
		op, err = operation.NewCLVFromBytes(addrMode, arg0, arg1)
	case operation.IsMnemonicValidCMP(mnemonic):
		op, err = operation.NewCMPFromBytes(addrMode, arg0, arg1)
	case operation.IsMnemonicValidCPX(mnemonic):
		op, err = operation.NewCPXFromBytes(addrMode, arg0, arg1)
	case operation.IsMnemonicValidCPY(mnemonic):
		op, err = operation.NewCPYFromBytes(addrMode, arg0, arg1)
	case operation.IsMnemonicValidDEC(mnemonic):
		op, err = operation.NewDECFromBytes(addrMode, arg0, arg1)
	case operation.IsMnemonicValidDEX(mnemonic):
		op, err = operation.NewDEXFromBytes(addrMode, arg0, arg1)
	case operation.IsMnemonicValidDEY(mnemonic):
		op, err = operation.NewDEYFromBytes(addrMode, arg0, arg1)
	case operation.IsMnemonicValidEOR(mnemonic):
		op, err = operation.NewEORFromBytes(addrMode, arg0, arg1)
	case operation.IsMnemonicValidINC(mnemonic):
		op, err = operation.NewINCFromBytes(addrMode, arg0, arg1)
	case operation.IsMnemonicValidINX(mnemonic):
		op, err = operation.NewINXFromBytes(addrMode, arg0, arg1)
	case operation.IsMnemonicValidINY(mnemonic):
		op, err = operation.NewINYFromBytes(addrMode, arg0, arg1)
	case operation.IsMnemonicValidJMP(mnemonic):
		op, err = operation.NewJMPFromBytes(addrMode, arg0, arg1)
	case operation.IsMnemonicValidJSR(mnemonic):
		op, err = operation.NewJSRFromBytes(addrMode, arg0, arg1)
	case operation.IsMnemonicValidLDA(mnemonic):
		op, err = operation.NewLDAFromBytes(addrMode, arg0, arg1)
	case operation.IsMnemonicValidLDX(mnemonic):
		op, err = operation.NewLDXFromBytes(addrMode, arg0, arg1)
	case operation.IsMnemonicValidLDY(mnemonic):
		op, err = operation.NewLDYFromBytes(addrMode, arg0, arg1)
	case operation.IsMnemonicValidLSR(mnemonic):
		op, err = operation.NewLSRFromBytes(addrMode, arg0, arg1)
	case operation.IsMnemonicValidNOP(mnemonic):
		op, err = operation.NewNOPFromBytes(addrMode, arg0, arg1)
	case operation.IsMnemonicValidORA(mnemonic):
		op, err = operation.NewORAFromBytes(addrMode, arg0, arg1)
	case operation.IsMnemonicValidPHA(mnemonic):
		op, err = operation.NewPHAFromBytes(addrMode, arg0, arg1)
	case operation.IsMnemonicValidPHP(mnemonic):
		op, err = operation.NewPHPFromBytes(addrMode, arg0, arg1)
	case operation.IsMnemonicValidPLA(mnemonic):
		op, err = operation.NewPLAFromBytes(addrMode, arg0, arg1)
	case operation.IsMnemonicValidPLP(mnemonic):
		op, err = operation.NewPLPFromBytes(addrMode, arg0, arg1)
	case operation.IsMnemonicValidROL(mnemonic):
		op, err = operation.NewROLFromBytes(addrMode, arg0, arg1)
	case operation.IsMnemonicValidROR(mnemonic):
		op, err = operation.NewRORFromBytes(addrMode, arg0, arg1)
	case operation.IsMnemonicValidRTI(mnemonic):
		op, err = operation.NewRTIFromBytes(addrMode, arg0, arg1)
	case operation.IsMnemonicValidRTS(mnemonic):
		op, err = operation.NewRTSFromBytes(addrMode, arg0, arg1)
	case operation.IsMnemonicValidSBC(mnemonic):
		op, err = operation.NewSBCFromBytes(addrMode, arg0, arg1)
	case operation.IsMnemonicValidSEC(mnemonic):
		op, err = operation.NewSECFromBytes(addrMode, arg0, arg1)
	case operation.IsMnemonicValidSED(mnemonic):
		op, err = operation.NewSEDFromBytes(addrMode, arg0, arg1)
	case operation.IsMnemonicValidSEI(mnemonic):
		op, err = operation.NewSEIFromBytes(addrMode, arg0, arg1)
	case operation.IsMnemonicValidSTA(mnemonic):
		op, err = operation.NewSTAFromBytes(addrMode, arg0, arg1)
	case operation.IsMnemonicValidSTX(mnemonic):
		op, err = operation.NewSTXFromBytes(addrMode, arg0, arg1)
	case operation.IsMnemonicValidSTY(mnemonic):
		op, err = operation.NewSTYFromBytes(addrMode, arg0, arg1)
	case operation.IsMnemonicValidTAX(mnemonic):
		op, err = operation.NewTAYFromBytes(addrMode, arg0, arg1)
	case operation.IsMnemonicValidTAY(mnemonic):
		op, err = operation.NewTAYFromBytes(addrMode, arg0, arg1)
	case operation.IsMnemonicValidTSX(mnemonic):
		op, err = operation.NewTSXFromBytes(addrMode, arg0, arg1)
	case operation.IsMnemonicValidTXA(mnemonic):
		op, err = operation.NewTXAFromBytes(addrMode, arg0, arg1)
	case operation.IsMnemonicValidTXS(mnemonic):
		op, err = operation.NewTXSFromBytes(addrMode, arg0, arg1)
	case operation.IsMnemonicValidTYA(mnemonic):
		op, err = operation.NewTYAFromBytes(addrMode, arg0, arg1)
	default:
		return nil, operation.InvalidMnemonicError{
			Mnemonic: mnemonic,
		}
	}

	if err != nil {
		return nil, err
	}

	return op, nil
}

func breakdownText(line string) (string, uint8, uint8, uint8, error) {
	lineBySpace := strings.Split(line, " ")
	mnemonic := lineBySpace[0]

	if len(lineBySpace) == 1 {
		// AddrModeImplied
		// e.g. XXX
		return mnemonic, operation.AddrModeImplied, 0, 0, nil
	}

	arg0Field := lineBySpace[1]

	if arg0Field == "A" {
		// AddrModeAccumulator
		// e.g. XXX A
		return mnemonic, operation.AddrModeAccumulator, 0, 0, nil
	}

	switch arg0Field[0] {
	case '#':
		// AddrModeImmediate
		// e.g.: XXX %$NN
		arg0, err := strconv.ParseUint(arg0Field[2:4], 16, 8)
		if err != nil {
			return "", 0, 0, 0, err
		}

		return mnemonic, operation.AddrModeImmediate, uint8(arg0), 0, nil
	case '$':
		if arg0Field[len(arg0Field)-1] == ',' {
			arg1 := lineBySpace[2]

			switch arg1 {
			case "X":
				arg0Str := arg0Field[1 : len(arg0Field)-1]
				switch len(arg0Str) {
				case 2:
					// AddrModeZeroX
					// e.g.: XXX $NN, X
					arg0, err := strconv.ParseUint(arg0Str, 16, 8)
					if err != nil {
						return "", 0, 0, 0, err
					}
					return mnemonic, operation.AddrModeZeroX, uint8(arg0), 0, nil
				case 4:
					// AddrModeAbsoluteX
					// e.g.: XXX $NNNN, X
					argWord, err := strconv.ParseUint(arg0Str, 16, 16)
					if err != nil {
						return "", 0, 0, 0, err
					}
					wordParts := util.BreakWordIntoBytes(uint16(argWord))
					return mnemonic, operation.AddrModeAbsoluteX, wordParts[0], wordParts[1], nil
				}
			case "Y":
				arg0Str := arg0Field[1 : len(arg0Field)-1]
				switch len(arg0Str) {
				case 2:
					// AddrModeZeroY
					// e.g.: XXX $NN, Y
					arg0, err := strconv.ParseUint(arg0Str, 16, 8)
					if err != nil {
						return "", 0, 0, 0, err
					}
					return mnemonic, operation.AddrModeZeroY, uint8(arg0), 0, nil
				case 4:
					// AddrModeAbsoluteY
					// e.g.: XXX $NNNN, Y
					argWord, err := strconv.ParseUint(arg0Str, 16, 16)
					if err != nil {
						return "", 0, 0, 0, err
					}
					wordParts := util.BreakWordIntoBytes(uint16(argWord))
					return mnemonic, operation.AddrModeAbsoluteY, wordParts[0], wordParts[1], nil
				}
			}
		} else {
			arg0Str := arg0Field[1:]
			switch len(arg0Str) {
			case 2:
				// AddrModeRelative | AddrModeZero
				// e.g.: XXX $NN
				arg0, err := strconv.ParseUint(arg0Str, 16, 8)
				if err != nil {
					return "", 0, 0, 0, err
				}
				return mnemonic, operation.AddrModeRelative, uint8(arg0), 0, nil
			case 4:
				// AddrModeAbsolute
				// e.g.: XXX $NNNN
				argWord, err := strconv.ParseUint(arg0Str, 16, 16)
				if err != nil {
					return "", 0, 0, 0, err
				}
				wordParts := util.BreakWordIntoBytes(uint16(argWord))
				return mnemonic, operation.AddrModeAbsolute, wordParts[0], wordParts[1], nil
			}
		}
	case '(':
		if arg0Field[len(arg0Field)-1] == ')' {
			// AddrModeIndirect
			// e.g.: XXX ($NNNN)
			arg0Str := arg0Field[2:6]
			argWord, err := strconv.ParseUint(arg0Str, 16, 16)
			if err != nil {
				return "", 0, 0, 0, err
			}
			wordParts := util.BreakWordIntoBytes(uint16(argWord))
			return mnemonic, operation.AddrModeIndirect, wordParts[0], wordParts[1], nil
		} else {
			arg1Field := lineBySpace[2]

			if arg1Field[len(arg1Field)-1] == ')' {
				// AddrModeIndirectX
				// e.g.: XXX ($NN, X)
				arg0Str := arg0Field[2:4]
				arg0, err := strconv.ParseUint(arg0Str, 16, 8)
				if err != nil {
					return "", 0, 0, 0, err
				}
				return mnemonic, operation.AddrModeIndirectX, uint8(arg0), 0, nil
			} else {
				// AddrModeIndirectY
				// e.g.: XXX ($NN), Y
				arg0Str := arg0Field[2:4]
				arg0, err := strconv.ParseUint(arg0Str, 16, 8)
				if err != nil {
					return "", 0, 0, 0, err
				}
				return mnemonic, operation.AddrModeIndirectY, uint8(arg0), 0, nil
			}
		}
	}

	return "", 0, 0, 0, operation.InvalidSyntaxError{
		Line: line,
	}
}
