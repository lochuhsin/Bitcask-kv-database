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
	String    DataType = "STRING"
	Tombstone DataType = "TOMBSTONE"
)

type Base interface {
	Format() []byte // should be ValueDataType::ValueLen::Value
	GetVal() any
	GetType() DataType
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
	Val        []byte
	CreateTime int64 // timestamp time.Now().UnixMicro()
}

func InitEntry(key []byte, val []byte) Entry {
	return Entry{
		key, val, time.Now().UnixMicro(),
	}
}

func InitTombEntry(key []byte) Entry {
	return Entry{
		key, util.StringToBytes(setting.Config.TOMBSTONE), time.Now().UnixMicro(),
	}
}
