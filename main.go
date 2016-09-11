package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"time"
)

const (
	description = "horeb: Speaking in tongues via stdout"
)

func usage() {
	fmt.Printf("\n%s\n\n", description)
	flag.PrintDefaults()
	fmt.Println()
}

func main() {
	rand.Seed(time.Now().UnixNano())

	flag.Usage = usage

	color := flag.Bool("c", false, "colorize output")
	dump := flag.Bool("d", false, "print all Blocks")
	list := flag.Bool("l", false, "list all Block names and codepoint ranges")
	nchars := flag.Int("n", 30, "number of characters to print")
	flag.Parse()

	blocks := []string{"all"}
	if flag.NArg() > 0 {
		blocks = flag.Args()
		// special all value means all blocks
		if blocks[0] == "all" {
			// FIXME: remove blocks[0] ("all") to avoid error
			for k := range Blocks {
				blocks = append(blocks, k)
			}
		}
	}

	switch {
	case *color:
		// TODO: implement colorized output
		return
	case *list:
		printBlocks(false)
	case *dump:
		printBlocks(true)
	case len(blocks) == 1:
		if b, ok := Blocks[blocks[0]]; ok {
			b.PrintRandom(*nchars)
		} else {
			log.Fatalf("Unknown block name: %s\n", blocks[0])
		}
	case len(blocks) > 1:
		bm := map[string]*UnicodeBlock{}
		for _, b := range blocks {
			if val, ok := Blocks[b]; ok {
				bm[b] = val
			} else {
				log.Printf("Unknown block name: %s\n", b)
			}
		}
		if len(bm) > 0 {
			for i := 0; i < *nchars; i++ {
				block, err := RandomBlock(bm)
				if err != nil {
					log.Fatal(err)
				}
				fmt.Printf("%c ", block.RandomRune())
			}
		}
		fmt.Println()
	}
}
