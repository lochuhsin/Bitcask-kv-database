package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"rebitcask/discovery/settings"
	"sync"
	"time"
)

/**
 * Refactor this entire class
 */

type PeerCacheSchema struct {
	Name string `json:"Name"`
	Ip   string `json:"Ip"`
}

func (p PeerCacheSchema) MarshalBinary() ([]byte, error) {
	return json.Marshal(p)
}

type peerCache string

var PeerCache peerCache = "peer"

var pMu sync.Mutex

func (p *peerCache) Add(ctx context.Context, peer PeerCacheSchema) bool {
	pMu.Lock()
	defer pMu.Unlock()
	client := GetClient()

	if p.Count(ctx) >= settings.Config.CLUSTER_MEMBER_COUNT {
		return false
	}

	err := client.Set(ctx, fmt.Sprintf("%v::%v", *p, peer.Name), peer, time.Hour*24).Err()
	if err != nil {
		panic(err)
	}
	return true
}

func (p *peerCache) Count(ctx context.Context) int {
	var cursor uint64
	client := GetClient()
	// NOTE: 1000 is a magic number that is way larger then the amount of the cluster
	// refactor this
	result, _, err := client.Scan(ctx, cursor, fmt.Sprintf("%v::*", *p), 1000).Result()
	if err != nil {
		panic(err)
	}
	return len(result)
}

func (p *peerCache) GetAll(ctx context.Context) []PeerCacheSchema {
	var cursor uint64
	client := GetClient()
	// NOTE: 1000 is a magic number that is way larger then the amount of the cluster
	// refactor this
	result, _, err := client.Scan(ctx, cursor, fmt.Sprintf("%v::*", *p), 1000).Result()
	if err != nil {
		panic(err)
	}

	peers := make([]PeerCacheSchema, len(result))

	for i, r := range result {
		obj := PeerCacheSchema{}
		json.Unmarshal([]byte(r), &obj)
		peers[i] = obj
	}
	return peers
}
