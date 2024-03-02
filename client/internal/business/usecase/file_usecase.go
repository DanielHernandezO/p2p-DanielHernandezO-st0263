package usecase

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/DanielHernandezO/p2p/client/internal/business/constant"
	"github.com/DanielHernandezO/p2p/client/internal/business/gateway"
	"github.com/DanielHernandezO/p2p/client/internal/infrastructure/config"
)

type FileUsecase interface {
	PostFiles(context context.Context) string
}

type fileUsecase struct {
	location   string
	fileServer gateway.FileServiceGateway
}

func NewFileUsecase(fileServer gateway.FileServiceGateway) *fileUsecase {
	return &fileUsecase{
		location:   config.GetStringPropetyBykey(config.FilesLocation),
		fileServer: fileServer,
	}
}

func (f *fileUsecase) PostFiles(context context.Context) string {
	err := filepath.Walk(f.location, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			answer := f.fileServer.PostFile(context, path)
			if answer != constant.PostedFile {
				return fmt.Errorf(answer)
			}
			return nil
		}

		return nil
	})

	if err != nil {
		return err.Error()
	}

	return constant.PostedFile
}
