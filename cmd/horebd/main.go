package main

import (
	"fmt"

	"github.com/qjcg/horeb/pkg/horeb"
	"github.com/spf13/pflag"
)

func main() {
	pflag.BoolP("debug", "d", false, "debug-level logging")
	pflag.StringP("ip", "i", "0.0.0.0", "IP address to listen on")
	pflag.BoolP("json", "j", false, "JSON-formatted logging")
	pflag.UintP("port", "p", 9999, "TCP port to listen on")
	pflag.BoolP("version", "v", false, "print version")
	pflag.Parse()

	conf := NewConf(pflag.CommandLine)

	if conf.viper.GetBool("version") {
		fmt.Println(horeb.Version)
		return
	}

	if err := conf.listenAndServe(); err != nil {
		conf.logger.Fatal(err)
	}
}
