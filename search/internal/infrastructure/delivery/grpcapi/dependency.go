package grpcapi

import (
	"github.com/DanielHernandezO/p2p/search/internal/business/usecase"
	"github.com/DanielHernandezO/p2p/search/internal/infrastructure/config"
	"github.com/DanielHernandezO/p2p/search/internal/infrastructure/delivery/grpcapi/handler"
	"github.com/DanielHernandezO/p2p/search/internal/infrastructure/repository/provider"
	"github.com/DanielHernandezO/p2p/search/internal/infrastructure/repository/queue/publisher"
	"github.com/DanielHernandezO/p2p/search/internal/infrastructure/repository/storage"
)

type Dependencies struct {
	searchHandler config.SearchFileServer
	joinHandler   config.JoinServer
	fetchHandler  config.FetchServer
}

func buildDependencies() *Dependencies {

	//queue
	downloadQueue := publisher.NewDownloadQueue()
	downloadQueue.Start()

	//storage
	peerStorage := storage.NewPeerStorage()

	//providers
	peerProvider := provider.NewPeerProvider()
	fileManagmentProvider := provider.NewFileManagmentProvider()

	//usecases
	joinUsecase := usecase.NewJoinUsecase(peerProvider, peerStorage)
	searchUsecase := usecase.NewSearchUsecase(fileManagmentProvider, peerProvider, peerStorage, downloadQueue)
	fetchUsecase := usecase.NewFetchUsecase(peerProvider, peerStorage)

	return &Dependencies{
		searchHandler: handler.NewSearchHandler(searchUsecase),
		joinHandler:   handler.NewJoinHandler(joinUsecase),
		fetchHandler:  handler.NewFetchHandler(fetchUsecase),
	}
}
