package domain

type FilesResponse struct {
	FilePath string
}

type Files struct {
	Location string   `json:"location"`
	Names    []string `json:"names"`
}
