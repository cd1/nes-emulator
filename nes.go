package nes

import (
	"bytes"
	"fmt"
	"log"
	"math"

	"github.com/cd1/nes-emulator/cpu"
	"github.com/cd1/nes-emulator/parser"
)

const (
	MemorySize          = 65536
	PRGROMStart         = 0x8000
	InitialStackAddress = 0x0100
	ResetVectorAddress  = 0xFFFC
)

type NES struct {
	CPU    CPU
	Memory Memory

	Verbose bool
}

func (nes *NES) Reset() uint8 {
	nes.CPU.ProgramCounter = 0xC000          // nes.ReadWord(ResetVectorAddress)
	nes.CPU.StackPointer = math.MaxUint8 - 2 // math.MaxUint8
	nes.CPU.SetStatus(StatusInterrupt|StatusUnused, true)

	nes.Memory = NewMemory(MemorySize)

	return 7
}

func (nes *NES) Run(game Game) error {
	resetCycles := nes.Reset()

	totalCycles := uint64(resetCycles)

	nes.loadGameInMemory(game)

	disassembleConfig := parser.DisassembleConfig{
		DisplayBytes:         true,
		DisplayMemoryAddress: true,
	}

	for {
		op, err := parser.ConvertBinaryToOperation(bytes.NewReader(nes.Memory[nes.CPU.ProgramCounter:]))
		if err != nil {
			return err
		}

		if nes.Verbose {
			var str bytes.Buffer

			if err = parser.ConvertOperationToText(op, &str, disassembleConfig, nes.CPU.ProgramCounter, nes); err != nil {
				return err
			}

			fmt.Printf("%-47v A:%02X X:%02X Y:%02X P:%02X SP:%02X PPU:%3v,%3v CYC:%v\n",
				str.String(), nes.CPU.Accumulator, nes.CPU.IndexX, nes.CPU.IndexY, nes.CPU.Status, nes.CPU.StackPointer, -1, -1, totalCycles)
		}

		cycles, err := op.ExecuteIn(nes)
		if err != nil {
			return err
		}

		totalCycles += uint64(cycles)
	}

	return nil
}

func (nes *NES) loadGameInMemory(game Game) {
	switch game.Header.PRGBankCount() {
	case 1:
		// load ROM twice, in 0x8000-0xBFFF and also in 0xC000-0xFFFF
		copy(nes.Memory[PRGROMStart:PRGROMStart+PRGBankSize], game.PRG)
		copy(nes.Memory[PRGROMStart+PRGBankSize:PRGROMStart+2*PRGBankSize], game.PRG)
	case 2:
		// load ROM in 0x8000-0xFFFF
		copy(nes.Memory[PRGROMStart:PRGROMStart+2*PRGBankSize], game.PRG)
	default:
		log.Printf("unexpected PRG bank count (%v); ROM was not loaded into memory", game.Header.PRGBankCount())
	}
}

func (nes *NES) GetAccumulator() uint8 {
	return nes.CPU.Accumulator
}

func (nes *NES) SetAccumulator(value uint8) {
	nes.CPU.Accumulator = value
}

func (nes *NES) GetIndexX() uint8 {
	return nes.CPU.IndexX
}

func (nes *NES) SetIndexX(value uint8) {
	nes.CPU.IndexX = value
}

func (nes *NES) GetIndexY() uint8 {
	return nes.CPU.IndexY
}

func (nes *NES) SetIndexY(value uint8) {
	nes.CPU.IndexY = value
}

func (nes *NES) GetProgramCounter() uint16 {
	return nes.CPU.ProgramCounter
}

func (nes *NES) SetProgramCounter(value uint16) {
	nes.CPU.ProgramCounter = value
}

func (nes *NES) GetStackPointer() uint8 {
	return nes.CPU.StackPointer
}

func (nes *NES) SetStackPointer(value uint8) {
	nes.CPU.StackPointer = value
}

func (nes *NES) GetStatus() uint8 {
	return nes.CPU.Status
}

func (nes *NES) SetStatus(value uint8) {
	// ignore StatusBreak from the received value
	effectiveValue := value & ^StatusBreak | StatusUnused
	nes.CPU.Status = effectiveValue
}

func (nes *NES) IsStatusBreak() bool {
	return nes.CPU.GetStatus(StatusBreak)
}

func (nes *NES) SetStatusBreak(isSet bool) {
	nes.CPU.SetStatus(StatusBreak, isSet)
}

func (nes *NES) IsStatusCarry() bool {
	return nes.CPU.GetStatus(StatusCarry)
}

func (nes *NES) SetStatusCarry(isSet bool) {
	nes.CPU.SetStatus(StatusCarry, isSet)
}

func (nes *NES) IsStatusDecimal() bool {
	return nes.CPU.GetStatus(StatusDecimal)
}

func (nes *NES) SetStatusDecimal(isSet bool) {
	nes.CPU.SetStatus(StatusDecimal, isSet)
}

func (nes *NES) IsStatusInterrupt() bool {
	return nes.CPU.GetStatus(StatusInterrupt)
}

func (nes *NES) SetStatusInterrupt(isSet bool) {
	nes.CPU.SetStatus(StatusInterrupt, isSet)
}

func (nes *NES) IsStatusNegative() bool {
	return nes.CPU.GetStatus(StatusNegative)
}

func (nes *NES) SetStatusNegative(isSet bool) {
	nes.CPU.SetStatus(StatusNegative, isSet)
}

func (nes *NES) IsStatusOverflow() bool {
	return nes.CPU.GetStatus(StatusOverflow)
}

