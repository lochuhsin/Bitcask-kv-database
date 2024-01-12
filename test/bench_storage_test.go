package test

import (
	"rebitcask/internal/storage"
	"testing"
)

func BenchmarkFullSearchStorageGet(b *testing.B) {
	keys, _ := GenerateLowDuplicateRandom(b.N)
	for _, k := range keys {
		_, _ = storage.Get(k)
	}
}

func BenchmarkStorageSet(b *testing.B) {
	keys, vals := GenerateLowDuplicateRandom(b.N)
	for i, k := range keys {
		_ = storage.Set(k, vals[i])
	}
}

func BenchmarkStorageGet(b *testing.B) {
	keys, _ := GenerateLowDuplicateRandom(b.N)
	for _, k := range keys {
		_, _ = storage.Get(k)
	}
}

func BenchmarkStorageDelete(b *testing.B) {
	keys, _ := GenerateLowDuplicateRandom(b.N)
	for _, k := range keys {
		_ = storage.Delete(k)
	}
}

func BenchmarkStorageSetGet(b *testing.B) {
	keys, vals := GenerateLowDuplicateRandom(b.N)
	for i, k := range keys {
		_ = storage.Set(k, vals[i])
		_, _ = storage.Get(k)
	}
}
