package memory

import (
	"rebitcask/internal/dao"
	"rebitcask/internal/util"
	"sort"
	"sync"
)

type value struct {
	createTime int64
	val        dao.Entry
}

type Hash struct {
	keyvalue map[string]value
	mu       *sync.Mutex
	frozen   bool
}

func NewHash() *Hash {
	return &Hash{keyvalue: map[string]value{}, mu: &sync.Mutex{}, frozen: false}
}

func (m *Hash) Get(k []byte) (b dao.Entry, status bool) {
	kString := util.BytesToString(k)
	if val, ok := m.keyvalue[kString]; ok {
		return val.val, true
	}
	return dao.Entry{}, false
}

func (m *Hash) Set(entry dao.Entry) {
	m.keyvalue[util.BytesToString(entry.Key)] = value{entry.CreateTime, entry}
}

func (m *Hash) GetSize() int {
	return len(m.keyvalue)
}

func (m *Hash) GetAll() []dao.Entry {
	arr := make([]dao.Entry, 0, len(m.keyvalue))
	for _, v := range m.keyvalue {
		arr = append(arr, v.val)
	}
	sort.Slice(arr, func(i, j int) bool {
		return util.BytesToString(arr[i].Key) <= util.BytesToString(arr[j].Key)
	})
	return arr
}
