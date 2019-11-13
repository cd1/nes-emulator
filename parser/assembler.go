package parser

import (
	"bufio"
	"encoding/binary"
	"io"
	"strconv"
	"strings"

	"github.com/cd1/nes-emulator/cpu"
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

func convertTextToOperation(line string) (cpu.Operation, error) {
	mnemonic, addrMode, arg0, arg1, err := breakdownText(line)
	if err != nil {
		return nil, err
	}

	var op cpu.Operation

	switch {
	case cpu.IsMnemonicValidADC(mnemonic):
		op, err = cpu.NewADCFromBytes(addrMode, arg0, arg1)
	case cpu.IsMnemonicValidAND(mnemonic):
		op, err = cpu.NewANDFromBytes(addrMode, arg0, arg1)
	case cpu.IsMnemonicValidASL(mnemonic):
		op, err = cpu.NewASLFromBytes(addrMode, arg0, arg1)
	case cpu.IsMnemonicValidBCC(mnemonic):
		op, err = cpu.NewBCCFromBytes(addrMode, arg0, arg1)
	case cpu.IsMnemonicValidBCS(mnemonic):
		op, err = cpu.NewBCSFromBytes(addrMode, arg0, arg1)
	case cpu.IsMnemonicValidBEQ(mnemonic):
		op, err = cpu.NewBEQFromBytes(addrMode, arg0, arg1)
	case cpu.IsMnemonicValidBIT(mnemonic):
		op, err = cpu.NewBITFromBytes(addrMode, arg0, arg1)
	case cpu.IsMnemonicValidBMI(mnemonic):
		op, err = cpu.NewBMIFromBytes(addrMode, arg0, arg1)
	case cpu.IsMnemonicValidBNE(mnemonic):
		op, err = cpu.NewBNEFromBytes(addrMode, arg0, arg1)
	case cpu.IsMnemonicValidBPL(mnemonic):
		op, err = cpu.NewBPLFromBytes(addrMode, arg0, arg1)
	case cpu.IsMnemonicValidBRK(mnemonic):
		op, err = cpu.NewBRKFromBytes(addrMode, arg0, arg1)
	case cpu.IsMnemonicValidBVC(mnemonic):
		op, err = cpu.NewBVCFromBytes(addrMode, arg0, arg1)
	case cpu.IsMnemonicValidBVS(mnemonic):
		op, err = cpu.NewBVSFromBytes(addrMode, arg0, arg1)
	case cpu.IsMnemonicValidCLC(mnemonic):
		op, err = cpu.NewCLCFromBytes(addrMode, arg0, arg1)
	case cpu.IsMnemonicValidCLD(mnemonic):
		op, err = cpu.NewCLDFromBytes(addrMode, arg0, arg1)
	case cpu.IsMnemonicValidCLI(mnemonic):
		op, err = cpu.NewCLIFromBytes(addrMode, arg0, arg1)
	case cpu.IsMnemonicValidCLV(mnemonic):
		op, err = cpu.NewCLVFromBytes(addrMode, arg0, arg1)
	case cpu.IsMnemonicValidCMP(mnemonic):
		op, err = cpu.NewCMPFromBytes(addrMode, arg0, arg1)
	case cpu.IsMnemonicValidCPX(mnemonic):
		op, err = cpu.NewCPXFromBytes(addrMode, arg0, arg1)
	case cpu.IsMnemonicValidCPY(mnemonic):
		op, err = cpu.NewCPYFromBytes(addrMode, arg0, arg1)
	case cpu.IsMnemonicValidDEC(mnemonic):
		op, err = cpu.NewDECFromBytes(addrMode, arg0, arg1)
	case cpu.IsMnemonicValidDEX(mnemonic):
		op, err = cpu.NewDEXFromBytes(addrMode, arg0, arg1)
	case cpu.IsMnemonicValidDEY(mnemonic):
		op, err = cpu.NewDEYFromBytes(addrMode, arg0, arg1)
	case cpu.IsMnemonicValidEOR(mnemonic):
		op, err = cpu.NewEORFromBytes(addrMode, arg0, arg1)
	case cpu.IsMnemonicValidINC(mnemonic):
		op, err = cpu.NewINCFromBytes(addrMode, arg0, arg1)
	case cpu.IsMnemonicValidINX(mnemonic):
		op, err = cpu.NewINXFromBytes(addrMode, arg0, arg1)
	case cpu.IsMnemonicValidINY(mnemonic):
		op, err = cpu.NewINYFromBytes(addrMode, arg0, arg1)
	case cpu.IsMnemonicValidJMP(mnemonic):
		op, err = cpu.NewJMPFromBytes(addrMode, arg0, arg1)
	case cpu.IsMnemonicValidJSR(mnemonic):
		op, err = cpu.NewJSRFromBytes(addrMode, arg0, arg1)
	case cpu.IsMnemonicValidLDA(mnemonic):
		op, err = cpu.NewLDAFromBytes(addrMode, arg0, arg1)
	case cpu.IsMnemonicValidLDX(mnemonic):
		op, err = cpu.NewLDXFromBytes(addrMode, arg0, arg1)
	case cpu.IsMnemonicValidLDY(mnemonic):
		op, err = cpu.NewLDYFromBytes(addrMode, arg0, arg1)
	case cpu.IsMnemonicValidLSR(mnemonic):
		op, err = cpu.NewLSRFromBytes(addrMode, arg0, arg1)
	case cpu.IsMnemonicValidNOP(mnemonic):
		op, err = cpu.NewNOPFromBytes(addrMode, arg0, arg1)
	case cpu.IsMnemonicValidORA(mnemonic):
		op, err = cpu.NewORAFromBytes(addrMode, arg0, arg1)
	case cpu.IsMnemonicValidPHA(mnemonic):
		op, err = cpu.NewPHAFromBytes(addrMode, arg0, arg1)
	case cpu.IsMnemonicValidPHP(mnemonic):
		op, err = cpu.NewPHPFromBytes(addrMode, arg0, arg1)
	case cpu.IsMnemonicValidPLA(mnemonic):
		op, err = cpu.NewPLAFromBytes(addrMode, arg0, arg1)
	case cpu.IsMnemonicValidPLP(mnemonic):
		op, err = cpu.NewPLPFromBytes(addrMode, arg0, arg1)
	case cpu.IsMnemonicValidROL(mnemonic):
		op, err = cpu.NewROLFromBytes(addrMode, arg0, arg1)
	case cpu.IsMnemonicValidROR(mnemonic):
		op, err = cpu.NewRORFromBytes(addrMode, arg0, arg1)
	case cpu.IsMnemonicValidRTI(mnemonic):
		op, err = cpu.NewRTIFromBytes(addrMode, arg0, arg1)
	case cpu.IsMnemonicValidRTS(mnemonic):
		op, err = cpu.NewRTSFromBytes(addrMode, arg0, arg1)
	case cpu.IsMnemonicValidSBC(mnemonic):
		op, err = cpu.NewSBCFromBytes(addrMode, arg0, arg1)
	case cpu.IsMnemonicValidSEC(mnemonic):
		op, err = cpu.NewSECFromBytes(addrMode, arg0, arg1)
	case cpu.IsMnemonicValidSED(mnemonic):
		op, err = cpu.NewSEDFromBytes(addrMode, arg0, arg1)
	case cpu.IsMnemonicValidSEI(mnemonic):
		op, err = cpu.NewSEIFromBytes(addrMode, arg0, arg1)
	case cpu.IsMnemonicValidSTA(mnemonic):
		op, err = cpu.NewSTAFromBytes(addrMode, arg0, arg1)
	case cpu.IsMnemonicValidSTX(mnemonic):
		op, err = cpu.NewSTXFromBytes(addrMode, arg0, arg1)
	case cpu.IsMnemonicValidSTY(mnemonic):
		op, err = cpu.NewSTYFromBytes(addrMode, arg0, arg1)
	case cpu.IsMnemonicValidTAX(mnemonic):
		op, err = cpu.NewTAYFromBytes(addrMode, arg0, arg1)
	case cpu.IsMnemonicValidTAY(mnemonic):
		op, err = cpu.NewTAYFromBytes(addrMode, arg0, arg1)
	case cpu.IsMnemonicValidTSX(mnemonic):
		op, err = cpu.NewTSXFromBytes(addrMode, arg0, arg1)
	case cpu.IsMnemonicValidTXA(mnemonic):
		op, err = cpu.NewTXAFromBytes(addrMode, arg0, arg1)
	case cpu.IsMnemonicValidTXS(mnemonic):
		op, err = cpu.NewTXSFromBytes(addrMode, arg0, arg1)
	case cpu.IsMnemonicValidTYA(mnemonic):
		op, err = cpu.NewTYAFromBytes(addrMode, arg0, arg1)
	default:
		return nil, cpu.InvalidMnemonicError{
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
		return mnemonic, cpu.AddrModeImplied, 0, 0, nil
	}

	arg0Field := lineBySpace[1]

	if arg0Field == "A" {
		// AddrModeAccumulator
		// e.g. XXX A
		return mnemonic, cpu.AddrModeAccumulator, 0, 0, nil
	}

	switch arg0Field[0] {
	case '#':
		// AddrModeImmediate
		// e.g.: XXX %$NN
		arg0, err := strconv.ParseUint(arg0Field[2:4], 16, 8)
		if err != nil {
			return "", 0, 0, 0, err
		}

		return mnemonic, cpu.AddrModeImmediate, uint8(arg0), 0, nil
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
					return mnemonic, cpu.AddrModeZeroX, uint8(arg0), 0, nil
				case 4:
					// AddrModeAbsoluteX
					// e.g.: XXX $NNNN, X
					argWord, err := strconv.ParseUint(arg0Str, 16, 16)
					if err != nil {
						return "", 0, 0, 0, err
					}
					wordParts := util.BreakWordIntoBytes(uint16(argWord))
					return mnemonic, cpu.AddrModeAbsoluteX, wordParts[0], wordParts[1], nil
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
					return mnemonic, cpu.AddrModeZeroY, uint8(arg0), 0, nil
				case 4:
					// AddrModeAbsoluteY
					// e.g.: XXX $NNNN, Y
					argWord, err := strconv.ParseUint(arg0Str, 16, 16)
					if err != nil {
						return "", 0, 0, 0, err
					}
					wordParts := util.BreakWordIntoBytes(uint16(argWord))
					return mnemonic, cpu.AddrModeAbsoluteY, wordParts[0], wordParts[1], nil
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
				return mnemonic, cpu.AddrModeRelative, uint8(arg0), 0, nil
			case 4:
				// AddrModeAbsolute
				// e.g.: XXX $NNNN
				argWord, err := strconv.ParseUint(arg0Str, 16, 16)
				if err != nil {
					return "", 0, 0, 0, err
				}
				wordParts := util.BreakWordIntoBytes(uint16(argWord))
				return mnemonic, cpu.AddrModeAbsolute, wordParts[0], wordParts[1], nil
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
			return mnemonic, cpu.AddrModeIndirect, wordParts[0], wordParts[1], nil
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
				return mnemonic, cpu.AddrModeIndirectX, uint8(arg0), 0, nil
			} else {
				// AddrModeIndirectY
				// e.g.: XXX ($NN), Y
				arg0Str := arg0Field[2:4]
				arg0, err := strconv.ParseUint(arg0Str, 16, 8)
				if err != nil {
					return "", 0, 0, 0, err
				}
				return mnemonic, cpu.AddrModeIndirectY, uint8(arg0), 0, nil
			}
		}
	}

	return "", 0, 0, 0, cpu.InvalidSyntaxError{
		Line: line,
	}
}
