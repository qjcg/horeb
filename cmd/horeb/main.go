package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"

	"github.com/qjcg/horeb/internal/horeb"
	"github.com/samber/mo"
	"golang.org/x/exp/slog"
)

const (
	description = "horeb: Speaking in tongues via stdout"
)

// config contains our application config.
type config struct {
	debug                  *bool
	listBlocks             *bool
	listBlocksWithContents *bool
	nChars                 *int
	ofs                    *string
	stream                 *bool
	streamDelay            *time.Duration
	version                *bool
	blocks                 []string
}

func getConf(w io.Writer, args []string) mo.Result[*config] {
	var err error

	fs := flag.NewFlagSet("horeb", flag.ExitOnError)

	fs.Usage = func() {
		fmt.Fprintf(w, "\n%s\n\n", description)
		fs.PrintDefaults()
		fmt.Fprintln(w)
	}

	conf := config{
		debug:                  fs.Bool("d", false, "print debug log messages"),
		listBlocks:             fs.Bool("l", false, "list all blocks"),
		listBlocksWithContents: fs.Bool("L", false, "list all blocks along with their contents"),
		nChars:                 fs.Int("n", 30, "number of runes to generate"),
		ofs:                    fs.String("o", " ", "output field separator"),
		stream:                 fs.Bool("s", false, "generate an endless stream of runes"),
		streamDelay:            fs.Duration("D", time.Millisecond*30, "stream delay"),
		version:                fs.Bool("v", false, "print version"),
	}
	if err = fs.Parse(args); err != nil {
		return mo.Err[*config](err)
	}

	conf.blocks = []string{"all"}
	if fs.NArg() > 0 {
		conf.blocks = fs.Args()
	}

	slog.Debug("configuration parsed from command line args", "conf", conf, "args", fs.Args())

	return mo.Ok(&conf)
}

func main() {
	os.Exit(Main())
}

func Main() int {
	rand.Seed(time.Now().UnixNano())

	conf, err := getConf(os.Stderr, os.Args[1:]).Get()
	if err != nil {
		slog.Error("error getting flags", err)
		return 1
	}

	if *conf.version {
		fmt.Println(Version)
		return 0
	}

	// special value means all blocks
	if conf.blocks[0] == "all" {
		// remove "all" value after use
		conf.blocks = conf.blocks[:0]
		for k := range horeb.Blocks {
			conf.blocks = append(conf.blocks, k)
		}
	}

	switch {
	case *conf.listBlocks:
		horeb.ListBlocks(os.Stdout)
	case *conf.listBlocksWithContents:
		horeb.ListBlocksWithContents(os.Stdout)

	// PrintRandom or StreamRandom from a _single_ block.
	case len(conf.blocks) == 1:
		b, ok := horeb.Blocks[conf.blocks[0]]
		if !ok {
			err := errors.New("unknown block")
			slog.Error("Unknown block", err, "block", conf.blocks[0])
			return 1
		}

		if *conf.stream {
			b.StreamRandom(os.Stdout, *conf.streamDelay, *conf.ofs)
		} else {
			b.PrintRandom(os.Stdout, *conf.nChars, *conf.ofs)
		}

	// Print a RandomRune or stream from two or more blocks.
	case len(conf.blocks) > 1:
		bm := map[string]horeb.UnicodeBlock{}
		for _, b := range conf.blocks {
			val, ok := horeb.Blocks[b]
			if !ok {
				slog.Warn("Unknown block", "block", b)
				continue
			}
			bm[b] = val
		}
		if len(bm) > 0 {
			defer fmt.Println()
			if *conf.stream {
				ticker := time.NewTicker(*conf.streamDelay)
				for range ticker.C {

					block, err := horeb.RandomBlock(bm)
					if err != nil {
						slog.Error("Error getting random block", err)
						return 1
					}
					fmt.Printf("%c%s", block.RandomRune(), *conf.ofs)
				}
			} else {
				for i := 0; i < *conf.nChars; i++ {
					block, err := horeb.RandomBlock(bm)
					if err != nil {
						slog.Error("Error getting random block", err)
						return 1
					}
					fmt.Printf("%c%s", block.RandomRune(), *conf.ofs)
				}
			}
		}
	}

	return 0
}
