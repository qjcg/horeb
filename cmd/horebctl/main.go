package main

import (
	"context"
	"flag"
	"fmt"
	"io"

	pb "github.com/qjcg/horeb/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

func getRuneStream(client pb.HorebClient, rr *pb.RuneRequest) {
	grpclog.Printf("Sent: %#v", rr)

	stream, err := client.GetStream(context.Background(), rr)
	if err != nil {
		grpclog.Fatalf("%v.GetStream(_) = _, %v", client, err)
	}

	for {
		streamedRune, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			grpclog.Fatal("stream receive error: %v", err)
		}
		grpclog.Printf("Got: %#v", streamedRune)
	}
}

func main() {
	ip := flag.String("i", "127.0.0.1", "ip address of horebd server")
	port := flag.Int("p", 9999, "TCP port of horebd server")
	num := flag.Int("n", 10, "number of runes to request")
	flag.Parse()

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", *ip, *port), opts...)
	if err != nil {
		grpclog.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewHorebClient(conn)

	getRuneStream(
		client,
		&pb.RuneRequest{Num: int32(*num)},
	)
}
