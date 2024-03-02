package gateway

import "github.com/DanielHernandezO/p2p/search/internal/business/domain"

type DownloadQueue interface {
	AddMessage(message domain.QueueMessage) error
}
