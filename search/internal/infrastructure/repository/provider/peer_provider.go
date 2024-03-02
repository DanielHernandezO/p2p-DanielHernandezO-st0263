package provider

import (
	"context"
	"fmt"
	"log"

	"github.com/DanielHernandezO/p2p/search/internal/business/domain"
	"github.com/DanielHernandezO/p2p/search/internal/infrastructure/config"
	"google.golang.org/grpc"
)

type peerProvider struct {
}

func NewPeerProvider() *peerProvider {
	return &peerProvider{}
}

func (p *peerProvider) ExecuteJoin(context *context.Context, socket *domain.Socket, peers map[string]string) bool {
	exitConnections := 0
	for key, value := range peers {
		conn := checkPeerConnection(key, value)
		if conn == nil {
			continue
		}
		defer conn.Close()
		client := config.NewJoinClient(conn)
		body := &config.JoinRequest{
			Ip:   socket.Ip,
			Port: socket.Port,
		}
		_, err := client.Join(*context, body)
		if err != nil {
			continue
		}
		exitConnections = exitConnections + 1
	}
	return exitConnections != 0
}

func (p *peerProvider) SearchFile(context *context.Context, fileRequested *domain.FileRequested, peers map[string]string) bool {
	foundFiles := 0
	for key, value := range peers {
		conn := checkPeerConnection(key, value)
		if conn == nil {
			continue
		}
		defer conn.Close()
		client := config.NewSearchFileClient(conn)
		bodyRequest := &config.FileRequested{
			Id:           fileRequested.Id,
			Name:         fileRequested.Name,
			IpResponse:   fileRequested.Ip,
			PortResponse: fileRequested.Port,
		}
		_, err := client.Search(*context, bodyRequest)
		if err != nil {
			continue
		}
		foundFiles = foundFiles + 1
	}
	return foundFiles != 0
}

func (p *peerProvider) ExecuteFetch(context *context.Context, peers map[string]string) map[string]string {
	newPeers := make(map[string]string)
	for ip, port := range peers {
		newPeers[ip] = port
	}
	for key, value := range peers {
		conn := checkPeerConnection(key, value)
		if conn == nil {
			continue
		}
		defer conn.Close()
		client := config.NewFetchClient(conn)
		body := &config.VoidRequest{}
		knownPeers, err := client.Fetch(*context, body)
		if err != nil {
			continue
		}
		for ip, port := range knownPeers.Peers {
			newPeers[ip] = port
		}
	}
	return newPeers
}

func checkPeerConnection(ip string, port string) *grpc.ClientConn {

	address := fmt.Sprintf("%s:%s", ip, port)
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Printf("error conecting to:%s", address)
		return nil
	}
	return conn

}
