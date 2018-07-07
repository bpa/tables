package main

import (
	"log"

	"github.com/bpa/tables/data"
	"github.com/bpa/tables/notify"
	"github.com/bpa/tables/server"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func main() {
	pflag.IntP("port", "p", 3000, "help message for flagname")
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config: %s", err)
	}
	data.Initialize()
	notify.Initialize()
	server.Listen(viper.GetInt("port"))
}
