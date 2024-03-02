package provider

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/DanielHernandezO/p2p/download/internal/business/constant"
	"github.com/DanielHernandezO/p2p/download/internal/business/domain"
	"github.com/DanielHernandezO/p2p/download/internal/infraestructure/config"
)

const (
	providerName = "download"
)

type downloadProvider struct {
	baseUrl string
}

func NewDownloadProvider() *downloadProvider {
	return &downloadProvider{
		baseUrl: config.GetStringPropetyBykey(fmt.Sprintf(config.BaseUrl, providerName)),
	}
}

func (d *downloadProvider) SendMessage(message *domain.QueueMessage) error {
	url := fmt.Sprintf(d.baseUrl, message.Ip, message.Port)

	jsonMessage, err := json.Marshal(message)
	if err != nil {
		log.Print(constant.MarshalError)
		return fmt.Errorf(constant.MarshalError)
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonMessage))
	if err != nil {
		log.Print(constant.HttpError)
		return fmt.Errorf(constant.HttpError)
	}

	defer resp.Body.Close()

	if isError(resp.StatusCode) {
		log.Print(constant.HttpError)
		return fmt.Errorf(constant.HttpError)
	}
	log.Printf(constant.SentMessage)

	return nil
}
