package bootstrap

type registerSchema struct {
	Name string `json:"serverName"`
	Ip   string `json:"serverIP"`
}

type peerSchema struct {
	Name string `json:"serverName"`
	Ip   string `json:"serverIP"`
}

type peerListSchema struct {
	Peers []peerSchema `json:"peers"`
}
