package util

import "testing"

func TestJoinBytesInWord(t *testing.T) {
	var b0 uint8 = 0x01
	var b1 uint8 = 0x02

	word := JoinBytesInWord([]uint8{b0, b1})
	if exp := uint16(0x0201); word != exp {
		t.Errorf("unexpected result; got=0x%04X, want=0x%04X", word, exp)
	}
}

func TestBreakWordIntoBytes(t *testing.T) {
	var word uint16 = 0x0201

	bytes := BreakWordIntoBytes(word)
	if expB0 := uint8(0x01); bytes[0] != expB0 {
		t.Errorf("unexpected byte #0 from word 0x%04X; got=0x%04X, want=0x%04X", word, bytes[0], expB0)
	}
	if expB1 := uint8(0x02); bytes[1] != expB1 {
		t.Errorf("unexpected byte #1 from word 0x%04X; got=0x%04X, want=0x%04X", word, bytes[1], expB1)
	}
}

func BenchmarkJoinBytesInWord(b *testing.B) {
	bytes := []uint8{0x01, 0x02}

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		_ = JoinBytesInWord(bytes)
	}
}

func BenchmarkBreakWordIntoBytes(b *testing.B) {
	var word uint16 = 0x0201

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		_ = BreakWordIntoBytes(word)
	}
}
