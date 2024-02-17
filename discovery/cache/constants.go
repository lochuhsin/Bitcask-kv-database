package cache

type CacheType string

const (
	PEER         CacheType = "peer"
	CLUSTER      CacheType = "cluster"
	PEER_COUNTER CacheType = "peerCounter"
)

type ClusterKeyType string

const (
	CLUSTER_STATUS ClusterKeyType = "cluster_status"
)
