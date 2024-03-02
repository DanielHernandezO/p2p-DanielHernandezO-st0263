package restapi

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/DanielHernandezO/p2p/download/internal/infraestructure/config"
	"github.com/DanielHernandezO/p2p/download/internal/infraestructure/delivery"
	"github.com/gin-gonic/gin"
)

type restApi struct{}

func NewRestApi() delivery.Strategy {
	return &restApi{}
}

func (r *restApi) Start() {
	config.LoadConfigs()
	dependencies := buildDependencies()

	router := gin.Default()

	mapUrls(router, dependencies)

	go dependencies.downloadSuscriber.Suscribe()

	go func() {
		router.Run(fmt.Sprintf(":%s", config.GetStringPropetyBykey(config.Port)))
	}()

	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, syscall.SIGINT, syscall.SIGTERM)
	<-stopChan

	log.Println("Apagando el servidor...")
}
