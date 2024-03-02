package gateway

import "github.com/DanielHernandezO/p2p/search/internal/business/domain"

type PeerStorageGateway interface {
	AddPeer(socket *domain.Socket)
	AddPeers(peers map[string]string)
	DeletePeer(ip string)
	GetPeers() map[string]string
}
