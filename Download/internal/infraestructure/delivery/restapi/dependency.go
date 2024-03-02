package restapi

import (
	"github.com/DanielHernandezO/p2p/download/internal/business/usecase"
	"github.com/DanielHernandezO/p2p/download/internal/infraestructure/delivery/restapi/handler"
	"github.com/DanielHernandezO/p2p/download/internal/infraestructure/repository/provider"
	"github.com/DanielHernandezO/p2p/download/internal/infraestructure/repository/queue/suscriber"
)

type Dependencies struct {
	downloadSuscriber suscriber.DownloadSuscriber
	saveHandler       handler.SaveHandler
}

func buildDependencies() *Dependencies {

	//provider
	downloadProvider := provider.NewDownloadProvider()

	//usecase
	downloadUsecase := usecase.NewDownloadUsecase(downloadProvider)
	saveUsecase := usecase.NewSaveUsecase()

	//suscriber
	downloadSuscriber := suscriber.NewDownloadSuscriber(downloadUsecase)

	//handler
	saveHandler := handler.NewSaveHander(saveUsecase)

	return &Dependencies{
		downloadSuscriber: downloadSuscriber,
		saveHandler:       saveHandler,
	}
}
