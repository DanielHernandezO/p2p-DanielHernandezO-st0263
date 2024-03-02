package publisher

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/DanielHernandezO/p2p/search/internal/business/constant"
	"github.com/DanielHernandezO/p2p/search/internal/business/domain"
	"github.com/DanielHernandezO/p2p/search/internal/infrastructure/config"
	"github.com/streadway/amqp"
)

var (
	rabbit    = "rabbit"
	queueName = "download"
)

type downloadPublisher struct {
	baseUrl string
}

func NewDownloadQueue() *downloadPublisher {
	user := config.GetStringPropetyBykey(fmt.Sprintf(config.User, rabbit))
	password := config.GetStringPropetyBykey(fmt.Sprintf(config.Password, rabbit))
	url := config.GetStringPropetyBykey(fmt.Sprintf(config.BaseUrl, rabbit))
	return &downloadPublisher{
		baseUrl: fmt.Sprintf(url, user, password),
	}
}

func (d *downloadPublisher) Start() {
	ip := config.GetStringPropetyBykey(fmt.Sprintf(config.IpConfig, rabbit))
	port := config.GetStringPropetyBykey(fmt.Sprintf(config.PortConfig, rabbit))
	address := fmt.Sprintf("%s:%s/", ip, port)
	url := d.baseUrl + address

	conn, err := amqp.Dial(url)
	if err != nil {
		log.Fatal(constant.PostingMessageError)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal(constant.CreatingChannelError)
	}
	defer ch.Close()

	exchange := config.GetStringPropetyBykey(fmt.Sprintf(config.ExchangeQueue, queueName))
	queue := config.GetStringPropetyBykey(fmt.Sprintf(config.QueueName, queueName))

	err = ch.ExchangeDeclare(
		exchange,
		"direct",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal(constant.PostingMessageError)
	}

	q, err := ch.QueueDeclare(
		queue,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal(constant.PostingMessageError)
	}

	err = ch.QueueBind(
		q.Name,
		queue,
		exchange,
		false,
		nil,
	)
	if err != nil {
		log.Fatal(constant.PostingMessageError)
	}
}

func (d *downloadPublisher) AddMessage(message domain.QueueMessage) error {
	ip := config.GetStringPropetyBykey(fmt.Sprintf(config.IpConfig, rabbit))
	port := config.GetStringPropetyBykey(fmt.Sprintf(config.PortConfig, rabbit))
	address := fmt.Sprintf("%s:%s/", ip, port)
	url := d.baseUrl + address

	conn, err := amqp.Dial(url)
	if err != nil {
		log.Print(constant.PostingMessageError)
		return fmt.Errorf(constant.PostingMessageError)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Print(constant.CreatingChannelError)
		return fmt.Errorf(constant.CreatingChannelError)
	}
	defer ch.Close()

	exchange := config.GetStringPropetyBykey(fmt.Sprintf(config.ExchangeQueue, queueName))
	queue := config.GetStringPropetyBykey(fmt.Sprintf(config.QueueName, queueName))

	jsonBody, err := json.Marshal(message)
	if err != nil {
		log.Print(constant.MarshalJsonError)
		return fmt.Errorf(constant.MarshalJsonError)
	}

	err = ch.Publish(
		exchange,
		queue,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/json",
			Body:        jsonBody,
		},
	)
	if err != nil {
		log.Print(constant.MessageNotPosted)
		return fmt.Errorf(constant.MessageNotPosted)
	}
	log.Printf(constant.MessagePosted, queue)
	return nil
}
