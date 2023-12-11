package storage

import (
	"errors"
	"fmt"
	"os"
	"rebitcask/internal/settings"
	"rebitcask/internal/storage/cache"
	"rebitcask/internal/storage/dao"
	"rebitcask/internal/storage/memory"
	"rebitcask/internal/storage/service"
)

// move these to env
const (
	cacheType  cache.CacheType  = cache.CBF
	memoryType memory.ModelType = memory.HASH
)

func Init() {
	/**
	 * Should call this, whenever the server is up
	 */
	settings.InitENV()
	env := settings.ENV
	service.CacheInit(cacheType)
	service.MemoryInit(memoryType)
	service.SegmentInit()
	segDir := fmt.Sprintf("%s%s", env.LogPath, env.SegmentFolder)
	os.MkdirAll(segDir, os.ModePerm)
}

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
	if service.CGet(k) {
		_k, err := daoConverter(k)
		if err != nil {
			panic(err) // TODO: better handling
		}
		d, status := service.MGet(_k.(dao.NilString))

		if status && d.GetType() != dao.Tombstone {
			return d.GetVal(), true
		}

		d, status = service.SGet(_k.(dao.NilString))

		if status && d.GetType() != dao.Tombstone {
			return d.GetVal(), true
		}
	} else {
		return *new(any), false
	}
	return *new(any), false
}

func Set(k string, v any) error {
	service.CSet(k)

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
	service.CDelete(k)
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
