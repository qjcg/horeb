package main

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// Conf stores this application's configuration and dependencies.
type Conf struct {
	logger *logrus.Logger
	viper  *viper.Viper
}

// NewConf initializes a new Conf value.
func NewConf(f *pflag.FlagSet) *Conf {
	conf := Conf{
		viper:  viper.New(),
		logger: logrus.New(),
	}

	if err := conf.viper.BindPFlags(f); err != nil {
		conf.logger.Fatal(err)
	}
	conf.viper.SetEnvPrefix("HOREBCTL")
	conf.viper.AutomaticEnv()

	if conf.viper.GetBool("debug") {
		conf.logger.SetLevel(logrus.DebugLevel)
	}

	if conf.viper.GetBool("json") {
		conf.logger.SetFormatter(&logrus.JSONFormatter{})
	}

	return &conf
}
