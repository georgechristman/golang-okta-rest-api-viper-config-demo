package main

import (
	"log"

	"github.com/georgechristman/golang-okta-rest-api-viper-config-demo/api"
	"github.com/georgechristman/golang-okta-rest-api-viper-config-demo/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	server, err := api.NewServer(config)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
