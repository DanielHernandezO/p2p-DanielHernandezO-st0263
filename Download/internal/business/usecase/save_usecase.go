package usecase

import (
	"log"
	"os"

	"github.com/DanielHernandezO/p2p/download/internal/business/domain"
	"github.com/DanielHernandezO/p2p/download/internal/infraestructure/config"
)

type SaveUsecase interface {
	Save(file *domain.QueueMessage) error
}

type saveUsecase struct {
	location string
}

func NewSaveUsecase() *saveUsecase {
	return &saveUsecase{
		location: config.GetStringPropetyBykey(config.Location),
	}
}
func (s *saveUsecase) Save(file *domain.QueueMessage) error {
	err := os.WriteFile(s.location+file.Name, file.Content, 0644)
	if err != nil {
		log.Fatal("Error al guardar el archivo:", err)
		return err
	}
	return nil
}
