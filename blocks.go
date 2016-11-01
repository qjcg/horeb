package main

import (
	"errors"
	"fmt"
	"math/rand"
	"sort"
	"strconv"
)

// UnicodeBlock values represent a contiguous range of Unicode codepoints.
type UnicodeBlock struct {
	Start, End rune
}

// Blocks is a map of short string labels to UnicodeBlock values.
var Blocks = map[string]UnicodeBlock{

	// Basic Multilingual Plane (0000-ffff)
	// https://en.wikipedia.org/wiki/Plane_(Unicode)#Basic_Multilingual_Plane
	"hebrew":         UnicodeBlock{0x0590, 0x05ff},
	"currency":       UnicodeBlock{0x20a0, 0x20cf},
	"letterlike":     UnicodeBlock{0x2100, 0x214f},
	"misc_technical": UnicodeBlock{0x2300, 0x23ff},
	"geometric":      UnicodeBlock{0x25a0, 0x25ff},
	"misc_symbols":   UnicodeBlock{0x2600, 0x26ff},
	"dingbats":       UnicodeBlock{0x2700, 0x27bf},

	// Supplementary Multilingual Plane (10000-1ffff)
	// https://en.wikipedia.org/wiki/Plane_(Unicode)#Supplementary_Multilingual_Plane
	"aegean_nums":        UnicodeBlock{0x10100, 0x1013f},
	"ancient_greek_nums": UnicodeBlock{0x10140, 0x1018f},
	"phaistos_disc":      UnicodeBlock{0x101d0, 0x101ff},
	"math_alnum":         UnicodeBlock{0x1d400, 0x1d7ff},
	"emoji":              UnicodeBlock{0x1f300, 0x1f5ff},
	"mahjong":            UnicodeBlock{0x1f000, 0x1f02f},
	"dominos":            UnicodeBlock{0x1f030, 0x1f09f},
	"playing_cards":      UnicodeBlock{0x1f0a0, 0x1f0ff},
}

// RandomBlock returns a UnicodeBlock at random from a map[string]UnicodeBlock provided as argument.
func RandomBlock(m map[string]UnicodeBlock) (UnicodeBlock, error) {
	if len(m) == 0 {
		return UnicodeBlock{}, errors.New("Empty map provided")
	}
	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	randKey := keys[rand.Intn(len(keys))]
	return m[randKey], nil
}

func printBlocks(all bool) {
	// Create a slice of alphabetically-sorted keys.
	var keys []string
	for k := range Blocks {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		b := Blocks[k]
		fmt.Printf("%5x %5x  %s\n", b.Start, b.End, k)
		if all {
			b.Print()
			fmt.Println()
		}
	}
}

// RandomRune returns a single rune at random from UnicodeBlock.
func (b UnicodeBlock) RandomRune() rune {
	return rune(rand.Intn(int(b.End-b.Start)) + int(b.Start) + 1)
}

// Print prints all printable runes in UnicodeBlock.
func (b UnicodeBlock) Print() {
	for i := b.Start; i <= b.End; i++ {
		if strconv.IsPrint(i) {
			fmt.Printf("%c ", i)
		}
	}
	fmt.Println()
}

// PrintRandom prints n random runes from UnicodeBlock.
func (b UnicodeBlock) PrintRandom(n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%c ", b.RandomRune())
	}
	fmt.Println()
}
