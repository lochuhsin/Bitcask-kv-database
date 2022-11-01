package models

type KVPair struct {
	Key string
	Val []byte
}

// MemoryModel TODO convert internal memory model to interface
type MemoryModel interface {
	Init()
	Get()
	Set()
	GetSize()
	GetAll()
	GetAllValueUnder()
	GetAllValueOver()
}
