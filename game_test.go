package nes

import (
	"os"
	"testing"
)

const nesTestFileName = "sample/nestest.nes"

func TestLoadGame(t *testing.T) {
	nesTestFile, err := os.Open(nesTestFileName)
	if err != nil {
		t.Skipf("failed to open nestest: %v", err)
	}
	defer nesTestFile.Close()

	if _, err := LoadGame(nesTestFile); err != nil {
		t.Fatal(err)
	}
}

func BenchmarkLoadGame(b *testing.B) {
	b.StopTimer()

	for n := 0; n < b.N; n++ {
		nesTestFile, err := os.Open(nesTestFileName)
		if err != nil {
			b.Skipf("failed to open nestest: %v", err)
		}

		b.StartTimer()

		if _, err = LoadGame(nesTestFile); err != nil {
			b.Fatal(err)
		}

		b.StopTimer()

		if err = nesTestFile.Close(); err != nil {
			b.Logf("failed to close nestest: %v", err)
		}
	}
}
