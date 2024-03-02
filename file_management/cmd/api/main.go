package main

import "github.com/DanielHernandezO/p2p/file_management/internal/infrastructure/delivery/restapi"

func main() {
	deliveryStrategy := restapi.NewRestApi()
	deliveryStrategy.Start()
}
