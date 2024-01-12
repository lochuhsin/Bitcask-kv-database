package test

import (
	"fmt"
	"os"
	"rebitcask/internal/settings"
	"rebitcask/internal/storage"
	"testing"
	"time"
)

func setup() {
	storage.Init()
}

func teardown() {
	time.Sleep(time.Millisecond * 100)
	removeSegment()
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}

func TestStorageSet(t *testing.T) {
	env := settings.ENV
	dataCount := env.MemoryCountLimit*10 + 1

	keys, vals := GenerateLowDuplicateRandom(dataCount)
	for i, k := range keys {
		err := storage.Set(k, vals[i])
		if err != nil {
			t.Error("Something went wrong while setting")
		}
	}
}

func TestStorageGet(t *testing.T) {

	env := settings.ENV
	dataCount := env.MemoryCountLimit*10 + 1

	keys, vals := GenerateLowDuplicateRandom(dataCount)
	for i, k := range keys {
		err := storage.Set(k, vals[i])
		if err != nil {
			t.Error("Something went wrong while setting")
		}
	}
}

func TestStorageDelete(t *testing.T) {

	env := settings.ENV
	dataCount := env.MemoryCountLimit*10 + 1
	keys, _ := GenerateLowDuplicateRandom(dataCount)
	for _, k := range keys {
		err := storage.Delete(k)
		if err != nil {
			t.Error("Something went wrong while setting")
		}
	}
}

func TestStorageSetGet(t *testing.T) {

	env := settings.ENV
	dataCount := env.MemoryCountLimit*10 + 1

	keys, vals := GenerateLowDuplicateRandom(dataCount)
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
	env := settings.ENV
	dataCount := env.MemoryCountLimit*10 + 1

	keys, vals := GenerateLowDuplicateRandom(dataCount)
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

	for i, k := range keys {
		val, status := storage.Get(k)
		if status {
			str := fmt.Sprintf("the key should not exist: %v", val)
			t.Error(str)
			fmt.Println(val == vals[i])
		}
	}
}

func TestEmptyGet(t *testing.T) {
	keys, _ := GenerateLowDuplicateRandom(100)

	for _, k := range keys {
		_, status := storage.Get(k)
		if status {
			t.Error("the key should not exist")
		}
	}
}
