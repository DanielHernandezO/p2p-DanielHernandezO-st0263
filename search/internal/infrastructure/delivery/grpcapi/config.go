package grpcapi

import (
	"fmt"
	"log"
	"net"

	"github.com/DanielHernandezO/p2p/search/internal/infrastructure/config"
	"github.com/DanielHernandezO/p2p/search/internal/infrastructure/delivery"
	"google.golang.org/grpc"
)

var host = "host"

type grpcapi struct{}

func NewGrpcApi() delivery.Strategy {
	return &grpcapi{}
}

func (r *grpcapi) Start() {
	config.LoadConfigs()
	dependencies := buildDependencies()

	lis, err := net.Listen(config.GetStringPropetyBykey(config.Network), fmt.Sprintf(":%s", config.GetStringPropetyBykey(fmt.Sprintf(config.PortConfig, host))))
	if err != nil {
		log.Fatalf("cannot create listener: %s", err)
	}

	serverRegister := grpc.NewServer()
	config.RegisterSearchFileServer(serverRegister, dependencies.searchHandler)
	config.RegisterJoinServer(serverRegister, dependencies.joinHandler)
	config.RegisterFetchServer(serverRegister, dependencies.fetchHandler)
	log.Printf("gRPC server listening on :%s", config.GetStringPropetyBykey(fmt.Sprintf(config.PortConfig, host)))
	err = serverRegister.Serve(lis)
	if err != nil {
		log.Fatalf("impossible to serve: %s", err)
	}
}
