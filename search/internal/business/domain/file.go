package domain

type FileRequested struct {
	Id   string
	Name string
	Ip   string
	Port string
}

type FilesResponse struct {
	Bytes []byte
}
