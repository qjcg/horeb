package horeb

import (
	"testing"
)

func TestRandomRune(t *testing.T) {
	var r rune
	for _, b := range Blocks {
		r = b.RandomRune()
		if r < b.Start || r > b.End {
			t.Errorf("rune %c outside of expected range: %x - %x\n", r, b.Start, b.End)
		}
	}
}

func TestRandomBlock(t *testing.T) {
	b, err := RandomBlock(Blocks)
	if err != nil {
		t.Fatal("RandomBlock error", err)
	}
	b.RandomRune()
}

func BenchmarkRandomRune(b *testing.B) {
	ub := UnicodeBlock{0x0000, 0x10ffff}
	for i := 0; i < b.N; i++ {
		ub.RandomRune()
	}
}

func BenchmarkRandomBlock(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if _, err := RandomBlock(Blocks); err != nil {
			b.Error(err)
		}
	}
}
