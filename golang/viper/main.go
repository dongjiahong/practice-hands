package main

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Name    string
	Version string
	Sms     struct {
		Apple struct {
			Name string
		}
		Orange struct {
			Name string
		}
	}
}

var c *Config

func main() {
	v := viper.New()
	v.SetConfigName("app")
	v.AddConfigPath("conf")
	v.SetConfigType("toml")
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := v.Unmarshal(&c); err != nil {
		panic(err)
	}
	fmt.Printf("%+v", c)
}
