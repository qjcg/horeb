package main

import (
	"context"
	"fmt"
	"io"

	"github.com/qjcg/horeb/pkg/horeb"
	pb "github.com/qjcg/horeb/proto"

	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
	"google.golang.org/grpc"
)

func main() {
	pflag.StringP("block", "b", "geometric", "unicode block")
	pflag.BoolP("debug", "d", false, "debug logging")
	pflag.StringP("host", "i", "localhost", "host to connect to")
	pflag.BoolP("json", "j", false, "JSON-formatted logging")
	pflag.UintP("number", "n", 5, "number of runes to print")
	pflag.UintP("port", "p", 9999, "TCP port to connect to")
	pflag.StringP("separator", "s", " ", "rune separator")
	pflag.BoolP("version", "v", false, "print version")
	pflag.Parse()

	conf := NewConf(pflag.CommandLine)

	if conf.viper.GetBool("version") {
		fmt.Println(horeb.Version)
		return
	}

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", conf.viper.GetString("host"), conf.viper.GetUint("port")), opts...)
	if err != nil {
		conf.logger.Fatalf("Failed to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewHorebClient(conn)
	request := &pb.RuneRequest{
		Num:   int32(conf.viper.GetUint("number")),
		Block: conf.viper.GetString("block"),
	}

	if err := conf.getRuneStream(client, request); err != nil {
		conf.logger.Fatal(err)
	}
}

func (conf *Conf) getRuneStream(client pb.HorebClient, r *pb.RuneRequest) error {
	conf.logger.WithFields(logrus.Fields{
		"block": r.Block,
		"num":   r.Num,
	}).Debugf("Requested")

	stream, err := client.GetStream(context.Background(), r)
	if err != nil {
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

	conf.logger.WithFields(logrus.Fields{
		"block": r.Block,
		"num":   r.Num,
	}).Debugf("Received")

	return nil
}
