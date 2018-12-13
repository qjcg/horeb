package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/qjcg/horeb"
	pb "github.com/qjcg/horeb/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

var logger = grpclog.NewLoggerV2(os.Stderr, os.Stderr, os.Stderr)

func init() {
	grpclog.SetLoggerV2(logger)
}

func getRuneStream(client pb.HorebClient, rr *pb.RuneRequest) {
	logger.Infof("SENT: %#v\n", rr)

	stream, err := client.GetStream(context.Background(), rr)
	if err != nil {
		logger.Fatalf("%v.GetStream(_) = _, %v\n", client, err)
	}

	for {
		streamedRune, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			logger.Errorf("stream receive error: %v\n", err)
		}
		fmt.Println(streamedRune.R)
	}
	logger.Infof("RECEIVED: %d %s", rr.Num, rr.Block)
}

func main() {
	block := flag.String("b", "geometric", "unicode block name")
	ip := flag.String("i", "127.0.0.1", "ip address of horebd server")
	port := flag.Int("p", 9999, "TCP port of horebd server")
	num := flag.Int("n", 10, "number of runes to request")
	version := flag.Bool("v", false, "print version")
	flag.Parse()

	if *version {
		fmt.Println(horeb.Version)
		return
	}

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", *ip, *port), opts...)
	if err != nil {
		logger.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewHorebClient(conn)

	getRuneStream(
		client,
		&pb.RuneRequest{Num: int32(*num), Block: *block},
	)
}
