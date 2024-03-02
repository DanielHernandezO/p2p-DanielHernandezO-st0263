package handler

import (
	"context"

	"github.com/DanielHernandezO/p2p/search/internal/business/domain"
	"github.com/DanielHernandezO/p2p/search/internal/business/usecase"
	"github.com/DanielHernandezO/p2p/search/internal/infrastructure/config"
)

type joinHandler struct {
	joinUsecase usecase.JoinUsecase
	config.UnimplementedJoinServer
}

func NewJoinHandler(joinUsecase usecase.JoinUsecase) *joinHandler {
	return &joinHandler{
		joinUsecase: joinUsecase,
	}
}

func (j *joinHandler) ExecuteJoin(context context.Context, joinPeersRequest *config.JoinPeersRequest) (*config.Response, error) {
	socket := &domain.Socket{
		Ip:   joinPeersRequest.Ip,
		Port: joinPeersRequest.Port,
	}
	peers := joinPeersRequest.Peers
	response, err := j.joinUsecase.ExecuteJoin(&context, socket, peers)
	return response, err
}

func (j *joinHandler) Join(context context.Context, joinRequest *config.JoinRequest) (*config.Response, error) {
	socket := &domain.Socket{
		Ip:   joinRequest.Ip,
		Port: joinRequest.Port,
	}
	response, err := j.joinUsecase.Join(&context, socket)
	return response, err
}
