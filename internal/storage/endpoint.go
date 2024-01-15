package storage

import (
	"errors"
	"rebitcask/internal/storage/dao"
	"rebitcask/internal/storage/service"
)

func Get(k string) (any, bool) {
	/**
	 * First, we pass through counting bloom filter
	 * if exists, then we continue the next step
	 * else: return directly
	 *
	 * Second, check if the value exists in memory
	 *
	 * Third, check if the value exists in current open segment
	 *
	 * Finally, check if the value exists in old segment
	 *
	 * Note: exists meaning that the key exists, and the value is not tombstoneed
	 */
	// check if exists in cbf
	_k, err := daoConverter(k)
	if err != nil {
		panic(err) // TODO: better handling
	}

	m, status := service.MGet(_k.(dao.NilString))
	if status {
		return filterTombstone(m)
	}

	s, status := service.SGet(_k.(dao.NilString))
	if status {
		return filterTombstone(s)
	}
	return *new(any), false
}

func Set(k string, v any) error {
	_k, err := daoConverter(k)
	if err != nil {
		return err
	}
	_v, err := daoConverter(v)
	if err != nil {
		return err
	}

	service.MSet(_k.(dao.NilString), _v)
	return nil
}

func Delete(k string) error {
	_k, err := daoConverter(k)
	if err != nil {
		return err
	}
	service.MDelete(_k.(dao.NilString))
	return nil
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

func daoConverter(v any) (dao.Base, error) {
	switch v := v.(type) {
	case int:
		return dao.NilInt{IsNil: false, Val: v}, nil
	case float64:
		return dao.NilFloat{IsNil: false, Val: v}, nil
	case byte:
		return dao.NilByte{IsNil: false, Val: v}, nil
	case string:
		return dao.NilString{IsNil: false, Val: v}, nil
	case bool:
		return dao.NilBool{IsNil: false, Val: v}, nil

	default:
		return nil, errors.New("invalid data type")
	}
}

func filterTombstone(val dao.Base) (any, bool) {
	if val.GetType() == dao.Tombstone {
		return new(*any), false
	}

	return val.GetVal(), true
}
