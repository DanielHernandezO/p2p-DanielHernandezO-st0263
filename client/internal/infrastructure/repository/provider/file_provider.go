package provider

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/DanielHernandezO/p2p/client/internal/business/constant"
	"github.com/DanielHernandezO/p2p/client/internal/infrastructure/config"
)

const (
	providerName = "filemanagment"
)

type fileProvider struct {
	baseUrl string
	ip      string
	port    string
}

func NewFileProvider() *fileProvider {
	return &fileProvider{
		baseUrl: config.GetStringPropetyBykey(fmt.Sprintf(config.BaseUrl, providerName)),
		ip:      config.GetStringPropetyBykey(fmt.Sprintf(config.IpConfig, providerName)),
		port:    config.GetStringPropetyBykey(fmt.Sprintf(config.PortConfig, providerName)),
	}
}

func (f *fileProvider) PostFile(context context.Context, filePath string) string {
	url := fmt.Sprintf(f.baseUrl, f.ip, f.port) + config.GetStringPropetyBykey(config.PostFilePath)
	file, err := os.Open(filePath)
	if err != nil {
		log.Print(err.Error())
		return fmt.Sprintf(constant.ErrorFilemanagment, filePath)
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile("file", filepath.Base(filePath))
	if err != nil {
		log.Print(err.Error())
		return fmt.Sprintf(constant.ErrorFilemanagment, filePath)
	}

	_, err = io.Copy(part, file)
	if err != nil {
		log.Print(err.Error())
		return fmt.Sprintf(constant.ErrorFilemanagment, filePath)
	}

	err = writer.Close()
	if err != nil {
		log.Print(err.Error())
		return fmt.Sprintf(constant.ErrorFilemanagment, filePath)
	}

	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		log.Print(err.Error())
		return fmt.Sprintf(constant.ErrorFilemanagment, filePath)
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Print(err.Error())
		return fmt.Sprintf(constant.ErrorFilemanagment, filePath)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("error with status %v", resp.StatusCode)
		return fmt.Sprintf(constant.ErrorFilemanagment, filePath)
	}

	return constant.PostedFile
}
