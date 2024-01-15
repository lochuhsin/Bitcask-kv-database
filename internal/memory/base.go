package memory

import (
	"rebitcask/internal/dao"
)

type IMemory interface {
	Get(dao.NilString) (dao.Base, bool)
	Set(dao.Pair) // if the memory is in frozen state, close set operation
	GetSize() int
	GetAll() []dao.Pair // Expected order by key from small to large
	GetAllValueUnder(dao.NilString) []dao.Pair
	Reset()
	Isfrozen() bool
	Setfrozen(bool)
	Clone() IMemory
}
