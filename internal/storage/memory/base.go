package memory

import "rebitcask/internal/storage/dao"

type MemoryBase interface {
	Get(dao.NilString) (dao.Base, bool)
	Set(dao.Pair)
	GetSize() int
	GetAll() []dao.Pair // Expected order by key from small to large
	GetAllValueUnder(dao.NilString) []dao.Pair
	Reset()
}
