package usecase

import (
	"context"

	"github.com/DanielHernandezO/p2p/search/internal/business/gateway"
	"github.com/DanielHernandezO/p2p/search/internal/infrastructure/config"
)

type FetchUsecase interface {
	ExecuteFetch(context *context.Context) *config.Peers
	Fetch(*context.Context) *config.Peers
}

type fetchUsecase struct {
	peerProvider gateway.PeerServiceGateway
	peerStorage  gateway.PeerStorageGateway
}

func NewFetchUsecase(peerProvider gateway.PeerServiceGateway, peerStorage gateway.PeerStorageGateway) *fetchUsecase {
	return &fetchUsecase{
		peerProvider: peerProvider,
		peerStorage:  peerStorage,
	}
}

func (f *fetchUsecase) ExecuteFetch(context *context.Context) *config.Peers {
	peers := f.peerStorage.GetPeers()
	peers = f.peerProvider.ExecuteFetch(context, peers)
	f.peerStorage.AddPeers(peers)
	return &config.Peers{
		Peers: peers,
	}
}

func (f *fetchUsecase) Fetch(*context.Context) *config.Peers {
	peers := f.peerStorage.GetPeers()
	return &config.Peers{
		Peers: peers,
	}
}
