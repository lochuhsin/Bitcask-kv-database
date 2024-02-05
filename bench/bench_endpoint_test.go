package test

import (
	"os"
	"rebitcask"
	"testing"
	"time"
)

func setup() {
	rebitcask.Setup(".env.bench")
}

func teardown() {
	time.Sleep(time.Second * 3)
	removeSegment()
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}

func BenchmarkFullSearchStorageGet(b *testing.B) {
	keys, _ := GenerateLowDuplicateRandom(b.N)
	for _, k := range keys {
		_, _ = rebitcask.Get(k)
	}
}

func BenchmarkStorageSet(b *testing.B) {
	keys, vals := GenerateLowDuplicateRandom(b.N)
	for i, k := range keys {
		_ = rebitcask.Set(k, vals[i])
	}
}

func BenchmarkStorageGet(b *testing.B) {
	keys, _ := GenerateLowDuplicateRandom(b.N)
	for _, k := range keys {
		_, _ = rebitcask.Get(k)
	}
}

func BenchmarkStorageDelete(b *testing.B) {
	keys, vals := GenerateLowDuplicateRandom(b.N)
	for i, k := range keys {
		_ = rebitcask.Set(k, vals[i])
	}
}

func BenchmarkStorageSetGet(b *testing.B) {
	keys, vals := GenerateLowDuplicateRandom(b.N)
	for i, k := range keys {
		_ = rebitcask.Set(k, vals[i])
		_, _ = rebitcask.Get(k)
	}
}
