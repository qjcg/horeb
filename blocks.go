package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

type UnicodeBlock struct {
	start, end rune
}

type BlockMap map[string]*UnicodeBlock

func (bmap BlockMap) RandomBlock() *UnicodeBlock {
	var keys []string
	for k := range bmap {
		keys = append(keys, k)
	}
	randKey := keys[rand.Intn(len(keys))]
	return bmap[randKey]
}

func printBlocks(all bool) {
	for name, block := range Blocks {
		fmt.Printf("%5x, %5x: %s\n", block.start, block.end, name)
		if all {
			block.Print()
			fmt.Println()
		}
	}
}

func (b *UnicodeBlock) RandomCodePoint() rune {
	return rune(rand.Intn(int(b.end-b.start)) + int(b.start) + 1)
}

func (b *UnicodeBlock) Print() {
	for i := b.start; i <= b.end; i++ {
		if strconv.IsPrint(i) {
			fmt.Printf("%c ", i)
		}
	}
	fmt.Println()
}

// Print num random characters from block.
func (b *UnicodeBlock) PrintRandom(num int) {
	for i := 0; i < num; i++ {
		fmt.Printf("%c ", b.RandomCodePoint())
	}
	fmt.Println()
}

// implementing the flag.Value interface
func (b *UnicodeBlock) Set(value string) error {
	vals := strings.Split(value, ",")
	if len(vals) == 1 {
		end, err := strconv.ParseInt(vals[0], 0, 0)
		if err != nil {
			return err
		}
		b.end = rune(end)
	} else if len(vals) == 2 {
		start, err := strconv.ParseInt(vals[0], 0, 0)
		end, err := strconv.ParseInt(vals[1], 0, 0)
		if err != nil {
			return err
		}
		b.start = rune(start)
		b.end = rune(end)
	}
	return nil
}

// implementing the flag.Value interface
func (b *UnicodeBlock) String() string {
	return fmt.Sprint(*b)
}

// For info about fonts supporting specific unicode blocks, see for example:
// http://www.fileformat.info/info/unicode/block/index.htm
var Blocks = BlockMap{
	// Basic Multilingual Plane (0000-ffff)
	// https://en.wikipedia.org/wiki/Plane_(Unicode)#Basic_Multilingual_Plane
	"hebrew":         &UnicodeBlock{0x0590, 0x05ff},
	"currency":       &UnicodeBlock{0x20a0, 0x20cf},
	"letterlike":     &UnicodeBlock{0x2100, 0x214f},
	"misc_technical": &UnicodeBlock{0x2300, 0x23ff},
	"geometric":      &UnicodeBlock{0x25a0, 0x25ff},
	"misc_symbols":   &UnicodeBlock{0x2600, 0x26ff},
	"dingbats":       &UnicodeBlock{0x2700, 0x27bf},
	// Supplementary Multilingual Plane (10000-1ffff)
	// https://en.wikipedia.org/wiki/Plane_(Unicode)#Supplementary_Multilingual_Plane
	"aegean_nums":        &UnicodeBlock{0x10100, 0x1013f},
	"ancient_greek_nums": &UnicodeBlock{0x10140, 0x1018f},
	"phaistos_disc":      &UnicodeBlock{0x101d0, 0x101ff},
	"math_alnum":         &UnicodeBlock{0x1d400, 0x1d7ff},
	"emoji":              &UnicodeBlock{0x1f300, 0x1f5ff},
	"mahjong":            &UnicodeBlock{0x1f000, 0x1f02f},
	"dominos":            &UnicodeBlock{0x1f030, 0x1f09f},
	"playing_cards":      &UnicodeBlock{0x1f0a0, 0x1f0ff},
}
