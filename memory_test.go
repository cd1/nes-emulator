package nes

import "testing"

const memorySize = 1024

func TestMemory_Byte(t *testing.T) {
	m := NewMemory(memorySize)

	var addr uint16 = 0x12
	var value uint8 = 0x34

	m.WriteByte(addr, value)
	if valueRead := m.ReadByte(addr); valueRead != value {
		t.Errorf("unexpected value read from memory; got=%v, want=%v", valueRead, value)
	}
}

func TestMemory_Word(t *testing.T) {
	m := NewMemory(memorySize)

	var addr uint16 = 0x12
	var value uint16 = 0x3456

	m.WriteWord(addr, value)
	if valueRead := m.ReadWord(addr); valueRead != value {
		t.Errorf("unexpected value read from memory; got=%v, want=%v", valueRead, value)
	}
}

func BenchmarkNewMemory(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = NewMemory(memorySize)
	}
}

func BenchmarkMemory_ReadByte(b *testing.B) {
	var addr uint16 = 0x12

	mem := NewMemory(memorySize)

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		_ = mem.ReadByte(addr)
	}
}

func BenchmarkMemory_ReadWord(b *testing.B) {
	var addr uint16 = 0x12

	mem := NewMemory(memorySize)

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		_ = mem.ReadWord(addr)
	}
}

func BenchmarkMemory_WriteByte(b *testing.B) {
	var addr uint16 = 0x12
	var value uint8 = 0x34

	mem := NewMemory(memorySize)

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		mem.WriteByte(addr, value)
	}
}

func BenchmarkMemory_WriteWord(b *testing.B) {
	var addr uint16 = 0x12
	var value uint16 = 0x3456

	mem := NewMemory(memorySize)

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		mem.WriteWord(addr, value)
	}
}
