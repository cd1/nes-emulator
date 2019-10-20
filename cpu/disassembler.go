package cpu

import (
	"encoding/binary"
	"fmt"
	"io"

	"github.com/cd1/nes-emulator/cpu/operation"
	"github.com/cd1/nes-emulator/util"
)

type DisassembleConfig struct {
	DisplayMemoryAddress bool
	DisplayBytes         bool
}

// Disassemble parses a binary stream and returns a list of
// 6502 operations. If the format is invalid, an error is returned.
func Disassemble(r io.Reader, w io.Writer, cfg DisassembleConfig) error {
	var currentMemoryAddress uint16 = 0x0600

	for {
		op, err := disassembleOperation(r)
		if err != nil {
			if err == io.EOF {
				break
			}

			return err
		}

		if err = convertOperationToText(op, w, cfg, &currentMemoryAddress); err != nil {
			return err
		}
	}

	return nil
}

func disassembleOperation(r io.Reader) (operation.Operation, error) {
	var opCode uint8
	var err error

	if err = binary.Read(r, binary.LittleEndian, &opCode); err != nil {
		return nil, err
	}

	var op operation.Operation

	switch {
	case operation.IsOpCodeValidADC(opCode):
		op, err = operation.NewADCBinary(opCode, r)
	case operation.IsOpCodeValidAND(opCode):
		op, err = operation.NewANDBinary(opCode, r)
	case operation.IsOpCodeValidASL(opCode):
		op, err = operation.NewASLBinary(opCode, r)
	case operation.IsOpCodeValidBCC(opCode):
		op, err = operation.NewBCCBinary(opCode, r)
	case operation.IsOpCodeValidBCS(opCode):
		op, err = operation.NewBCSBinary(opCode, r)
	case operation.IsOpCodeValidBEQ(opCode):
		op, err = operation.NewBEQBinary(opCode, r)
	case operation.IsOpCodeValidBIT(opCode):
		op, err = operation.NewBITBinary(opCode, r)
	case operation.IsOpCodeValidBMI(opCode):
		op, err = operation.NewBMIBinary(opCode, r)
	case operation.IsOpCodeValidBNE(opCode):
		op, err = operation.NewBNEBinary(opCode, r)
	case operation.IsOpCodeValidBPL(opCode):
		op, err = operation.NewBPLBinary(opCode, r)
	case operation.IsOpCodeValidBRK(opCode):
		op, err = operation.NewBRKBinary(opCode, r)
	case operation.IsOpCodeValidBVC(opCode):
		op, err = operation.NewBVCBinary(opCode, r)
	case operation.IsOpCodeValidBVS(opCode):
		op, err = operation.NewBVSBinary(opCode, r)
	case operation.IsOpCodeValidCLC(opCode):
		op, err = operation.NewCLCBinary(opCode, r)
	case operation.IsOpCodeValidCLD(opCode):
		op, err = operation.NewCLDBinary(opCode, r)
	case operation.IsOpCodeValidCLI(opCode):
		op, err = operation.NewCLIBinary(opCode, r)
	case operation.IsOpCodeValidCLV(opCode):
		op, err = operation.NewCLVBinary(opCode, r)
	case operation.IsOpCodeValidCMP(opCode):
		op, err = operation.NewCMPBinary(opCode, r)
	case operation.IsOpCodeValidCPX(opCode):
		op, err = operation.NewCPXBinary(opCode, r)
	case operation.IsOpCodeValidCPY(opCode):
		op, err = operation.NewCPYBinary(opCode, r)
	case operation.IsOpCodeValidDEC(opCode):
		op, err = operation.NewDECBinary(opCode, r)
	case operation.IsOpCodeValidDEX(opCode):
		op, err = operation.NewDEXBinary(opCode, r)
	case operation.IsOpCodeValidDEY(opCode):
		op, err = operation.NewDEYBinary(opCode, r)
	case operation.IsOpCodeValidEOR(opCode):
		op, err = operation.NewEORBinary(opCode, r)
	case operation.IsOpCodeValidINC(opCode):
		op, err = operation.NewINCBinary(opCode, r)
	case operation.IsOpCodeValidINX(opCode):
		op, err = operation.NewINXBinary(opCode, r)
	case operation.IsOpCodeValidINY(opCode):
		op, err = operation.NewINYBinary(opCode, r)
	case operation.IsOpCodeValidJMP(opCode):
		op, err = operation.NewJMPBinary(opCode, r)
	case operation.IsOpCodeValidJSR(opCode):
		op, err = operation.NewJSRBinary(opCode, r)
	case operation.IsOpCodeValidLDA(opCode):
		op, err = operation.NewLDABinary(opCode, r)
	case operation.IsOpCodeValidLDX(opCode):
		op, err = operation.NewLDXBinary(opCode, r)
	case operation.IsOpCodeValidLDY(opCode):
		op, err = operation.NewLDYBinary(opCode, r)
	case operation.IsOpCodeValidLSR(opCode):
		op, err = operation.NewLSRBinary(opCode, r)
	case operation.IsOpCodeValidNOP(opCode):
		op, err = operation.NewNOPBinary(opCode, r)
	case operation.IsOpCodeValidORA(opCode):
		op, err = operation.NewORABinary(opCode, r)
	case operation.IsOpCodeValidPHA(opCode):
		op, err = operation.NewPHABinary(opCode, r)
	case operation.IsOpCodeValidPHP(opCode):
		op, err = operation.NewPHPBinary(opCode, r)
	case operation.IsOpCodeValidPLA(opCode):
		op, err = operation.NewPLABinary(opCode, r)
	case operation.IsOpCodeValidPLP(opCode):
		op, err = operation.NewPLPBinary(opCode, r)
	case operation.IsOpCodeValidROL(opCode):
		op, err = operation.NewROLBinary(opCode, r)
	case operation.IsOpCodeValidROR(opCode):
		op, err = operation.NewRORBinary(opCode, r)
	case operation.IsOpCodeValidRTI(opCode):
		op, err = operation.NewRTIBinary(opCode, r)
	case operation.IsOpCodeValidRTS(opCode):
		op, err = operation.NewRTSBinary(opCode, r)
	case operation.IsOpCodeValidSBC(opCode):
		op, err = operation.NewSBCBinary(opCode, r)
	case operation.IsOpCodeValidSEC(opCode):
		op, err = operation.NewSECBinary(opCode, r)
	case operation.IsOpCodeValidSED(opCode):
		op, err = operation.NewSEDBinary(opCode, r)
	case operation.IsOpCodeValidSEI(opCode):
		op, err = operation.NewSEIBinary(opCode, r)
	case operation.IsOpCodeValidSTA(opCode):
		op, err = operation.NewSTABinary(opCode, r)
	case operation.IsOpCodeValidSTX(opCode):
		op, err = operation.NewSTXBinary(opCode, r)
	case operation.IsOpCodeValidSTY(opCode):
		op, err = operation.NewSTYBinary(opCode, r)
	case operation.IsOpCodeValidTAX(opCode):
		op, err = operation.NewTAXBinary(opCode, r)
	case operation.IsOpCodeValidTAY(opCode):
		op, err = operation.NewTAYBinary(opCode, r)
	case operation.IsOpCodeValidTSX(opCode):
		op, err = operation.NewTSXBinary(opCode, r)
	case operation.IsOpCodeValidTXA(opCode):
		op, err = operation.NewTXABinary(opCode, r)
	case operation.IsOpCodeValidTXS(opCode):
		op, err = operation.NewTXSBinary(opCode, r)
	case operation.IsOpCodeValidTYA(opCode):
		op, err = operation.NewTYABinary(opCode, r)
	default:
		return nil, operation.InvalidOpCodeError{
			OpCode: opCode,
		}
	}

	if err != nil {
		return nil, err
	}

	return op, nil
}

func convertOperationToText(op operation.Operation, w io.Writer, cfg DisassembleConfig, memoryAddress *uint16) error {
	if cfg.DisplayMemoryAddress {
		if _, err := fmt.Fprintf(w, "%04X\t", *memoryAddress); err != nil {
			return err
		}

		*memoryAddress += uint16(op.Size())
	}

	if cfg.DisplayBytes {
		switch op.Size() {
		case 1:
			if _, err := fmt.Fprintf(w, "$%02X        \t", op.Code()); err != nil {
				return err
			}
		case 2:
			if _, err := fmt.Fprintf(w, "$%02X $%02X    \t", op.Code(), op.ByteArg()); err != nil {
				return err
			}
		case 3:
			args := util.BreakWordIntoBytes(op.WordArg())
			if _, err := fmt.Fprintf(w, "$%02X $%02X $%02X\t", op.Code(), args[0], args[1]); err != nil {
				return err
			}
		}
	}

	if _, err := fmt.Fprintln(w, op); err != nil {
		return err
	}

	return nil
}
