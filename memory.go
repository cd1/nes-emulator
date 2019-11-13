package nes

import "github.com/cd1/nes-emulator/util"

type Memory []uint8

func NewMemory(size uint64) Memory {
	return Memory(make([]uint8, size))
}

func (m Memory) ReadByte(address uint16) uint8 {
	return m[address]
}

func (m Memory) ReadWord(address uint16) uint16 {
	return util.JoinBytesInWord(m[address : address+2])
}

func (m Memory) ReadWordSamePage(address uint16) uint16 {
	addr0 := address
	addr1 := address&0xFF00 + uint16(uint8(address&0x00FF)+1)

	return util.JoinBytesInWord([]uint8{m[addr0], m[addr1]})
}

func (m Memory) WriteByte(address uint16, value uint8) {
	m[address] = value
}

func (m Memory) WriteWord(address uint16, value uint16) {
	copy(m[address:address+2], util.BreakWordIntoBytes(value))
}
