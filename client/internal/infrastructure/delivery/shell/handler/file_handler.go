package handler

import (
	"context"

	"github.com/DanielHernandezO/p2p/client/internal/business/usecase"
)

type FileHandler interface {
	PostFiles() string
}

type fileHandler struct {
	fileUsecase usecase.FileUsecase
}

func NewFileHandler(fileUsecase usecase.FileUsecase) *fileHandler {
	return &fileHandler{
		fileUsecase: fileUsecase,
	}
}

func (f *fileHandler) PostFiles() string {
	context := context.Background()
	answer := f.fileUsecase.PostFiles(context)
	return answer
}
