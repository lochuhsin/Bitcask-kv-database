package cache

type PeerSchema struct {
	HostPort string `json:"HostPort"`
}

type ClusterSchema struct {
	Value string `json:"Value"`
}

type PeerCounterSchema struct{}
