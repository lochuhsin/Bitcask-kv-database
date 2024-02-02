package core

type dataRequestSchema struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type dataDeleteSchema struct {
	Key string `json:"key"`
}
