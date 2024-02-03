package discovery

type registerRequestSchema struct {
	ServerIP   string `json:"serverIp"`
	ServerName string `json:"serverName"`
}

type status string

const (
	RED    status = "red"
	GREEN  status = "green"
	Yellow status = "yellow"
)

type getClusterStatusSchema struct {
	Status status `json:"status"`
}

type peerSchema struct {
	Name string `json:"serverName"`
	Ip   string `json:"serverIP"`
}

type peerListResponseSchema struct {
	Peers []peerSchema `json:"peers"`
}
