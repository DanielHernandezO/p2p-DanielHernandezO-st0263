package gateway

import "context"

type ServerServiceGateway interface {
	ExecuteJoin(context context.Context) string
	Search(context context.Context, filename string) string
}
