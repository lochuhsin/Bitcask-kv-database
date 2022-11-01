package models

import "sync"

type Hash struct {
	mu       sync.Mutex
	keyvalue map[string]Item
}

func (m *Hash) Init() {
	m.keyvalue = make(map[string]Item)
}

func (m *Hash) Get(k *string) (b Item, status bool) {
	if val, ok := m.keyvalue[*k]; ok {
		return val, true
	}
	return *new(Item), false
}

func (m *Hash) Set(k string, v Item) {
	m.mu.Lock()
	m.keyvalue[k] = v
	m.mu.Unlock()
}

func (m *Hash) GetSize() int {
	return len(m.keyvalue)
}

// GetAll TODO: optimize this
func (m *Hash) GetAll() *[]KVPair {
	arr := make([]KVPair, 0, len(m.keyvalue))
	for k, v := range m.keyvalue {
		arr = append(arr, KVPair{
			Key: k,
			Val: v,
		})
	}
	return &arr
}

// SetMemory TODO: optimize this
func (m *Hash) SetMemory(kvMap map[string]Item) {
	m.keyvalue = kvMap
}
