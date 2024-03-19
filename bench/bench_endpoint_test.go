package test

import (
	"os"
	"rebitcask"
	"rebitcask/internal/setting"
	"testing"
	"time"
)

func setup() {
	rebitcask.Setup(".env.bench")
}

func teardown() {
	time.Sleep(time.Second * 100)
	os.RemoveAll(setting.Config.DATA_FOLDER_PATH)
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}

func BenchmarkEmptyGet_100(b *testing.B) {
	keys, _ := GenerateLowDuplicateRandom(100)
	benchmarkEmptyGet(keys, b)
}

func BenchmarkEmptyGet_1000(b *testing.B) {
	keys, _ := GenerateLowDuplicateRandom(1000)
	benchmarkEmptyGet(keys, b)
}

func BenchmarkEmptyGet_10000(b *testing.B) {
	keys, _ := GenerateLowDuplicateRandom(10000)
	benchmarkEmptyGet(keys, b)
}

func benchmarkEmptyGet(keys []string, b *testing.B) {
	for i := 0; i < b.N; i++ {
		func() {
			for _, k := range keys {
				_, _ = rebitcask.Get(k)
			}
		}()
	}
}

func BenchmarkSet_100(b *testing.B) {
	keys, vals := GenerateLowDuplicateRandom(100)
	for i := 0; i < b.N; i++ {
		benchmarkSet(keys, vals, b)
	}
}

func BenchmarkSet_1000(b *testing.B) {
	keys, vals := GenerateLowDuplicateRandom(1000)
	for i := 0; i < b.N; i++ {
		benchmarkSet(keys, vals, b)
	}
}

func benchmarkSet(keys, vals []string, b *testing.B) {
	for i := 0; i < b.N; i++ {
		func() {
			for i, k := range keys {
				_ = rebitcask.Set(k, vals[i])
			}
		}()
	}
}

func BenchmarkDelete_100(b *testing.B) {
	keys, vals := GenerateLowDuplicateRandom(100)
	for i := 0; i < b.N; i++ {
		benchmarkDelete(keys, vals, b)
	}
}

func BenchmarkDelete_1000(b *testing.B) {
	keys, vals := GenerateLowDuplicateRandom(1000)
	for i := 0; i < b.N; i++ {
		benchmarkDelete(keys, vals, b)
	}
}

func benchmarkDelete(keys, vals []string, b *testing.B) {
	for i := 0; i < b.N; i++ {
		func() {
			for _, k := range keys {
				_ = rebitcask.Delete(k)
			}
		}()
	}
}
