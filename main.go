package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func printBlocks(all bool) {
	for name, block := range Blocks {
		fmt.Printf("%s: %U, %U\n", name, block.start, block.end)
		if all {
			block.Print()
			fmt.Println()
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	dumpBlocks := flag.Bool("d", false, "print all Blocks")
	listBlocks := flag.Bool("l", false, "list all Blocks")
	nchars := flag.Int("n", 300, "number of characters to print")
	block := flag.String("b", "geometric", "unicode block by name")
	flag.Parse()

	if *dumpBlocks {
		printBlocks(true)
		os.Exit(0)
	}

	if *listBlocks {
		printBlocks(false)
		os.Exit(0)
	}


	b, valid := Blocks[*block]
	// If invalid block passed in, use a default instead
	if !valid {
		b = Blocks["geometric"]
	}
	b.PrintRandom(*nchars)
}
