package main

import (
	"testing"
)

func TestRandomCodePoint(t *testing.T) {
	var cp rune
	for _, b := range Blocks {
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
