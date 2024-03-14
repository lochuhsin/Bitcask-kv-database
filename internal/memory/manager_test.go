package memory

import (
	"fmt"
	"rebitcask/internal/dao"
	"rebitcask/internal/memory/models"
	"strconv"
	"testing"
)

func TestMemoryManagerEmptyGet(t *testing.T) {
	entryCountLimit := 1
	blockChSize := 100
	blockStorage := NewMemoryStorage()
	manager := NewMemoryManager(
		blockStorage, entryCountLimit, blockChSize, models.ModelType("bst"),
	)
	_, status := manager.Get([]byte(""))
	if status {
		t.Error("Key Should not exists")
	}
}
func TestMemoryManagerSetBelowLimit(t *testing.T) {
	entryCount := 10
	entryCountLimit := 1000
	blockChSize := 10000
	blockStorage := NewMemoryStorage()
	manager := NewMemoryManager(
		blockStorage, entryCountLimit, blockChSize, models.ModelType("bst"),
	)

	for i := 0; i < entryCount; i++ {
		kStr := strconv.Itoa(i)
		vStr := strconv.Itoa(i)
		entry := dao.InitEntry([]byte(kStr), []byte(vStr))
		manager.Set(entry)
	}

	for i := 0; i < entryCount; i++ {
		kStr := strconv.Itoa(i)
		_, status := manager.Get([]byte(kStr))
		if !status {
			t.Error("entry should exists")
		}
	}
}

func TestMemoryManagerSetAboveLimit(t *testing.T) {
	entryCount := 100
	entryCountLimit := 1
	blockChSize := 10000
	blockStorage := NewMemoryStorage()
	manager := NewMemoryManager(
		blockStorage, entryCountLimit, blockChSize, models.ModelType("bst"),
	)

	for i := 0; i < entryCount; i++ {
		kStr := strconv.Itoa(i)
		vStr := strconv.Itoa(i)
		entry := dao.InitEntry([]byte(kStr), []byte(vStr))
		manager.Set(entry)
	}

	for i := 0; i < entryCount; i++ {
		kStr := strconv.Itoa(i)
		_, status := manager.Get([]byte(kStr))
		if !status {
			t.Error("entry should exists")
		}
	}
	blockCount := manager.GetTotalBlockCount()
	if blockCount != entryCount+1 {
		fmt.Println(blockCount)
		t.Error("wrong number of blocks")
	}
}
