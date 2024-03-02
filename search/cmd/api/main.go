package main

import "github.com/DanielHernandezO/p2p/search/internal/infrastructure/delivery/grpcapi"

func main() {
	deliveryStrategy := grpcapi.NewGrpcApi()
	deliveryStrategy.Start()
}
