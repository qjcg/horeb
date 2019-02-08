package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/coreos/go-systemd/activation"
	"github.com/qjcg/horeb"
	pb "github.com/qjcg/horeb/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

var logger = grpclog.NewLoggerV2(os.Stderr, os.Stderr, os.Stderr)

func main() {
	ip := flag.String("i", "0.0.0.0", "ip address to listen on")
	port := flag.String("p", "9999", "TCP port to listen on")
	version := flag.Bool("v", false, "print version")
	flag.Parse()

	if *version {
		fmt.Println(horeb.Version)
		return
	}

	var lis net.Listener
	listenFlags := fmt.Sprintf("%s:%s", *ip, *port)

	// Use systemd socket listener if available, otherwise fall back to
	// listenFlags parameters.
	listeners, err := activation.Listeners()
	if err != nil {
		log.Fatalf("Couldn't get listeners: %s\n", err)
	}

	logger.Infof("Systemd listeners: %#v", listeners)

	if len(listeners) == 1 {
		lis = listeners[0]
		logger.Infof("Using systemd listener: %#v\n", lis)
	} else {
		lis, err = net.Listen("tcp", listenFlags)
		if err != nil {
			logger.Fatalf("Failed to listen: %v", err)
		}
		logger.Infof("Using flag listener: %#v\n", lis)
	}

	s := grpc.NewServer()
	pb.RegisterHorebServer(s, &server{})

	logger.Infof("Horeb gRPC server listening on %#v", lis)

	if err := s.Serve(lis); err != nil {
		logger.Fatalf("failed to serve: %#v", err)
	}
}

// server is used to implement our horeb server.
type server struct{}

func (s server) GetStream(rr *pb.RuneRequest, stream pb.Horeb_GetStreamServer) error {
	logger.Infof("RECEIVED: %#v", rr)
	block, ok := horeb.Blocks[rr.Block]
	if !ok {
		logger.Errorf("Invalid block: %s\n", rr.Block)
		return fmt.Errorf("Invalid block: %s", rr.Block)
	}
	for i := 0; i < int(rr.Num); i++ {
		myRandomRune := pb.Rune{R: string(block.RandomRune())}
		if err := stream.Send(&myRandomRune); err != nil {
			return err
		}
	}
	logger.Infof("SENT: %d %s\n", rr.Num, rr.Block)
	return nil
}
