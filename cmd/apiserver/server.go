package main

import (
	server "github.com/BenWhiting/pod-spawner/internal/apiserver"
	"github.com/BenWhiting/pod-spawner/internal/apiserver/podrepository"
	"github.com/sirupsen/logrus"
)

func main() {
	// Logger
	log := logrus.New()

	//  Generate server config
	config, err := server.Generate()
	if err != nil {
		panic(err)
	}
	// Pod Provider
	pp := podrepository.New()

	// Server
	s, err := server.New(log, pp)
	if err != nil {
		panic(err)
	}

	s.Start(config)
}
