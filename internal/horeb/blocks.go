package horeb

import (
	"errors"
	"fmt"
	"io"
	"math/rand"
	"sort"
	"strconv"
	"time"

	"github.com/samber/lo"
)

// UnicodeBlock values represent a contiguous range of Unicode codepoints.
type UnicodeBlock struct {
	Start, End rune
}

// Blocks is a map of short string labels to UnicodeBlock values.
var Blocks = map[string]UnicodeBlock{

	// Basic Multilingual Plane (0000-ffff)
	// https://en.wikipedia.org/wiki/Plane_(Unicode)#Basic_Multilingual_Plane
	"hebrew":         {0x0590, 0x05ff},
	"currency":       {0x20a0, 0x20cf},
	"letterlike":     {0x2100, 0x214f},
	"arrows":         {0x2190, 0x21ff},
	"misc_technical": {0x2300, 0x23ff},
	"geometric":      {0x25a0, 0x25ff},
	"misc_symbols":   {0x2600, 0x26ff},
	"dingbats":       {0x2700, 0x27bf},
	"cyrillic":       {0x0400, 0x04ff},
	"tibetian":       {0x0f00, 0x0fff},

	// Supplementary Multilingual Plane (10000-1ffff)
	// https://en.wikipedia.org/wiki/Plane_(Unicode)#Supplementary_Multilingual_Plane
	"aegean_nums":        {0x10100, 0x1013f},
	"ancient_greek_nums": {0x10140, 0x1018f},
	"phaistos_disc":      {0x101d0, 0x101ff},
	"math_alnum":         {0x1d400, 0x1d7ff},
	"emoji":              {0x1f300, 0x1f5ff},
	"mahjong":            {0x1f000, 0x1f02f},
	"dominos":            {0x1f030, 0x1f09f},
	"playing_cards":      {0x1f0a0, 0x1f0ff},
	"chess":              {0x1fa00, 0x1fa53},
	"mayan_numerals":     {0x1d2e0, 0x1d2ff},
	"gothic":             {0x10330, 0x1034f},
}

// RandomBlock returns a UnicodeBlock at random from a map[string]UnicodeBlock provided as argument.
func RandomBlock(m map[string]UnicodeBlock) (UnicodeBlock, error) {
	nkeys := len(m)
	if nkeys < 1 {
		return UnicodeBlock{}, errors.New("Empty map provided")
	}
	keys := lo.Keys[string, UnicodeBlock](m)
	randKey := lo.Sample(keys)
	return m[randKey], nil
}

// getSortedKeys returns a slice of alphabetically-sorted keys.
func getSortedKeys() []string {
	var keys []string
	for k := range Blocks {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	return keys
}

// ListBlocks prints known blocks.
func ListBlocks(w io.Writer) {
	for _, k := range getSortedKeys() {
		b := Blocks[k]
		fmt.Fprintf(w, "%5x %5x  %s\n", b.Start, b.End, k)
	}
}

// ListBlocksWithContents prints all known blocks and their contents.
func ListBlocksWithContents(w io.Writer) {
	for _, k := range getSortedKeys() {
		b := Blocks[k]
		fmt.Fprintf(w, "%5x %5x  %s\n", b.Start, b.End, k)
		b.Print(w)
		fmt.Fprintln(w)
	}
}

// RandomRune returns a rune at random from UnicodeBlock.
func (b UnicodeBlock) RandomRune() rune {
	return rune(rand.Intn(int(b.End-b.Start)) + int(b.Start) + 1)
}

// Print prints all printable runes in UnicodeBlock.
func (b UnicodeBlock) Print(w io.Writer) {
	for i := b.Start; i <= b.End; i++ {

		// Skip non-printable runes.
		if !strconv.IsPrint(i) {
			continue
		}

		fmt.Fprintf(w, "%c ", i)
	}
	fmt.Fprintln(w)
}

// PrintRandom prints n random runes from UnicodeBlock.
func (b UnicodeBlock) PrintRandom(w io.Writer, n int, ofs string) {
	for i := 0; i < n; i++ {
		fmt.Fprintf(w, "%c%s", b.RandomRune(), ofs)
	}
	fmt.Fprintln(w)
}

// StreamRandom continuously writes random runes from UnicodeBlock.
func (b UnicodeBlock) StreamRandom(w io.Writer, interval time.Duration, ofs string) {
	for range time.NewTicker(interval).C {
		fmt.Fprintf(w, "%c%s", b.RandomRune(), ofs)
	}
}
