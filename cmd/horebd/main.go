package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/qjcg/horeb"
	pb "github.com/qjcg/horeb/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

// server is used to implement our horeb server.
type server struct{}

func (s server) GetStream(rr *pb.RuneRequest, stream pb.Horeb_GetStreamServer) error {
	grpclog.Printf("Got: %#v", rr)
	for i := 0; i < int(rr.Num); i++ {

		myRandomRune := pb.Rune{R: string(horeb.Blocks["geometric"].RandomRune())}
		if err := stream.Send(&myRandomRune); err != nil {
			return err
		}
		grpclog.Printf("Sent: %#v", myRandomRune)
	}
	return nil
}

func main() {
	ip := flag.String("i", "127.0.0.1", "ip address to listen on")
	port := flag.String("p", "9999", "TCP port to listen on")
	flag.Parse()

	listenString := fmt.Sprintf("%s:%s", *ip, *port)
	lis, err := net.Listen("tcp", listenString)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterHorebServer(s, &server{})

	grpclog.Printf("Horeb gRPC server listening on tcp://%s", listenString)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
