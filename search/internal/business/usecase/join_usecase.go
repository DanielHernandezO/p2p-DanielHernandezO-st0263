package usecase

import (
	"context"
	"fmt"
	"log"

	"github.com/DanielHernandezO/p2p/search/internal/business/constant"
	"github.com/DanielHernandezO/p2p/search/internal/business/domain"
	"github.com/DanielHernandezO/p2p/search/internal/business/gateway"
	"github.com/DanielHernandezO/p2p/search/internal/infrastructure/config"
)

type JoinUsecase interface {
	ExecuteJoin(context *context.Context, socket *domain.Socket, peers map[string]string) (*config.Response, error)
	Join(context *context.Context, socket *domain.Socket) (*config.Response, error)
}

type joinUsecase struct {
	peerProvider gateway.PeerServiceGateway
	peerStorage  gateway.PeerStorageGateway
}

func NewJoinUsecase(peerProvider gateway.PeerServiceGateway, peerStorage gateway.PeerStorageGateway) *joinUsecase {
	return &joinUsecase{
		peerProvider: peerProvider,
		peerStorage:  peerStorage,
	}
}

func (j *joinUsecase) ExecuteJoin(context *context.Context, socket *domain.Socket, peers map[string]string) (*config.Response, error) {
	j.peerStorage.AddPeers(peers)
	peers = j.peerStorage.GetPeers()
	response := j.peerProvider.ExecuteJoin(context, socket, peers)
	if response {
		log.Printf(constant.Connected, socket.Ip)
		return &config.Response{
			Message: constant.Connected,
		}, nil
	}
	log.Print(constant.ConnectingPeersError)
	return nil, fmt.Errorf(constant.ConnectingPeersError)
}

func (j *joinUsecase) Join(context *context.Context, socket *domain.Socket) (*config.Response, error) {
	j.peerStorage.AddPeer(socket)
	log.Printf(constant.Connected, socket.Ip)
	return &config.Response{
		Message: constant.Connected,
	}, nil
}
