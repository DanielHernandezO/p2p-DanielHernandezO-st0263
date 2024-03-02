package restapi

import (
	"fmt"

	"github.com/DanielHernandezO/p2p/file_management/internal/infrastructure/config"
	"github.com/DanielHernandezO/p2p/file_management/internal/infrastructure/delivery"
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

	router.Run(fmt.Sprintf(":%s", config.GetStringPropetyBykey(config.Port)))
}
