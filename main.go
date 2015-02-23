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
	plane := flag.String("p", "geometric", "unicode plane by name")
	printMap := flag.Bool("m", false, "print Planes map")
	flag.Parse()

	if *printMap {
		for k, v := range Planes {
			fmt.Printf("%-15s: %U, %U\n", k, v.low, v.high)
		}
		os.Exit(0)
	}

	// If valid plane provided as flag, use it
	if *plane != "" {
		p := Planes[*plane]
		if p.high != 0 {
			low = &p.low
			high = &p.high
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
