package gateway

import "context"

type FileServiceGateway interface {
	PostFile(context context.Context, filePath string) string
}
