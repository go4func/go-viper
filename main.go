package main

import (
	"fmt"

	"github.com/spf13/viper"
)

type Configuration struct {
	Server   ServerConfiguration
	Database DatabaseConfiguration
}
type DatabaseConfiguration struct {
	ConnectionUri string
}
type ServerConfiguration struct {
	Port int
}

type test struct {
	Config config
}

type config struct {
	Uri    string
	Worker int
}

func readConfig(cfName string) error {
	var cf test
	viper.AddConfigPath(".")
	viper.SetConfigName(cfName)

	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	err = viper.Unmarshal(&cf)
	if err != nil {
		return err
	}
	fmt.Println(viper.ConfigFileUsed(), cf)

	return err
}

func main() {

	//readConfig("config")
	readConfig("test.dev")
}
