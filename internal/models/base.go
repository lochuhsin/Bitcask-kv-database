package models

type KVPair struct {
	Key string
	Val Item
}

// MemoryModel TODO convert internal memory model to interface
type MemoryModel interface {
	Init()
	Get(*string) (Item, bool)
	Set(string, Item)
	GetSize() int
	GetAll() []KVPair
	GetAllValueUnder(*string) []KVPair
}

type Item struct {
	Val        []byte
	CreateTime string
	// TODO: add additional attribute from here
}
