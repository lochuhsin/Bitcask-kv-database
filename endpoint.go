package rebitcask

import (
	"errors"
	"rebitcask/internal/dao"
	"rebitcask/internal/memory"
	"rebitcask/internal/segment"
	"rebitcask/internal/util"
)

func Get(k string) (any, bool) {
	/**
	 * First, check does the value exist in memory
	 *
	 * Second, check does the value exist in segment
	 *
	 * Note: exists meaning that the key exists, and the value is not tombstone
	 */
	bytes := util.StringToBytes(k)
	m, status := memory.GetMemoryManager().Get(bytes)
	if status {
		return checkTombstone(m.Val)
	}

	s, status := segment.GetSegmentManager().GetValue(bytes)
	if status {
		return checkTombstone(s.Val)
	}
	return *new(any), false
}

func Set(k string, v any) error {
	_v, err := convertToBaseObjects(v)
	if err != nil {
		return err
	}
	entry := dao.InitEntry(util.StringToBytes(k), _v)
	err = memory.GetMemoryManager().Set(entry)
	return err
}

func Delete(k string) error {
	entry := dao.InitTombEntry(util.StringToBytes(k))
	return memory.GetMemoryManager().Set(entry)
}

func Exist() (bool, error) {
	panic("Not implemented error")
}

func BulkCreate(k string) error {
	panic("Not implemented error")
}

func BulkUpdate(k string) error {
	panic("Not implemented error")
}

func BulkUpsert(k string) error {
	panic("Not implemented error")
}

func BulkDelete(k string) error {
	panic("Not implemented error")
}

func BulkGet(k ...string) ([]string, error) {
	panic("Not implemented error")
}

func convertToBaseObjects(v any) (dao.Base, error) {
	switch v := v.(type) {
	case int:
		return dao.NilInt{IsNil: false, Val: v}, nil
	case float64:
		return dao.NilFloat{IsNil: false, Val: v}, nil
	case byte:
		return dao.NilByte{IsNil: false, Val: v}, nil
	case string:
		return dao.NilString{
			IsNil: false,
			Val:   util.StringToBytes(v),
		}, nil
	case bool:
		return dao.NilBool{IsNil: false, Val: v}, nil

	default:
		return nil, errors.New("invalid data type")
	}
}

func checkTombstone(val dao.Base) (any, bool) {
	if val.GetType() == dao.Tombstone {
		return new(*any), false
	}

	return val.GetVal(), true
}
