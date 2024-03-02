package storage

import (
	"fmt"
	"sync"

	"github.com/DanielHernandezO/p2p/search/internal/business/domain"
	"github.com/DanielHernandezO/p2p/search/internal/infrastructure/config"
)

var (
	mainPeer      = "main"
	alternatePeer = "alternate"
)

type peerStorage struct {
	mu    sync.Mutex
	peers map[string]string
}

func NewPeerStorage() *peerStorage {
	peers := make(map[string]string)
	mainIp := config.GetStringPropetyBykey(fmt.Sprintf(config.IpConfig, mainPeer))
	mainPort := config.GetStringPropetyBykey(fmt.Sprintf(config.PortConfig, mainPeer))
	alternateIp := config.GetStringPropetyBykey(fmt.Sprintf(config.IpConfig, alternatePeer))
	alternatePort := config.GetStringPropetyBykey(fmt.Sprintf(config.PortConfig, alternatePeer))
	peers[mainIp] = mainPort
	peers[alternateIp] = alternatePort
	return &peerStorage{
		peers: peers,
	}
}

func (p *peerStorage) AddPeer(socket *domain.Socket) {
	p.mu.Lock()
	defer p.mu.Unlock()

	p.peers[socket.Ip] = socket.Port
}

func (p *peerStorage) AddPeers(peers map[string]string) {
	p.mu.Lock()
	defer p.mu.Unlock()

	for ip, port := range peers {
		p.peers[ip] = port
	}
}

func (p *peerStorage) DeletePeer(ip string) {
	p.mu.Lock()
	defer p.mu.Unlock()
	delete(p.peers, ip)
}

func (p *peerStorage) GetPeers() map[string]string {
	return p.peers
}
