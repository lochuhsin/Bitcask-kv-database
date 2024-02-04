package cluster

type ClusterStatusSchema struct {
	Status ClusterStatus `json:"status"`
}

type ClusterConfigurationSchema struct {
	MemberCount int `json:"memberCount"`
}

type registerRequestSchema struct {
	Name string `json:"serverName"`
	Ip   string `json:"serverIP"`
}

type registerResponseSchema struct {
	Message string `json:"message"`
}

type peerSchema struct {
	Name string `json:"serverName"`
	Ip   string `json:"serverIP"`
}

type peerListResponseSchema struct {
	Peers []peerSchema `json:"peers"`
}
