package provider

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/DanielHernandezO/p2p/search/internal/business/constant"
	"github.com/DanielHernandezO/p2p/search/internal/business/domain"
	"github.com/DanielHernandezO/p2p/search/internal/infrastructure/config"
	"github.com/go-playground/validator/v10"
)

const (
	providerName = "filemanagment"
)

type fileManagmentProvider struct {
	baseUrl  string
	validate *validator.Validate
	ip       string
	port     string
}

func NewFileManagmentProvider() *fileManagmentProvider {
	return &fileManagmentProvider{
		baseUrl:  config.GetStringPropetyBykey(fmt.Sprintf(config.BaseUrl, providerName)),
		validate: validator.New(),
		ip:       config.GetStringPropetyBykey(fmt.Sprintf(config.IpConfig, providerName)),
		port:     config.GetStringPropetyBykey(fmt.Sprintf(config.PortConfig, providerName)),
	}
}

func (f *fileManagmentProvider) GetFile(fileId string) *domain.FilesResponse {
	url := fmt.Sprintf(f.baseUrl, f.ip, f.port) + fmt.Sprintf(config.GetStringPropetyBykey(config.GetFilePath), fileId)
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("error executing HTTP: %v", err)
		return nil
	}
	defer resp.Body.Close()

	if isError(resp.StatusCode) {
		log.Printf(constant.NotFound)
		return nil
	}
	log.Printf(constant.Found)

	fileContent, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("error reading file content")
		return nil
	}

	return &domain.FilesResponse{
		Bytes: fileContent,
	}
}
