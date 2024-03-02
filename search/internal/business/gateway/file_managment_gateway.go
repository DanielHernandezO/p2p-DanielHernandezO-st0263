package gateway

import "github.com/DanielHernandezO/p2p/search/internal/business/domain"

type FileManagmentProvider interface {
	GetFile(fileId string) *domain.FilesResponse
}
