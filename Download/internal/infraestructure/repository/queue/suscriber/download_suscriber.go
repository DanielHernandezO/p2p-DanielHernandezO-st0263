package suscriber

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/DanielHernandezO/p2p/download/internal/business/constant"
	"github.com/DanielHernandezO/p2p/download/internal/business/domain"
	"github.com/DanielHernandezO/p2p/download/internal/business/gateway"
	"github.com/DanielHernandezO/p2p/download/internal/infraestructure/config"
	"github.com/streadway/amqp"
)

var (
	rabbit    = "rabbit"
	queueName = "download"
)

type DownloadSuscriber interface {
	Suscribe()
}

type downloadSuscriber struct {
	downloadUsecase gateway.DownloadGateway
	baseUrl         string
}

func NewDownloadSuscriber(downloadUsecase gateway.DownloadGateway) *downloadSuscriber {
	user := config.GetStringPropetyBykey(fmt.Sprintf(config.User, rabbit))
	password := config.GetStringPropetyBykey(fmt.Sprintf(config.Password, rabbit))
	url := config.GetStringPropetyBykey(fmt.Sprintf(config.BaseUrl, rabbit))
	return &downloadSuscriber{
		downloadUsecase: downloadUsecase,
		baseUrl:         fmt.Sprintf(url, user, password),
	}
}

func (d *downloadSuscriber) Suscribe() {
	ip := config.GetStringPropetyBykey(fmt.Sprintf(config.IpConfig, rabbit))
	port := config.GetStringPropetyBykey(fmt.Sprintf(config.PortConfig, rabbit))
	address := fmt.Sprintf("%s:%s/", ip, port)
	url := d.baseUrl + address

	conn, err := amqp.Dial(url)
	if err != nil {
		log.Fatal(constant.RabbitMqError)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal(constant.RabbitMqError)
	}
	defer ch.Close()

	queue := config.GetStringPropetyBykey(fmt.Sprintf(config.QueueName, queueName))

	msgs, err := ch.Consume(
		queue,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal(constant.SuscriberError)
	}

	for msg := range msgs {
		var message *domain.QueueMessage
		err := json.Unmarshal(msg.Body, &message)
		if err != nil {
			log.Print(constant.UnmarshalError)
			continue
		}
		err = d.downloadUsecase.SendMessage(message)
		if err != nil {
			log.Printf(err.Error())
		}
		log.Printf(constant.ProcessedMessage)
	}
}
