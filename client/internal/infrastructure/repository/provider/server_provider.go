package provider

import (
	"context"
	"fmt"
	"log"

	"github.com/DanielHernandezO/p2p/client/internal/business/constant"
	"github.com/DanielHernandezO/p2p/client/internal/infrastructure/config"
	"github.com/google/uuid"
	"google.golang.org/grpc"
)

var server = "server"

type serverProvider struct {
	Ip    string
	Port  string
	Peers map[string]string
}

func NewServerProvider() *serverProvider {
	ip := config.GetStringPropetyBykey(fmt.Sprintf(config.IpConfig, server))
	port := config.GetStringPropetyBykey(fmt.Sprintf(config.PortConfig, server))
	peers := config.GetMapPropertyByKey(config.PeersConfig)
	return &serverProvider{
		Ip:    ip,
		Port:  port,
		Peers: peers,
	}
}

func (s *serverProvider) ExecuteJoin(context context.Context) string {
	conn := checkPeerConnection(s.Ip, s.Port)
	if conn == nil {
		return constant.ErrorServer
	}
	defer conn.Close()
	client := config.NewJoinClient(conn)
	body := &config.JoinPeersRequest{
		Ip:    s.Ip,
		Port:  s.Port,
		Peers: s.Peers,
	}
	_, err := client.ExecuteJoin(context, body)
	if err != nil {
		return constant.ErrorServer
	}

	return constant.Connected
}

func (s *serverProvider) Search(context context.Context, filename string) string {
	downloadIp := config.GetStringPropetyBykey(fmt.Sprintf(config.IpConfig, "download"))
	downloadPort := config.GetStringPropetyBykey(fmt.Sprintf(config.PortConfig, "download"))
	conn := checkPeerConnection(s.Ip, s.Port)
	if conn == nil {
		return constant.ErrorServer
	}
	defer conn.Close()
	client := config.NewSearchFileClient(conn)
	body := &config.FileRequested{
		Id:           fmt.Sprintf("%s-%s", s.Ip, uuid.New()),
		Name:         filename,
		IpResponse:   downloadIp,
		PortResponse: downloadPort,
	}
	answer, err := client.Search(context, body)
	if err != nil {
		return constant.ErrorServer
	}
	return answer.Message
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
