package cache

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"rebitcask/discovery/settings"
	"sync"
	"time"
)

type peer struct {
	sync.Mutex
}

func (p *peer) Get(ctx context.Context, key string) (any, bool) {
	panic("Not implemented error")
}

func (p *peer) Set(ctx context.Context, key string, val ISchema) error {
	p.Lock()
	defer p.Unlock()

	var cursor uint64
	client := GetClient()

	keyPrefix := p.storeKeyPrefix()
	keys, _, err := client.Scan(ctx, cursor, fmt.Sprintf("%v::*", keyPrefix), 1000).Result()

	if err != nil {
		panic(err)
	}

	if len(keys) >= settings.Config.CLUSTER_MEMBER_COUNT {
		return errors.New("peer seats are full")
	}

	storeKey := p.storeKeyFormat(key)
	err = client.Set(ctx, storeKey, val, time.Hour).Err()
	if err != nil {
		panic("something went wrong while setting to cache client")
	}
	return nil
}

func (p *peer) GetAllValues(ctx context.Context) any {
	var cursor uint64
	client := GetClient()
	// NOTE: 1000 is a magic number that is way larger then the amount of the cluster
	// refactor this
	keyPrefix := p.storeKeyPrefix()
	keys, _, err := client.Scan(ctx, cursor, fmt.Sprintf("%v::*", keyPrefix), 1000).Result()
	if err != nil {
		panic(err)
	}

	values := client.MGet(ctx, keys...).Val()
	peers := make([]PeerSchema, len(values))

	for i, v := range values {
		obj := PeerSchema{}
		json.Unmarshal([]byte(v.(string)), &obj)
		peers[i] = obj
	}
	return peers
}

func (p *peer) GetAllKeys(ctx context.Context) []string {
	panic("Not implemented error")
}

func (p *peer) Count(ctx context.Context) int {
	p.Lock()
	defer p.Unlock()

	var cursor uint64
	client := GetClient()

	keyPrefix := p.storeKeyPrefix()
	keys, _, err := client.Scan(ctx, cursor, fmt.Sprintf("%v::*", keyPrefix), 1000).Result()

	if err != nil {
		panic(err)
	}
	return len(keys)
}

func (p *peer) storeKeyFormat(key string) string {
	return fmt.Sprintf("%v::%v", string(PEER), key)
}

func (p *peer) storeKeyPrefix() string {
	return fmt.Sprintf("%v::", string(PEER))
}

type cluster struct {
	sync.Mutex
}

func (c *cluster) Get(ctx context.Context, key string) (any, bool) {
	if !c.keyChecker(key) {
		return nil, false
	}
	client := GetClient()
	storeKey := c.storeKeyFormat(key)
	val := client.Get(ctx, storeKey).Val()
	obj := ClusterSchema{}
	json.Unmarshal([]byte(val), &obj)
	return obj, false
}

func (c *cluster) Set(ctx context.Context, key string, val ISchema) error {
	c.Lock()
	defer c.Unlock()
	if !c.keyChecker(key) {
		return errors.New("invalid cluster key")
	}

	client := GetClient()
	storeKey := c.storeKeyFormat(key)
	return client.Set(ctx, storeKey, val, 1000).Err()
}

func (c *cluster) GetAllValues(ctx context.Context) any {
	panic("Not implemented error")
}

func (c *cluster) GetAllKeys(ctx context.Context) []string {
	panic("Not implemented error")
}

func (c *cluster) Count(ctx context.Context) int {
	panic("Not implemented error")
}

func (c *cluster) keyChecker(key string) bool {
	switch ClusterKeyType(key) {
	case CLUSTER_STATUS:
		return true
	default:
		return false
	}
}

func (c *cluster) storeKeyFormat(key string) string {
	return fmt.Sprintf("%v::%v", string(CLUSTER), key)
}

type peerCounter struct {
}

func (p peerCounter) Get(ctx context.Context, key string) (any, bool) {
	panic("Not implemented error")
}

func (p peerCounter) Set(ctx context.Context, key string, val ISchema) error {
	panic("Not implemented error")
}

func (p peerCounter) GetAllValues(ctx context.Context) any {
	panic("Not implemented error")
}

func (p peerCounter) GetAllKeys(ctx context.Context) []string {
	panic("Not implemented error")
}

func (p peerCounter) Count(ctx context.Context) int {
	panic("Not implemented error")
}

func Get(ctx context.Context, cType CacheType, key string) (any, bool) {
	cache := cacheSelector(cType)
	return cache.Get(ctx, key)
}

func Set(ctx context.Context, cType CacheType, key string, val ISchema) error {
	cache := cacheSelector(cType)
	return cache.Set(ctx, key, val)
}

func Count(ctx context.Context, cType CacheType) int {
	cache := cacheSelector(cType)
	return cache.Count(ctx)
}

func cacheSelector(cType CacheType) ICache {
	switch cType {
	case PEER:
		return &peer{}
	case CLUSTER:
		return &cluster{}
	case PEER_COUNTER:
		return peerCounter{}
	default:
		panic(fmt.Sprintf("Cache type: %v not implemented", cType))
	}
}
