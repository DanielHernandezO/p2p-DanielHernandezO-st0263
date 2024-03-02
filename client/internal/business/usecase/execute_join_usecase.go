package usecase

import (
	"context"

	"github.com/DanielHernandezO/p2p/client/internal/business/gateway"
)

type ExecuteJoinUsecase interface {
	ExecuteJoin(context context.Context) string
	Search(context context.Context, filename string) string
}

type executeJoinUsecase struct {
	serverService gateway.ServerServiceGateway
}

func NewExecuteJoin(serverService gateway.ServerServiceGateway) *executeJoinUsecase {
	return &executeJoinUsecase{
		serverService: serverService,
	}
}

func (e *executeJoinUsecase) ExecuteJoin(context context.Context) string {
	answer := e.serverService.ExecuteJoin(context)
	return answer
}

func (e *executeJoinUsecase) Search(context context.Context, filename string) string {
	answer := e.serverService.Search(context, filename)
	return answer
}
