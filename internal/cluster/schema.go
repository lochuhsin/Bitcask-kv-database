package cluster

type registerRequestSchema struct {
	ServerIP   string `json:"serverIp"`
	ServerName string `json:"serverName"`
}

type ClusterStatus string

const (
	RED    ClusterStatus = "red"
	GREEN  ClusterStatus = "green"
	Yellow ClusterStatus = "yellow"
)

type getClusterStatusSchema struct {
	Status string `json:"status"`
}

type peerSchema struct {
	Name string `json:"serverName"`
	Ip   string `json:"serverIP"`
}

type peerListResponseSchema struct {
	Peers []peerSchema `json:"peers"`
}
