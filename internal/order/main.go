package main

import (
	"github.com/spf13/viper"
	"github.com/zandala88/gorder/common/config"
	"log"
)

func init() {
	if err := config.NewViperConfig(); err != nil {
		log.Fatalf("could not read config: %v", err)
	}
}

func main() {
	log.Printf("config: %v", viper.Get("order"))
}
