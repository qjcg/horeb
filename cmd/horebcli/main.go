package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/qjcg/horeb"
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

	dump := flag.Bool("d", false, "dump all blocks")
	list := flag.Bool("l", false, "list all blocks")
	nchars := flag.Int("n", 30, "number of runes to generate")
	ofs := flag.String("o", " ", "output field separator")
	version := flag.Bool("v", false, "print version")
	flag.Parse()

	if *version {
		fmt.Println(Version)
		return
	}

	blocks := []string{"all"}
	if flag.NArg() > 0 {
		blocks = flag.Args()
	}
	// special value means all blocks
	if blocks[0] == "all" {
		// remove "all" value after use
		blocks = blocks[:0]
		for k := range horeb.Blocks {
			blocks = append(blocks, k)
		}
	}

	switch {
	case *list:
		horeb.PrintBlocks(false)
	case *dump:
		horeb.PrintBlocks(true)
	case len(blocks) == 1:
		b, ok := horeb.Blocks[blocks[0]]
		if !ok {
			log.Fatalf("Unknown block: %s\n", blocks[0])
		}
		b.PrintRandom(*nchars, *ofs)
	case len(blocks) > 1:
		bm := map[string]horeb.UnicodeBlock{}
		for _, b := range blocks {
			val, ok := horeb.Blocks[b]
			if !ok {
				log.Printf("Unknown block: %s\n", b)
				continue
			}
			bm[b] = val
		}
		if len(bm) > 0 {
			for i := 0; i < *nchars; i++ {
				block, err := horeb.RandomBlock(bm)
				if err != nil {
					log.Fatal(err)
				}
				fmt.Printf("%c%s", block.RandomRune(), *ofs)
			}
		}
		fmt.Println()
	}
}
