package models

type KVPair struct {
	Key string
	Val Item
}

type MemoryModel interface {
	Init()
	Get(string) (Item, bool)
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

type DataType interface {
	ToBytes() ([]byte, error)
}

func (i *NilInt) ToBytes() ([]byte, error) {
	panic("Not Implemented yet")
}

func (i *NilString) ToBytes() ([]byte, error) {
	panic("Not Implemented yet")
}

func (i *NilBool) ToBytes() ([]byte, error) {
	panic("Not Implemented yet")
}

func (i *NilByte) ToBytes() ([]byte, error) {
	panic("Not Implemented yet")
}

type NilInt struct {
	IsNil bool
	Val   int
}

type NilString struct {
	IsNil bool
	Val   string
}

type NilBool struct {
	IsNil bool
	Val   bool
}

type NilByte struct {
	IsNil bool
	Val   byte
}
