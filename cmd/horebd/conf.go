package main

import (
	"fmt"
	"net"

	"github.com/qjcg/horeb/pkg/horeb"
	pb "github.com/qjcg/horeb/proto"
	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

// Conf represents this app's configuration.
type Conf struct {
	logger *logrus.Logger
	viper  *viper.Viper
}

// NewConf returns a pointer to an initialized Conf value.
func NewConf(fs *pflag.FlagSet) *Conf {
	conf := Conf{
		logger: logrus.New(),
		viper:  viper.New(),
	}

	if err := conf.viper.BindPFlags(fs); err != nil {
		conf.logger.Fatal(err)
	}

	conf.viper.SetEnvPrefix("HOREBD")
	conf.viper.AutomaticEnv()

	if conf.viper.GetBool("debug") {
		conf.logger.SetLevel(logrus.DebugLevel)
	}

	if conf.viper.GetBool("json") {
		conf.logger.SetFormatter(&logrus.JSONFormatter{})
	}

	conf.logger.WithFields(logrus.Fields{
		"ip":   conf.viper.GetString("ip"),
		"port": conf.viper.GetUint("port"),
	}).Infof("Listening on gRPC socket")

	return &conf
}

func (c *Conf) socketAddress() string {
	return fmt.Sprintf(
		"%s:%s",
		c.viper.GetString("ip"),
		c.viper.GetString("port"),
	)
}

func (c *Conf) listenAndServe() error {
	grpcServer := grpc.NewServer()
	pb.RegisterHorebServer(grpcServer, c)

	listener, err := net.Listen("tcp", c.socketAddress())
	if err != nil {
		return err
	}

	if err := grpcServer.Serve(listener); err != nil {
		return err
	}

	return nil
}

// GetStream implements the pb.HorebServer interface.
func (c *Conf) GetStream(r *pb.RuneRequest, stream pb.Horeb_GetStreamServer) error {
	block, ok := horeb.Blocks[r.Block]
	if !ok {
		c.logger.Errorf("Invalid block: %s\n", r.Block)
		return fmt.Errorf("Invalid block: %s", r.Block)
	}

	for i := 0; i < int(r.Num); i++ {
		myRandomRune := pb.Rune{R: string(block.RandomRune())}
		if err := stream.Send(&myRandomRune); err != nil {
			return err
		}
	}

	c.logger.WithFields(logrus.Fields{
		"block": r.Block,
		"num":   r.Num,
	}).Info("Sent response")

	return nil
}
