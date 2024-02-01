package cluster

type ClusterStatus string

// Borrowing from elasticsearch cluster lol
const (
	RED   ClusterStatus = "red"
	YELLO ClusterStatus = "yello"
	GREEN ClusterStatus = "green"
)
