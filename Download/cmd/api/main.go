package main

import "github.com/DanielHernandezO/p2p/download/internal/infraestructure/delivery/restapi"

func main() {
	deliveryStrategy := restapi.NewRestApi()
	deliveryStrategy.Start()
}
