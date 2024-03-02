package usecase

import (
	"context"
	"fmt"

	"github.com/DanielHernandezO/p2p/search/internal/business/constant"
	"github.com/DanielHernandezO/p2p/search/internal/business/domain"
	"github.com/DanielHernandezO/p2p/search/internal/business/gateway"
	"github.com/DanielHernandezO/p2p/search/internal/infrastructure/config"
)

var (
	proccesedRequest map[string]bool
	host             = "host"
)

type SearchUsecase interface {
	Search(context *context.Context, fileRequested *domain.FileRequested) (*config.Response, error)
}

type searchUsecase struct {
	fileManagment gateway.FileManagmentProvider
	peerProvider  gateway.PeerServiceGateway
	peerStorage   gateway.PeerStorageGateway
	downloadQueue gateway.DownloadQueue
}

func NewSearchUsecase(fileManagment gateway.FileManagmentProvider, peerProvider gateway.PeerServiceGateway, peerStorage gateway.PeerStorageGateway, downloadQueue gateway.DownloadQueue) *searchUsecase {
	proccesedRequest = make(map[string]bool)
	return &searchUsecase{
		fileManagment: fileManagment,
		peerProvider:  peerProvider,
		peerStorage:   peerStorage,
		downloadQueue: downloadQueue,
	}
}

func (s *searchUsecase) Search(context *context.Context, fileRequested *domain.FileRequested) (*config.Response, error) {
	if proccesedRequest[fileRequested.Id] {
		return nil, fmt.Errorf(constant.AlreadyProccessed)
	}
	proccesedRequest[fileRequested.Id] = true

	response := s.fileManagment.GetFile(fileRequested.Name)
	peers := s.peerStorage.GetPeers()
	found := s.peerProvider.SearchFile(context, fileRequested, peers)
	if response == nil && !found {
		return nil, fmt.Errorf(constant.NotFound)
	}

	if response != nil {
		iphost := config.GetStringPropetyBykey(fmt.Sprintf(config.IpConfig, host))
		messageQueue := domain.QueueMessage{
			Name:    iphost + "-" + fileRequested.Name,
			Ip:      fileRequested.Ip,
			Port:    fileRequested.Port,
			Content: response.Bytes,
		}
		err := s.downloadQueue.AddMessage(messageQueue)
		if err != nil {
			return nil, err
		}
	}

	return &config.Response{
		Message: constant.Found,
	}, nil
}
