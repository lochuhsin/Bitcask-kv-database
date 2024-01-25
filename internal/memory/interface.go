package memory

import (
	"rebitcask/internal/dao"
)

type IMemory interface {
	Get(dao.NilString) (dao.Base, bool)
	Set(dao.Entry) // if the memory is in frozen state, close set operation
	GetSize() int
	GetAll() []dao.Entry // Expected order by key from small to large
}
