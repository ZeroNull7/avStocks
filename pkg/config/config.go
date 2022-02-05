package config

import (
	"fmt"

	"github.com/spf13/viper"
)

var config *Config

type Config struct {
	Intraday struct {
		Extended struct {
			Get struct {
				Stocks []string `yaml:stocks`
			} `yaml:"get"`
		} `yaml:"extended"`
	} `yaml:"intraday"`
}

// Read the config file from the current directory and marshal
// into the conf config struct.
func get() *Config {

	if config == nil {
		err := viper.ReadInConfig()

		if err != nil {
			fmt.Printf("%v", err)
		}

		config = &Config{}
		err = viper.Unmarshal(config)
		if err != nil {
			fmt.Printf("unable to decode into config struct, %v", err)
		}
	}

	return config
}
