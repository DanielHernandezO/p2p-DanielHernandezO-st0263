package gateway

import "github.com/DanielHernandezO/p2p/download/internal/business/domain"

type DownloadGateway interface {
	SendMessage(message *domain.QueueMessage) error
}
