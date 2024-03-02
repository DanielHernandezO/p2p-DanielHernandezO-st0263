package restapi

import (
	"github.com/DanielHernandezO/p2p/file_management/internal/business/usecase"
	"github.com/DanielHernandezO/p2p/file_management/internal/infrastructure/delivery/restapi/handler"
)

type Dependencies struct {
	fileHandler handler.FileHandler
}

func buildDependencies() *Dependencies {
	fileUsecase := usecase.NewFileUsecase()
	return &Dependencies{
		fileHandler: handler.NewFileHandler(fileUsecase),
	}
}
