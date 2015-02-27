package main

import (
	"flag"
	"math/rand"
	"time"
)

var blockFlag UnicodeBlock

func init() {
	flag.Var(&blockFlag, "r", "codepoint range")
}

func main() {
	rand.Seed(time.Now().UnixNano())

	dump := flag.Bool("d", false, "print all Blocks")
	list := flag.Bool("l", false, "list all Blocks")
	nchars := flag.Int("n", 80, "number of characters to print")
	block := flag.String("b", "", "unicode block by name")
	flag.Parse()

	// If invalid Blocks key specified, use a default instead
	b, valid := Blocks[*block]
	if !valid {
		b = Blocks["geometric"]
	}

	switch {
	case *dump:
		if *block != "" {
			b.Print()
		} else {
			printBlocks(true)
		}
	case *list:
		printBlocks(false)
	case blockFlag.end > 0:
		b.start = blockFlag.start
		b.end = blockFlag.end
		b.PrintRandom(*nchars)
	default:
		b.PrintRandom(*nchars)
	}
}
