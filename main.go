package main

import (
	"flag"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	dump := flag.Bool("d", false, "print all Blocks")
	list := flag.Bool("l", false, "list all Block names and codepoint ranges")
	nchars := flag.Int("n", 80, "number of characters to print")
	blocksStr := flag.String("b", "geometric", "comma-separated unicode blocks by name")
	//blockRangeStr := flag.String("r", "", "one (end) or two (start,end) comma-separated unicode blocks by hex value")
	flag.Parse()

	// Convert multi-valued arguments to slices.
	blocks := strings.Split(*blocksStr, ",")
	//blockRange := strings.Split(*blockRangeStr, ",")

	switch {
	case *list:
		printBlocks(false)
	case *dump:
		if len(blocks) != 0 {
			for _, s := range blocks {
				b := Blocks[s]
				b.Print()
			}
		} else {
			printBlocks(true)
		}
	case len(blocks) == 1:
		Blocks[blocks[0]].PrintRandom(*nchars)
	case len(blocks) > 1:
		bm := map[string]*UnicodeBlock{}
		for _, b := range blocks {
			bm[b] = Blocks[b]
		}
		for i := 0; i < *nchars; i++ {
			fmt.Printf("%c ", RandomBlock(bm).RandomRune())
		}
		fmt.Println()
	}
}
