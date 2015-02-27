package main

import (
	"testing"
)

func TestRandomCodePoint(t *testing.T) {
	for _, b := range Blocks {
		var cp rune
		cp = b.RandomCodePoint()

		if cp < b.start || cp > b.end {
			t.Fail()
		}
	}
}

func BenchmarkRandomCodePoint(b *testing.B) {
	testBlock := &UnicodeBlock{0x0000, 0x10ffff}
	for i := 0; i < b.N; i++ {
		testBlock.RandomCodePoint()
	}
}
