package config

import (
	"bufio"
	"log"
	"os"
	"strings"
)

var properties map[string]string

var (
	Port       = "port"
	Location   = "files_location"
	IpConfig   = "ip_%s"
	PortConfig = "port_%s"
	QueueName  = "queue_%s"
	User       = "user_%s"
	Password   = "password_%s"
	BaseUrl    = "base_%s_url"
)

func LoadConfigs() {
	fileroute := "configs/application.properties"
	properties = make(map[string]string)

	file, err := os.Open(fileroute)
	if err != nil {
		log.Fatalf("Error loading configurations: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "=")
		if len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])
			properties[key] = value
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("Error loading configurations: %v", err)
	}
}

func GetStringPropetyBykey(key string) string {
	return properties[key]
}
