package handler

import (
	"context"

	"github.com/DanielHernandezO/p2p/client/internal/business/usecase"
)

type ExecuteJoinHandler interface {
	ExecuteJoin() string
	Search(filename string) string
}

type executeJoinHandler struct {
	executeJoinUsecase usecase.ExecuteJoinUsecase
}

func NewExecuteJoin(executeJoinUsecase usecase.ExecuteJoinUsecase) *executeJoinHandler {
	return &executeJoinHandler{
		executeJoinUsecase: executeJoinUsecase,
	}
}

func (e *executeJoinHandler) ExecuteJoin() string {
	context := context.Background()
	answer := e.executeJoinUsecase.ExecuteJoin(context)
	return answer
}

func (e *executeJoinHandler) Search(filename string) string {
	context := context.Background()
	answer := e.executeJoinUsecase.Search(context, filename)
	return answer
}
