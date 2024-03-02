package shell

import (
	"github.com/DanielHernandezO/p2p/client/internal/business/usecase"
	"github.com/DanielHernandezO/p2p/client/internal/infrastructure/delivery/shell/handler"
	"github.com/DanielHernandezO/p2p/client/internal/infrastructure/repository/provider"
)

type Dependencies struct {
	executeJoinHandler handler.ExecuteJoinHandler
	fileHandler        handler.FileHandler
}

func buildDependencies() *Dependencies {

	serverProvider := provider.NewServerProvider()
	fileProvider := provider.NewFileProvider()

	executeJoinUsecase := usecase.NewExecuteJoin(serverProvider)
	fileUsecase := usecase.NewFileUsecase(fileProvider)

	executeJoinHandler := handler.NewExecuteJoin(executeJoinUsecase)
	fileHandler := handler.NewFileHandler(fileUsecase)
	return &Dependencies{
		executeJoinHandler: executeJoinHandler,
		fileHandler:        fileHandler,
	}
}
