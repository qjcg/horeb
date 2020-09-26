package main

import (
	"context"
	"fmt"
	"io"

	"github.com/qjcg/horeb/pkg/horeb"
	pb "github.com/qjcg/horeb/proto"

	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

// Conf stores this application's configuration and dependenciees.
type Conf struct {
	logger *logrus.Logger
	viper  *viper.Viper
}

// NewConf initializes a new Conf value.
func NewConf(v *viper.Viper) *Conf {
	conf := Conf{
		viper:  v,
		logger: logrus.New(),
	}

	return &conf
}

func main() {
	pflag.StringP("block", "b", "geometric", "unicode block")
	pflag.BoolP("debug", "d", false, "debug logging")
	pflag.StringP("host", "i", "localhost", "host to listen on")
	pflag.UintP("number", "n", 5, "number of runes to print")
	pflag.UintP("port", "p", 9999, "TCP port to listen on")
	pflag.StringP("separator", "s", " ", "rune separator")
	pflag.BoolP("version", "v", false, "print version")
	pflag.Parse()

	conf := NewConf(viper.New())

	conf.viper.SetEnvPrefix("HOREBCTL")
	conf.viper.AutomaticEnv()

	if conf.viper.GetBool("version") {
		conf.logger.Infoln(horeb.Version)
		return
	}

	if conf.viper.GetBool("debug") {
		conf.logger.SetLevel(logrus.DebugLevel)
	}

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", conf.viper.GetString("host"), conf.viper.GetUint("port")), opts...)
	if err != nil {
		conf.logger.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewHorebClient(conn)

	err = conf.getRuneStream(
		client,
		&pb.RuneRequest{
			Num:   int32(conf.viper.GetInt("number")),
			Block: conf.viper.GetString("block"),
		},
	)
	if err != nil {
		conf.logger.Fatalln(err)
	}
}

func (conf *Conf) getRuneStream(client pb.HorebClient, rr *pb.RuneRequest) error {
	conf.logger.Debugf("SENT: %#v\n", rr)

	stream, err := client.GetStream(context.Background(), rr)
	if err != nil {
		conf.logger.Infof("%v.GetStream(_) = _, %v\n", client, err)
		return err
	}

	for {
		streamedRune, err := stream.Recv()
		if err == io.EOF {
			// Add a newline at the end of the stream before breaking.
			fmt.Println()
			break
		}
		if err != nil {
			conf.logger.Infof("stream receive error: %v\n", err)
		}
		fmt.Printf("%s%s", streamedRune.R, conf.viper.GetString("separator"))
	}
	conf.logger.Debugf("RECEIVED: %d %s", rr.Num, rr.Block)

	return nil
}
