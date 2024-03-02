package gateway

import (
	"context"

	"github.com/DanielHernandezO/p2p/search/internal/business/domain"
)

type PeerServiceGateway interface {
	ExecuteJoin(context *context.Context, socket *domain.Socket, peers map[string]string) bool
	SearchFile(context *context.Context, fileRequested *domain.FileRequested, peers map[string]string) bool
	ExecuteFetch(context *context.Context, peers map[string]string) map[string]string
}
