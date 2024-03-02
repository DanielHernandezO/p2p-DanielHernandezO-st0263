package handler

import (
	"net/http"

	"github.com/DanielHernandezO/p2p/download/internal/business/constant"
	"github.com/DanielHernandezO/p2p/download/internal/business/domain"
	"github.com/DanielHernandezO/p2p/download/internal/business/usecase"
	"github.com/gin-gonic/gin"
)

type SaveHandler interface {
	Save(c *gin.Context)
}

type saveHandler struct {
	saveUsecase usecase.SaveUsecase
}

func NewSaveHander(saveUsecase usecase.SaveUsecase) *saveHandler {
	return &saveHandler{
		saveUsecase: saveUsecase,
	}
}

func (s *saveHandler) Save(c *gin.Context) {
	var file domain.QueueMessage

	if err := c.BindJSON(&file); err != nil {
		c.JSON(http.StatusBadRequest, domain.Response{
			Message: constant.BadRequest,
			Code:    http.StatusBadGateway,
		})
		return
	}

	response := s.saveUsecase.Save(&file)

	if response != nil {
		c.JSON(http.StatusBadRequest, domain.Response{
			Message: constant.BadRequest,
			Code:    http.StatusBadGateway,
		})
		return
	}
	c.JSON(http.StatusOK, domain.Response{
		Message: constant.Saved,
		Code:    http.StatusOK,
	})
}
