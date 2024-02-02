package cluster

type ClusterStatusSchema struct {
	Status ClusterStatus `json:"status"`
}

type ClusterConfigurationSchema struct {
	MemberCount int `json:"memberCount"`
}
