package main

import (
	"fmt"
	"net"
	"os"

	"github.com/qjcg/horeb/pkg/horeb"
	pb "github.com/qjcg/horeb/proto"
	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

var logger = grpclog.NewLoggerV2(os.Stderr, os.Stderr, os.Stderr)

// Conf represents this app's configuration.
type Conf struct {
	logger *logrus.Logger
	viper  *viper.Viper
}

func main() {
	pflag.StringP("ip", "i", "0.0.0.0", "ip address to listen on")
	pflag.UintP("port", "p", 9999, "TCP port to listen on")
	pflag.BoolP("version", "v", false, "print version")
	pflag.Parse()

	viper.SetEnvPrefix("HOREBD")
	viper.AutomaticEnv()

	conf := Conf{
		logger: logrus.New(),
		viper:  viper.GetViper(),
	}

	if viper.GetBool("version") {
		fmt.Println(horeb.Version)
		return
	}

	var lis net.Listener
	listenFlags := fmt.Sprintf("%s:%v", viper.GetString("ip"), viper.GetString("port"))
	lis, err := net.Listen("tcp", listenFlags)
	if err != nil {
		conf.logger.Fatalf("Failed to listen: %v", err)
	}
	conf.logger.Infof("Using flag listener: %#v\n", lis)

	s := grpc.NewServer()
	pb.RegisterHorebServer(s, &server{})

	conf.logger.Infof("Horeb gRPC server listening on %#v", lis)

	if err := s.Serve(lis); err != nil {
		conf.logger.Fatalf("failed to serve: %#v", err)
	}
}

// server is used to implement our horeb server.
type server struct{}

func (s server) GetStream(conf *Conf, rr *pb.RuneRequest, stream pb.Horeb_GetStreamServer) error {
	conf.logger.Info("RECEIVED: %#v", rr)
	block, ok := horeb.Blocks[rr.Block]
	if !ok {
		conf.logger.Errorf("Invalid block: %s\n", rr.Block)
		return fmt.Errorf("Invalid block: %s", rr.Block)
	}
	for i := 0; i < int(rr.Num); i++ {
		myRandomRune := pb.Rune{R: string(block.RandomRune())}
		if err := stream.Send(&myRandomRune); err != nil {
			return err
		}
	}
	conf.logger.Infof("SENT: %d %s\n", rr.Num, rr.Block)
	return nil
}
