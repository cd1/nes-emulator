package nes

import "testing"

func TestNES_Status(t *testing.T) {
	var system NES

	if st := system.GetStatus(); st != 0x00 {
		t.Errorf("unexpected initial status; got=%02X, want=%02X", st, 0x00)
	}

	negZero := StatusNegative | StatusZero
	system.SetStatus(negZero)

	if st := system.GetStatus(); st != negZero {
		t.Errorf("unexpected status; got=%02X, want=%02X", st, negZero)
	}

	system.SetStatus(negZero | StatusUnused)

	if st := system.GetStatus(); st != negZero {
		t.Errorf("unexpected status when setting the UNUSED flag; got=%02X, want=%02X", st, negZero)
	}
}

func BenchmarkNES_SetStatus(b *testing.B) {
	var system NES

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		system.SetStatus(0xFF)
	}
}
