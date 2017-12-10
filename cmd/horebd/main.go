package main

import (
	"log"
	"net"

	"github.com/qjcg/horeb"
	pb "github.com/qjcg/horeb/proto"

	"google.golang.org/grpc"
)

// server is used to implement our horeb server.
type server struct{}

func (s server) GetStream(rr *pb.RuneRequest, stream pb.Horeb_GetStreamServer) error {
	for i := 0; i < int(rr.Num); i++ {

		myRandomRune := horeb.Blocks["geometric"].RandomRune()
		if err := stream.Send(&pb.Rune{string(myRandomRune)}); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	lis, err := net.Listen("tcp", ":9999")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterHorebServer(s, &server{})

	log.Println("Horeb gRPC server listening on tcp://0.0.0.0:9999")

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
