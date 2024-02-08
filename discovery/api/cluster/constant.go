package cluster

type ClusterStatus string

// Borrowing from elasticsearch cluster lol
const (
	RED    ClusterStatus = "red"
	YELLOW ClusterStatus = "yellow"
	GREEN  ClusterStatus = "green"
)
