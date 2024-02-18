package dao

import (
	"bytes"
	"rebitcask/internal/setting"
	"rebitcask/internal/util"
	"strconv"
	"time"
)

/**
 * TODO: add preserved keywords ::
 */

/**
 * DataStorage format is:
 * CRC::TimeStamp::KeyDataType::KeyLen::Key::ValueDataType::ValueLen::Value
 *
 * Note: KeyLen, ValueLen should be converted to
 * actual data bytes size in memory)
 *
 * Remove all Sprintf to string builder
 */

type DataType string

const (
	Int       DataType = "INT"
	Float     DataType = "FLOAT"
	Bool      DataType = "BOOL"
	String    DataType = "STRING"
	Byte      DataType = "BYTE"
	Nil       DataType = "NIL"
	Tombstone DataType = "TOMBSTONE"
)

type Base interface {
	Format() []byte // should be ValueDataType::ValueLen::Value
	GetVal() any
	GetType() DataType
}

func (i NilNil) Format() string {
	panic("Not Implemented yet")
}

func (i NilInt) Format() []byte {
	panic("Not Implemented yet")
}

func (i NilFloat) Format() []byte {
	panic("Not Implemented yet")
}

func (i NilString) Format() []byte {
	var builder bytes.Buffer
	str := i.GetVal().(string)
	builder.Write(util.StringToBytes(string(String)))
	builder.Write([]byte("::"))
	builder.Write(util.StringToBytes(strconv.Itoa(len(str))))
	builder.Write([]byte("::"))
	builder.WriteString(str)
	return builder.Bytes()
}

func (i NilBool) Format() []byte {
	panic("Not implemented yet")
}

func (i NilByte) Format() []byte {
	var builder bytes.Buffer
	builder.Write(util.StringToBytes(string(Byte)))
	builder.Write([]byte("::"))
	builder.WriteByte(1)
	builder.Write([]byte("::"))
	builder.WriteByte(i.GetVal().(byte))
	return builder.Bytes()
}

func (i NilTomb) Format() []byte {
	lenString := strconv.Itoa(len(setting.Config.TOMBSTONE))
	var builder bytes.Buffer
	builder.Write(util.StringToBytes(string(Tombstone)))
	builder.Write([]byte("::"))
	builder.Write(util.StringToBytes(lenString))
	builder.Write([]byte("::"))
	builder.Write(util.StringToBytes(string(setting.Config.TOMBSTONE)))
	return builder.Bytes()

}

type NilNil struct {
	IsNil bool
}

type NilInt struct {
	IsNil bool
	Val   int
}

func (i NilInt) GetVal() any {
	return i.Val
}

func (i NilInt) GetType() DataType {
	return Int
}

type NilString struct {
	IsNil bool
	Val   []byte
}

func (i NilString) GetVal() any {
	return string(i.Val)
}

func (i NilString) GetType() DataType {
	return String
}

type NilBool struct {
	IsNil bool
	Val   bool
}

func (i NilBool) GetVal() any {
	return i.Val
}

func (i NilBool) GetType() DataType {
	return Bool
}

type NilByte struct {
	IsNil bool
	Val   byte
}

func (i NilByte) GetVal() any {
	return i.Val
}

func (i NilByte) GetType() DataType {
	return Byte
}

// use 64 for convenience
type NilFloat struct {
	IsNil bool
	Val   float64
	Key   string
}

func (i NilFloat) GetVal() any {
	return i.Val
}

func (i NilFloat) GetType() DataType {
	return Float
}

// it's just ......naming conflicts to Nil is added ...
type NilTomb struct {
}

func (i NilTomb) GetVal() any {
	return setting.Config.TOMBSTONE
	// return tombstone value in envVar initialization
}
func (i NilTomb) GetType() DataType {
	return Tombstone
}

type Entry struct {
	Key        []byte
	Val        Base
	CreateTime int64 // timestamp time.Now().UnixMicro()
}

func InitEntry(key []byte, val Base) Entry {
	return Entry{
		key, val, time.Now().UnixMicro(),
	}
}

func InitTombEntry(key []byte) Entry {
	return Entry{
		key, NilTomb{}, time.Now().UnixMicro(),
	}
}
