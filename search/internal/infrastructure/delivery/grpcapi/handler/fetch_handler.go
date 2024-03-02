package handler

import (
	"context"

	"github.com/DanielHernandezO/p2p/search/internal/business/usecase"
	"github.com/DanielHernandezO/p2p/search/internal/infrastructure/config"
)

type fetchHanlder struct {
	fecthUsecase usecase.FetchUsecase
	config.UnimplementedFetchServer
}

func NewFetchHandler(fecthUsecase usecase.FetchUsecase) *fetchHanlder {
	return &fetchHanlder{
		fecthUsecase: fecthUsecase,
	}
}

func (f *fetchHanlder) ExecuteFetch(context context.Context, params *config.VoidRequest) (*config.Peers, error) {
	peers := f.fecthUsecase.ExecuteFetch(&context)
	return peers, nil
}

func (f *fetchHanlder) Fetch(context context.Context, params *config.VoidRequest) (*config.Peers, error) {
	peers := f.fecthUsecase.Fetch(&context)
	return peers, nil
}
