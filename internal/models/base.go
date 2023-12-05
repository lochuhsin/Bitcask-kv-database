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

type DataType string

const (
	Int    DataType = "INT"
	Float  DataType = "FLOAT"
	Bool   DataType = "BOOL"
	String DataType = "STRING"
	Byte   DataType = "BYTE"
	Nil    DataType = "NIL"
)

/**
 * TODO: add preserved keywords CRC::
 */

/**
 * DataStorage format is
 * CRC::Timestamp::KeyDataType::ValueDataType::KeyLen::ValueLen::Key::Value
 *
 * Note: KeyLen, ValueLen should be converted to
 * actual data bytes size in memory)
 */

type Data interface {
	ToData() ([]byte, error)
}

func (i *NilNil) ToData() ([]byte, error) {
	panic("Not implemented yet")
}

func (i *NilInt) ToData() ([]byte, error) {
	panic("Not Implemented yet")
}

func (i *NilFloat) ToData() ([]byte, error) {
	panic("Not Implemented yet")
}

func (i *NilString) ToData() ([]byte, error) {
	panic("Not Implemented yet")
}

func (i *NilBool) ToData() ([]byte, error) {
	panic("Not Implemented yet")
}

func (i *NilByte) ToData() ([]byte, error) {
	panic("Not Implemented yet")
}

type NilNil struct {
	IsNil bool
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

// use 64 for convenience
type NilFloat struct {
	IsNil bool
	Val   float64
	Key   string
}