func (nes *NES) SetStatusOverflow(isSet bool) {
	nes.CPU.SetStatus(StatusOverflow, isSet)
}

func (nes *NES) IsStatusUnused() bool {
	return nes.CPU.GetStatus(StatusUnused)
}

func (nes *NES) SetStatusUnused(isSet bool) {
	nes.CPU.SetStatus(StatusUnused, isSet)
}

func (nes *NES) IsStatusZero() bool {
	return nes.CPU.GetStatus(StatusZero)
}

func (nes *NES) SetStatusZero(isSet bool) {
	nes.CPU.SetStatus(StatusZero, isSet)
}

func mapMemoryAddress(address uint16) uint16 {
	var newAddress uint16

	if address < 0x2000 {
		// (0x0800, 0x0FFF): mirror of (0x0000, 0x07FF)
		// ... x1
		// (0x1800, 0x1FFF): mirror of (0x0000, 0x07FF)
		newAddress = address % 0x800
	} else if address >= 0x2008 && address < 0x4000 {
		// (0x2008, 0x200F): mirror of (0x2000, 0x2007)
		// ... x22
		// (0x3FF8, 0x3FFF): mirror of (0x2000, 0x2007)
		newAddress = 0x2000 + (address-0x2000)%0x08
	} else {
		newAddress = address
	}

	if address != newAddress {
		log.Printf("memory mapped from %04X -> %04X", address, newAddress)
	}
	return newAddress
}

func (nes *NES) ReadByte(address uint16) uint8 {
	return nes.Memory.ReadByte(mapMemoryAddress(address))
}

func (nes *NES) WriteByte(address uint16, value uint8) {
	nes.Memory.WriteByte(address, value)
}

func (nes *NES) ReadWord(address uint16) uint16 {
	return nes.Memory.ReadWord(mapMemoryAddress(address))
}

func (nes *NES) ReadWordSamePage(address uint16) uint16 {
	return nes.Memory.ReadWordSamePage(address)
}

func (nes *NES) WriteWord(address uint16, value uint16) {
	nes.Memory.WriteWord(address, value)
}

func (nes *NES) PushByteToStack(value uint8) {
	nes.WriteByte(InitialStackAddress+uint16(nes.CPU.StackPointer), value)
	nes.CPU.StackPointer--
}

func (nes *NES) PushWordToStack(value uint16) {
	nes.WriteWord(InitialStackAddress+uint16(nes.CPU.StackPointer-1), value)
	nes.CPU.StackPointer -= 2
}

func (nes *NES) PullByteFromStack() uint8 {
	nes.CPU.StackPointer++
	return nes.ReadByte(InitialStackAddress + uint16(nes.CPU.StackPointer))
}

func (nes *NES) PullWordFromStack() uint16 {
	value := nes.ReadWord(InitialStackAddress + uint16(nes.CPU.StackPointer+1))
	nes.CPU.StackPointer += 2

	return value
}

func (nes *NES) IncrementProgramCounter(value uint8) {
	nes.CPU.ProgramCounter += uint16(int8(value))
}

func inSamePage(addr0 uint16, addr1 uint16) bool {
	return addr0&0xFF00 == addr1&0xFF00
}

func (nes *NES) FetchOperand(op cpu.Operation) (uint16, uint8, bool) {
	var address uint16
	var operand uint8
	var pageCrossed bool

	switch op.AddressMode() {
	case cpu.AddrModeAccumulator:
		// operand not in memory
		operand = nes.CPU.Accumulator
	case cpu.AddrModeAbsolute:
		address = op.WordArg()
		operand = nes.ReadByte(address)
	case cpu.AddrModeAbsoluteX:
		address = op.WordArg() + uint16(nes.CPU.IndexX)
		operand = nes.ReadByte(address)
		pageCrossed = !inSamePage(op.WordArg(), address)
	case cpu.AddrModeAbsoluteY:
		address = op.WordArg() + uint16(nes.CPU.IndexY)
		operand = nes.ReadByte(address)
		pageCrossed = !inSamePage(op.WordArg(), address)
	case cpu.AddrModeImmediate:
		// operand not in memory
		operand = op.ByteArg()
	case cpu.AddrModeImplied:
		// no operand
	case cpu.AddrModeIndirect:
		address = nes.ReadWordSamePage(op.WordArg())
		operand = nes.ReadByte(address)
	case cpu.AddrModeIndirectX:
		address = nes.ReadWordSamePage(uint16(op.ByteArg() + nes.CPU.IndexX))
		operand = nes.ReadByte(address)
	case cpu.AddrModeIndirectY:
		innerAddress := nes.ReadWordSamePage(uint16(op.ByteArg()))
		address = innerAddress + uint16(nes.CPU.IndexY)
		operand = nes.ReadByte(address)
		pageCrossed = !inSamePage(innerAddress, address)
	case cpu.AddrModeRelative:
		// operand not in memory
		operand = op.ByteArg()
		pageCrossed = !inSamePage(nes.CPU.ProgramCounter+uint16(int8(op.Size()+operand)), nes.CPU.ProgramCounter+uint16(op.Size()))
	case cpu.AddrModeZero:
		address = uint16(op.ByteArg())
		operand = nes.ReadByte(address)
	case cpu.AddrModeZeroX:
		address = uint16(op.ByteArg() + nes.CPU.IndexX)
		operand = nes.ReadByte(address)
	case cpu.AddrModeZeroY:
		address = uint16(op.ByteArg() + nes.CPU.IndexY)
		operand = nes.ReadByte(address)
	default:
		log.Printf("failed to fetch operand: invalid address mode (%v)", op.AddressMode())
	}

	return address, operand, pageCrossed
}
