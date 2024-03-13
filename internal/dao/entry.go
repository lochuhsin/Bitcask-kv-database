package dao

import (
	"rebitcask/internal/setting"
	"rebitcask/internal/util"
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
