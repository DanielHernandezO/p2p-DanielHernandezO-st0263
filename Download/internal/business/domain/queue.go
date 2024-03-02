package domain

type QueueMessage struct {
	Name    string `json:"name"`
	Ip      string `json:"ip"`
	Port    string `json:"port"`
	Content []byte `json:"content"`
}
