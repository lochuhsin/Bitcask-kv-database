package cluster

type ClusterStatusSchema struct {
	Status ClusterStatus `json:"status"`
}

type ClusterConfigrationSchema struct {
	MemberCount int `json:"memberCount"`
}
