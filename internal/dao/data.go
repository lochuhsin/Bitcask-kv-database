package dao

import (
	"fmt"
	"rebitcask/internal/settings"
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
	Format() string // should be ValueDataType::ValueLen::Value
	GetVal() any
	GetType() DataType
}

func (i NilNil) Format() string {
	panic("Not Implemented yet")
}

func (i NilInt) Format() string {
	intstring := strconv.Itoa(i.GetVal().(int))
	return fmt.Sprintf("%v::%v::%v", Int, len(intstring), intstring)
}

func (i NilFloat) Format() string {
	floatstring := strconv.FormatFloat(i.GetVal().(float64), 'e', 10, 64)
	return fmt.Sprintf("%v::%v::%v", Float, len(floatstring), floatstring)
}

func (i NilString) Format() string {
	str := i.GetVal().(string)
	return fmt.Sprintf("%v::%v::%v", String, len(str), str)
}

func (i NilBool) Format() string {
	boolstr := strconv.FormatBool(i.GetVal().(bool))
	return fmt.Sprintf("%v::%v::%v", Bool, len(boolstr), boolstr)
}

func (i NilByte) Format() string {
	boolstr := string(i.GetVal().(byte))
	return fmt.Sprintf("%v::%v::%v", Byte, len(boolstr), boolstr)
}

func (i NilTomb) Format() string {
	return fmt.Sprintf("%v::%v::%v", Tombstone, len(settings.Config.TOMBSTONE), settings.Config.TOMBSTONE)

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
	Val   string
}

func (i NilString) GetVal() any {
	return i.Val
}

func (i NilString) GetType() DataType {
	return String
}

func (s NilString) IsLarger(d NilString) bool {
	if s.IsNil || d.IsNil {
		return s.nilCompare(d)
	}
	return s.Val > d.Val
}

func (s NilString) IsSmaller(d NilString) bool {
	if s.IsNil || d.IsNil {
		return s.nilCompare(d)
	}
	return s.Val < d.Val
}

func (s NilString) IsEqual(d NilString) bool {
	if s.IsNil || d.IsNil {
		return s.nilCompare(d)
	}
	return s.Val == d.Val
}

func (s NilString) nilCompare(d NilString) bool {
	if s.IsNil && !d.IsNil {
		return false
	} else if !s.IsNil && d.IsNil {
		return true
	} else {
		return false
	}
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
	return settings.Config.TOMBSTONE
	// return tombstone value in envVar initialization
}
func (i NilTomb) GetType() DataType {
	return Tombstone
}

type Entry struct {
	Key        NilString
	Val        Base
	CreateTime int64 // timestamp time.Now().UnixMicro()
}

func InitEntry(key NilString, val Base) Entry {
	return Entry{
		key, val, time.Now().UnixMicro(),
	}
}

func InitTombEntry(key NilString) Entry {
	return Entry{
		key, NilTomb{}, time.Now().UnixMicro(),
	}
}
