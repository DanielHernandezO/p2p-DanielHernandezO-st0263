package handler

import (
	"context"

	"github.com/DanielHernandezO/p2p/search/internal/business/domain"
	"github.com/DanielHernandezO/p2p/search/internal/business/usecase"
	"github.com/DanielHernandezO/p2p/search/internal/infrastructure/config"
)

type searchHandler struct {
	searchUsecase usecase.SearchUsecase
	config.UnimplementedSearchFileServer
}

func NewSearchHandler(searchUsecase usecase.SearchUsecase) *searchHandler {
	return &searchHandler{
		searchUsecase: searchUsecase,
	}
}

func (s *searchHandler) Search(context context.Context, bodyRequest *config.FileRequested) (*config.Response, error) {
	body := &domain.FileRequested{
		Id:   bodyRequest.Id,
		Name: bodyRequest.Name,
		Ip:   bodyRequest.IpResponse,
		Port: bodyRequest.PortResponse,
	}
	response, err := s.searchUsecase.Search(&context, body)
	return response, err
}
