package cache

import (
	"context"
	"encoding/json"
)

type ISchema interface {
	MarshalBinary() ([]byte, error)
}

func (p PeerSchema) MarshalBinary() ([]byte, error) {
	return json.Marshal(p)
}

func (c ClusterSchema) MarshalBinary() ([]byte, error) {
	return json.Marshal(c)
}

type ICache interface {
	Get(context.Context, string) (any, bool)
	Set(context.Context, string, ISchema) error
	GetAllValues(context.Context) any
	GetAllKeys(context.Context) []string
	Count(context.Context) int
}
