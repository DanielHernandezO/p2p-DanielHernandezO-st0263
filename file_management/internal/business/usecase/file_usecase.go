package usecase

import (
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"sync"

	"github.com/DanielHernandezO/p2p/file_management/internal/business/constant"
	"github.com/DanielHernandezO/p2p/file_management/internal/business/domain"
	"github.com/DanielHernandezO/p2p/file_management/internal/infrastructure/config"
	"github.com/gin-gonic/gin"
)

type FileUsecase interface {
	Post(c *gin.Context, file *multipart.FileHeader) *domain.Response
	GetFile(c *gin.Context, fileId string) (*domain.FilesResponse, *domain.Response)
}

type fileUsecase struct {
	location string
	mu       sync.Mutex
}

func NewFileUsecase() *fileUsecase {
	return &fileUsecase{
		location: config.GetStringPropetyBykey(config.Location),
	}
}

func (f *fileUsecase) Post(c *gin.Context, file *multipart.FileHeader) *domain.Response {
	f.mu.Lock()
	defer f.mu.Unlock()

	err := c.SaveUploadedFile(file, f.location+file.Filename)
	if err != nil {
		return &domain.Response{
			Message: constant.ErrorSaving,
			Code:    http.StatusInternalServerError,
		}
	}

	log.Print(constant.Added)
	return &domain.Response{
		Message: constant.Added,
		Code:    http.StatusOK,
	}
}

func (f *fileUsecase) GetFile(c *gin.Context, fileId string) (*domain.FilesResponse, *domain.Response) {
	f.mu.Lock()
	defer f.mu.Unlock()

	filePath := f.location + fileId
	_, err := os.Stat(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, &domain.Response{
				Code:    http.StatusNotFound,
				Message: constant.Notfound,
			}
		}
		return nil, &domain.Response{
			Code:    http.StatusInternalServerError,
			Message: constant.ServerError,
		}
	}

	return &domain.FilesResponse{
		FilePath: filePath,
	}, nil
}
