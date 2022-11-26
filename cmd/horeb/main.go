package main

import (
	"errors"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/qjcg/horeb/internal/horeb"
	"golang.org/x/exp/slog"
)

const (
	description = "horeb: Speaking in tongues via stdout"
)

func usage() {
	fmt.Fprintf(flag.CommandLine.Output(), "\n%s\n\n", description)
	flag.PrintDefaults()
	fmt.Fprintln(flag.CommandLine.Output())
}

func main() {
	rand.Seed(time.Now().UnixNano())

	flag.Usage = usage

	dumpFlag := flag.Bool("d", false, "dump all blocks")
	listFlag := flag.Bool("l", false, "list all blocks")
	nCharsFlag := flag.Int("n", 30, "number of runes to generate")
	ofsFlag := flag.String("o", " ", "output field separator")
	streamFlag := flag.Bool("s", false, "generate an endless stream of runes")
	streamDelayFlag := flag.Duration("D", time.Millisecond*30, "stream delay")
	versionFlag := flag.Bool("v", false, "print version")
	flag.Parse()

	if *versionFlag {
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
	case *listFlag:
		horeb.ListBlocks(os.Stdout)
	case *dumpFlag:
		horeb.DumpBlocks(os.Stdout)
	case len(blocks) == 1:
		b, ok := horeb.Blocks[blocks[0]]
		if !ok {
			err := errors.New("unknown block")
			slog.Error("Unknown block", err, "block", blocks[0])
		}

		if *streamFlag {
			ticker := time.NewTicker(*streamDelayFlag)
			for range ticker.C {
				fmt.Printf("%c%s", b.RandomRune(), *ofsFlag)
			}
		} else {
			b.PrintRandom(os.Stdout, *nCharsFlag, *ofsFlag)
		}
	case len(blocks) > 1:
		bm := map[string]horeb.UnicodeBlock{}
		for _, b := range blocks {
			val, ok := horeb.Blocks[b]
			if !ok {
				slog.Warn("Unknown block", "block", b)
				continue
			}
			bm[b] = val
		}
		if len(bm) > 0 {
			defer fmt.Println()
			if *streamFlag {
				ticker := time.NewTicker(*streamDelayFlag)
				for range ticker.C {

					block, err := horeb.RandomBlock(bm)
					if err != nil {
						slog.Error("Error getting random block", err)
						os.Exit(1)
					}
					fmt.Printf("%c%s", block.RandomRune(), *ofsFlag)
				}
			} else {
				for i := 0; i < *nCharsFlag; i++ {
					block, err := horeb.RandomBlock(bm)
					if err != nil {
						slog.Error("Error getting random block", err)
						os.Exit(1)
					}
					fmt.Printf("%c%s", block.RandomRune(), *ofsFlag)
				}
			}
		}
	}
}
