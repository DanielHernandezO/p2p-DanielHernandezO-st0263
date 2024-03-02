package usecase

import (
	"github.com/DanielHernandezO/p2p/download/internal/business/domain"
	"github.com/DanielHernandezO/p2p/download/internal/business/gateway"
)

type downloadUsecase struct {
	downloadProvider gateway.DownloadGateway
}

func NewDownloadUsecase(downloadProvider gateway.DownloadGateway) *downloadUsecase {
	return &downloadUsecase{
		downloadProvider: downloadProvider,
	}
}

func (d *downloadUsecase) SendMessage(message *domain.QueueMessage) error {
	err := d.downloadProvider.SendMessage(message)
	return err
}
