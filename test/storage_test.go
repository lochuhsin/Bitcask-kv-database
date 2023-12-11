package test

import (
	"rebitcask/internal/settings"
	"rebitcask/internal/storage"
	"testing"
)

func TestStorageSet(t *testing.T) {
	defer removeSegment()
	storage.Init()

	env := settings.ENV
	dataCount := env.MemoryCountLimit*10 + 1

	keys, vals := generateLowDuplicateRandomData(dataCount)
	for i, k := range keys {
		err := storage.Set(k, vals[i])
		if err != nil {
			t.Error("Something went wrong while setting")
		}
	}
}

func TestStorageGet(t *testing.T) {
	defer removeSegment()
	storage.Init()

	env := settings.ENV
	dataCount := env.MemoryCountLimit*10 + 1

	keys, vals := generateLowDuplicateRandomData(dataCount)
	for i, k := range keys {
		err := storage.Set(k, vals[i])
		if err != nil {
			t.Error("Something went wrong while setting")
		}
	}
}

func TestStorageDelete(t *testing.T) {
	defer removeSegment()
	storage.Init()
	env := settings.ENV
	dataCount := env.MemoryCountLimit*10 + 1
	keys, _ := generateLowDuplicateRandomData(dataCount)
	for _, k := range keys {
		err := storage.Delete(k)
		if err != nil {
			t.Error("Something went wrong while setting")
		}
	}
}

func TestStorageSetGet(t *testing.T) {
	defer removeSegment()
	storage.Init()

	env := settings.ENV
	dataCount := env.MemoryCountLimit*10 + 1

	keys, vals := generateLowDuplicateRandomData(dataCount)
	for i, k := range keys {
		err := storage.Set(k, vals[i])
		if err != nil {
			t.Error("Something went wrong while setting")
		}
	}

	for i, k := range keys {
		val, status := storage.Get(k)
		if !status {
			t.Error("the key should exist")
		}

		if val != vals[i] {
			t.Error("the value should be equal to the generated value")
		}
	}
}

func TestStorageSetDelete(t *testing.T) {
	defer removeSegment()
	storage.Init()

	env := settings.ENV
	dataCount := env.MemoryCountLimit*10 + 1

	keys, vals := generateLowDuplicateRandomData(dataCount)
	for i, k := range keys {
		err := storage.Set(k, vals[i])
		if err != nil {
			t.Error("Something went wrong while setting")
		}
	}

	for _, k := range keys {
		err := storage.Delete(k)
		if err != nil {
			t.Error("Delete operation should work")
		}

	}

	for _, k := range keys {
		_, status := storage.Get(k)
		if status {
			t.Error("the key should not exist")
		}
	}
}

func TestEmptyGet(t *testing.T) {
	defer removeSegment()
	storage.Init()
	keys, _ := generateLowDuplicateRandomData(100)

	for _, k := range keys {
		_, status := storage.Get(k)
		if status {
			t.Error("the key should not exist")
		}
	}
}
