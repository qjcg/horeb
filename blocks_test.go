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

func TestRandomBlock(t *testing.T) {
	b := Blocks.RandomBlock()
	b.RandomCodePoint()
}

func BenchmarkRandomCodePoint(b *testing.B) {
	ub := &UnicodeBlock{0x0000, 0x10ffff}
	for i := 0; i < b.N; i++ {
		ub.RandomCodePoint()
	}
}

func BenchmarkRandomBlock(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Blocks.RandomBlock()
	}
}
