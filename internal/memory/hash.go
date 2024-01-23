package memory

import (
	"rebitcask/internal/dao"
	"sort"
	"sync"
)

type value struct {
	createTime int64
	val        dao.Base
}

type Hash struct {
	keyvalue map[dao.NilString]value
	mu       *sync.Mutex
	frozen   bool
}

func NewHash() *Hash {
	return &Hash{keyvalue: map[dao.NilString]value{}, mu: &sync.Mutex{}, frozen: false}
}

func (m *Hash) Get(k dao.NilString) (b dao.Base, status bool) {
	if val, ok := m.keyvalue[k]; ok {
		return val.val, true
	}
	return nil, false
}

func (m *Hash) Set(pair dao.Pair) {
	m.keyvalue[pair.Key] = value{pair.CreateTime, pair.Val}
}

func (m *Hash) GetSize() int {
	return len(m.keyvalue)
}

func (m *Hash) GetAll() []dao.Pair {
	/**
	 * TODO: implement sort feature
	 */
	arr := make([]dao.Pair, 0, len(m.keyvalue))
	for k, v := range m.keyvalue {
		arr = append(arr, dao.Pair{
			Key:        k,
			Val:        v.val,
			CreateTime: v.createTime,
		})
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i].Key.IsSmaller(arr[j].Key)
	})
	return arr
}
