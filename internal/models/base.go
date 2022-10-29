package models

type KVPair struct {
	Key string
	Val []byte
}

type MemoryModel interface {
	Init()
	Get()
	Set()
	GetSize()
	GetAll()
}
