package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/qjcg/horeb"
	pb "github.com/qjcg/horeb/proto"
	"github.com/spf13/viper"

	"google.golang.org/grpc"
)

// Setup for global levelled logging.
var (
	info  = log.New(os.Stderr, "INFO ", log.LstdFlags)
	debug = log.New(ioutil.Discard, "DEBUG ", log.LstdFlags)
)

func main() {
	debugFlag := flag.Bool("d", false, "debug logging")
	version := flag.Bool("v", false, "print version")
	flag.Parse()

	conf := viper.New()
	conf.SetDefault("block", "geometric")
	conf.SetDefault("host", "localhost")
	conf.SetDefault("num", 5)
	conf.SetDefault("port", 9999)

	conf.SetEnvPrefix("HOREBCTL")
	conf.BindEnv("block")
	conf.BindEnv("host")
	conf.BindEnv("num")
	conf.BindEnv("port")

	if *version {
		fmt.Println(horeb.Version)
		return
	}

	if *debugFlag {
		debug.SetOutput(os.Stderr)
	}

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", conf.GetString("host"), conf.GetInt("port")), opts...)
	if err != nil {
		info.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewHorebClient(conn)

	err = getRuneStream(
		client,
		&pb.RuneRequest{
			Num:   int32(conf.GetInt("num")),
			Block: conf.GetString("block"),
		},
	)
	if err != nil {
		info.Fatal(err)
	}
}

func getRuneStream(client pb.HorebClient, rr *pb.RuneRequest) error {
	debug.Printf("SENT: %#v\n", rr)

	stream, err := client.GetStream(context.Background(), rr)
	if err != nil {
		info.Printf("%v.GetStream(_) = _, %v\n", client, err)
		return err
	}

	for {
		streamedRune, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			info.Printf("stream receive error: %v\n", err)
		}
		fmt.Println(streamedRune.R)
	}
	debug.Printf("RECEIVED: %d %s", rr.Num, rr.Block)

	return nil
}
