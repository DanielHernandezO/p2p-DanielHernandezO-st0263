package handler

import (
	"net/http"

	"github.com/DanielHernandezO/p2p/file_management/internal/business/constant"
	"github.com/DanielHernandezO/p2p/file_management/internal/business/domain"
	"github.com/DanielHernandezO/p2p/file_management/internal/business/usecase"
	"github.com/gin-gonic/gin"
)

type FileHandler interface {
	Post(c *gin.Context)
	GetFile(c *gin.Context)
}

type fileHandler struct {
	fileUsecase usecase.FileUsecase
}

func NewFileHandler(fileUsecase usecase.FileUsecase) *fileHandler {
	return &fileHandler{
		fileUsecase: fileUsecase,
	}
}

func (f *fileHandler) Post(c *gin.Context) {
	file, err := c.FormFile(constant.File)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.Response{
			Message: constant.BadRequest,
			Code:    http.StatusBadGateway,
		})
		return
	}

	response := f.fileUsecase.Post(c, file)

	c.JSON(response.Code, domain.Response{
		Message: response.Message,
		Code:    response.Code,
	})
}

func (f *fileHandler) GetFile(c *gin.Context) {
	fileId := c.Param(constant.FileId)
	response, err := f.fileUsecase.GetFile(c, fileId)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}
	c.File(response.FilePath)
}
