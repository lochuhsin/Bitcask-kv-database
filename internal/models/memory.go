package models

import "sync"

type MemoryMap struct {
	mu       sync.Mutex
	keyvalue map[string][]byte
}

type KVPair struct {
	Key string
	Val []byte
}

func (m *MemoryMap) Init() {
	m.keyvalue = make(map[string][]byte)
}

func (m *MemoryMap) Get(k *string) (b []byte, status bool) {
	if val, ok := m.keyvalue[*k]; ok {
		return val, true
	}
	return []byte(""), false
}

func (m *MemoryMap) Set(k string, v []byte) {
	m.mu.Lock()
	m.keyvalue[k] = v
	m.mu.Unlock()
}

func (m *MemoryMap) GetSize() int {
	return len(m.keyvalue)
}

// GetAll TODO: optimize this
func (m *MemoryMap) GetAll() *[]KVPair {
	arr := make([]KVPair, len(m.keyvalue))
	for k, v := range m.keyvalue {
		arr = append(arr, KVPair{
			Key: k,
			Val: v,
		})
	}
	return &arr
}

// SetMap TODO: optimize this
func (m *MemoryMap) SetMap(kvMap map[string][]byte) {
	m.keyvalue = kvMap
}
