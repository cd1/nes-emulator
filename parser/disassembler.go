package parser

import (
	"encoding/binary"
	"fmt"
	"io"

	"github.com/cd1/nes-emulator/cpu"
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
		op, err := ConvertBinaryToOperation(r)
		if err != nil {
			if err == io.EOF {
				break
			}

			return err
		}

		if err = ConvertOperationToText(op, w, cfg, currentMemoryAddress, nil); err != nil {
			return err
		}
		fmt.Fprintln(w)

		currentMemoryAddress += uint16(op.Size())
	}

	return nil
}

func ConvertBinaryToOperation(r io.Reader) (cpu.Operation, error) {
	var opCode uint8
	var err error

	if err = binary.Read(r, binary.LittleEndian, &opCode); err != nil {
		return nil, err
	}

	var op cpu.Operation

	switch {
	case cpu.IsOpCodeValidADC(opCode):
		op, err = cpu.NewADCBinary(opCode, r)
	case cpu.IsOpCodeValidAND(opCode):
		op, err = cpu.NewANDBinary(opCode, r)
	case cpu.IsOpCodeValidASL(opCode):
		op, err = cpu.NewASLBinary(opCode, r)
	case cpu.IsOpCodeValidBCC(opCode):
		op, err = cpu.NewBCCBinary(opCode, r)
	case cpu.IsOpCodeValidBCS(opCode):
		op, err = cpu.NewBCSBinary(opCode, r)
	case cpu.IsOpCodeValidBEQ(opCode):
		op, err = cpu.NewBEQBinary(opCode, r)
	case cpu.IsOpCodeValidBIT(opCode):
		op, err = cpu.NewBITBinary(opCode, r)
	case cpu.IsOpCodeValidBMI(opCode):
		op, err = cpu.NewBMIBinary(opCode, r)
	case cpu.IsOpCodeValidBNE(opCode):
		op, err = cpu.NewBNEBinary(opCode, r)
	case cpu.IsOpCodeValidBPL(opCode):
		op, err = cpu.NewBPLBinary(opCode, r)
	case cpu.IsOpCodeValidBRK(opCode):
		op, err = cpu.NewBRKBinary(opCode, r)
	case cpu.IsOpCodeValidBVC(opCode):
		op, err = cpu.NewBVCBinary(opCode, r)
	case cpu.IsOpCodeValidBVS(opCode):
		op, err = cpu.NewBVSBinary(opCode, r)
	case cpu.IsOpCodeValidCLC(opCode):
		op, err = cpu.NewCLCBinary(opCode, r)
	case cpu.IsOpCodeValidCLD(opCode):
		op, err = cpu.NewCLDBinary(opCode, r)
	case cpu.IsOpCodeValidCLI(opCode):
		op, err = cpu.NewCLIBinary(opCode, r)
	case cpu.IsOpCodeValidCLV(opCode):
		op, err = cpu.NewCLVBinary(opCode, r)
	case cpu.IsOpCodeValidCMP(opCode):
		op, err = cpu.NewCMPBinary(opCode, r)
	case cpu.IsOpCodeValidCPX(opCode):
		op, err = cpu.NewCPXBinary(opCode, r)
	case cpu.IsOpCodeValidCPY(opCode):
		op, err = cpu.NewCPYBinary(opCode, r)
	case cpu.IsOpCodeValidDCP(opCode):
		op, err = cpu.NewDCPBinary(opCode, r)
	case cpu.IsOpCodeValidDEC(opCode):
		op, err = cpu.NewDECBinary(opCode, r)
	case cpu.IsOpCodeValidDEX(opCode):
		op, err = cpu.NewDEXBinary(opCode, r)
	case cpu.IsOpCodeValidDEY(opCode):
		op, err = cpu.NewDEYBinary(opCode, r)
	case cpu.IsOpCodeValidEOR(opCode):
		op, err = cpu.NewEORBinary(opCode, r)
	case cpu.IsOpCodeValidINC(opCode):
		op, err = cpu.NewINCBinary(opCode, r)
	case cpu.IsOpCodeValidINX(opCode):
		op, err = cpu.NewINXBinary(opCode, r)
	case cpu.IsOpCodeValidINY(opCode):
		op, err = cpu.NewINYBinary(opCode, r)
	case cpu.IsOpCodeValidISB(opCode):
		op, err = cpu.NewISBBinary(opCode, r)
	case cpu.IsOpCodeValidJMP(opCode):
		op, err = cpu.NewJMPBinary(opCode, r)
	case cpu.IsOpCodeValidJSR(opCode):
		op, err = cpu.NewJSRBinary(opCode, r)
	case cpu.IsOpCodeValidLAX(opCode):
		op, err = cpu.NewLAXBinary(opCode, r)
	case cpu.IsOpCodeValidLDA(opCode):
		op, err = cpu.NewLDABinary(opCode, r)
	case cpu.IsOpCodeValidLDX(opCode):
		op, err = cpu.NewLDXBinary(opCode, r)
	case cpu.IsOpCodeValidLDY(opCode):
		op, err = cpu.NewLDYBinary(opCode, r)
	case cpu.IsOpCodeValidLSR(opCode):
		op, err = cpu.NewLSRBinary(opCode, r)
	case cpu.IsOpCodeValidNOP(opCode):
		op, err = cpu.NewNOPBinary(opCode, r)
	case cpu.IsOpCodeValidORA(opCode):
		op, err = cpu.NewORABinary(opCode, r)
	case cpu.IsOpCodeValidPHA(opCode):
		op, err = cpu.NewPHABinary(opCode, r)
	case cpu.IsOpCodeValidPHP(opCode):
		op, err = cpu.NewPHPBinary(opCode, r)
	case cpu.IsOpCodeValidPLA(opCode):
		op, err = cpu.NewPLABinary(opCode, r)
	case cpu.IsOpCodeValidPLP(opCode):
		op, err = cpu.NewPLPBinary(opCode, r)
	case cpu.IsOpCodeValidRLA(opCode):
		op, err = cpu.NewRLABinary(opCode, r)
	case cpu.IsOpCodeValidROL(opCode):
		op, err = cpu.NewROLBinary(opCode, r)
	case cpu.IsOpCodeValidROR(opCode):
		op, err = cpu.NewRORBinary(opCode, r)
	case cpu.IsOpCodeValidRRA(opCode):
		op, err = cpu.NewRRABinary(opCode, r)
	case cpu.IsOpCodeValidRTI(opCode):
		op, err = cpu.NewRTIBinary(opCode, r)
	case cpu.IsOpCodeValidRTS(opCode):
		op, err = cpu.NewRTSBinary(opCode, r)
	case cpu.IsOpCodeValidSAX(opCode):
		op, err = cpu.NewSAXBinary(opCode, r)
	case cpu.IsOpCodeValidSBC(opCode):
		op, err = cpu.NewSBCBinary(opCode, r)
	case cpu.IsOpCodeValidSEC(opCode):
		op, err = cpu.NewSECBinary(opCode, r)
	case cpu.IsOpCodeValidSED(opCode):
		op, err = cpu.NewSEDBinary(opCode, r)
	case cpu.IsOpCodeValidSEI(opCode):
		op, err = cpu.NewSEIBinary(opCode, r)
	case cpu.IsOpCodeValidSLO(opCode):
		op, err = cpu.NewSLOBinary(opCode, r)
	case cpu.IsOpCodeValidSRE(opCode):
		op, err = cpu.NewSREBinary(opCode, r)
	case cpu.IsOpCodeValidSTA(opCode):
		op, err = cpu.NewSTABinary(opCode, r)
	case cpu.IsOpCodeValidSTX(opCode):
		op, err = cpu.NewSTXBinary(opCode, r)
	case cpu.IsOpCodeValidSTY(opCode):
		op, err = cpu.NewSTYBinary(opCode, r)
	case cpu.IsOpCodeValidTAX(opCode):
		op, err = cpu.NewTAXBinary(opCode, r)
	case cpu.IsOpCodeValidTAY(opCode):
		op, err = cpu.NewTAYBinary(opCode, r)
	case cpu.IsOpCodeValidTSX(opCode):
		op, err = cpu.NewTSXBinary(opCode, r)
	case cpu.IsOpCodeValidTXA(opCode):
		op, err = cpu.NewTXABinary(opCode, r)
	case cpu.IsOpCodeValidTXS(opCode):
		op, err = cpu.NewTXSBinary(opCode, r)
	case cpu.IsOpCodeValidTYA(opCode):
		op, err = cpu.NewTYABinary(opCode, r)
	default:
		return nil, cpu.InvalidOpCodeError{
			OpCode: opCode,
		}
	}

	if err != nil {
		return nil, err
	}

	return op, nil
}

func ConvertOperationToText(op cpu.Operation, w io.Writer, cfg DisassembleConfig, memoryAddress uint16, env cpu.OperationEnvironment) error {
	if cfg.DisplayMemoryAddress {
		if _, err := fmt.Fprintf(w, "%04X  ", memoryAddress); err != nil {
			return err
		}
	}

	if cfg.DisplayBytes {
		var str string

		switch op.Size() {
		case 1:
			str = fmt.Sprintf("%02X", op.Code())
		case 2:
			str = fmt.Sprintf("%02X %02X", op.Code(), op.ByteArg())
		case 3:
			args := util.BreakWordIntoBytes(op.WordArg())
			str = fmt.Sprintf("%02X %02X %02X", op.Code(), args[0], args[1])
		}

		if _, err := fmt.Fprintf(w, "%-8v ", str); err != nil {
			return err
		}
	}

	if _, err := fmt.Fprint(w, op.StringWithEnv(env)); err != nil {
		return err
	}

	return nil
}
