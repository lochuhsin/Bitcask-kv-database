package models

type KVPair struct {
	Key string
	Val Item
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

type Item struct {
	Val        []byte
	CreateTime string
	// TODO: add additional attribute from here
}
