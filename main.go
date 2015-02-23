package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	low := flag.Int("l", 0, "low unicode codepoint")
	high := flag.Int("h", 0x10ffff, "high unicode codepoint")
	nchars := flag.Int("n", 300, "number of characters to print")
	wrapWidth := flag.Int("w", 60, "line width")
	block := flag.String("b", "geometric", "unicode block by name")
	printMap := flag.Bool("m", false, "print Blocks map")
	flag.Parse()

	if *printMap {
		for k, v := range Blocks {
			fmt.Printf("%-15s: %U, %U\n", k, v.low, v.high)
		}
		os.Exit(0)
	}

	// If valid block provided, use it
	if *block != "" {
		b := Blocks[*block]
		if b.high != 0 {
			low = &b.low
			high = &b.high
		}
	}

	defer fmt.Println()
	lineWidth := 0
	for i := 0; i < *nchars; i++ {
		lineWidth += 2
		if lineWidth >= *wrapWidth {
			fmt.Println()
			lineWidth = 0
		}
		randomCodepoint := rand.Intn(*high-*low) + *low
		fmt.Printf("%c ", randomCodepoint)
	}
}
