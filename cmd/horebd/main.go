package main

import (
	"flag"
	"fmt"
	"net"
	"os"

	"github.com/qjcg/horeb"
	pb "github.com/qjcg/horeb/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

var logger = grpclog.NewLoggerV2(os.Stderr, os.Stderr, os.Stderr)

// server is used to implement our horeb server.
type server struct{}

func (s server) GetStream(rr *pb.RuneRequest, stream pb.Horeb_GetStreamServer) error {
	logger.Infof("Got: %#v", rr)
	for i := 0; i < int(rr.Num); i++ {

		myRandomRune := pb.Rune{R: string(horeb.Blocks["geometric"].RandomRune())}
		if err := stream.Send(&myRandomRune); err != nil {
			return err
		}
		logger.Infof("Sent: %#v", myRandomRune)
	}
	return nil
}

func main() {
	ip := flag.String("i", "127.0.0.1", "ip address to listen on")
	port := flag.String("p", "9999", "TCP port to listen on")
	version := flag.Bool("v", false, "print version")
	flag.Parse()

	if *version {
		fmt.Println(horeb.Version)
		return
	}

	listenString := fmt.Sprintf("%s:%s", *ip, *port)
	lis, err := net.Listen("tcp", listenString)
	if err != nil {
		logger.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterHorebServer(s, &server{})

	logger.Infof("Horeb gRPC server listening on tcp://%s", listenString)

	if err := s.Serve(lis); err != nil {
		logger.Fatalf("failed to serve: %v", err)
	}
}
