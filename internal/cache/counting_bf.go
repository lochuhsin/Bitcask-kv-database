package cache

import (
	"hash/adler32"
	"sync"
)

/**
 * For current use case, we restrict the entire rebitcask key
 * should be string
 */

type CountingBloomFilter struct {
	hashArr [1000000]int
	mu      *sync.Mutex
}

func InitCBF() *CountingBloomFilter {
	return &CountingBloomFilter{hashArr: [1000000]int{}, mu: &sync.Mutex{}}
}

func (cbf *CountingBloomFilter) Get(s string) bool {
	hashNum1 := cbf.hash1(s)
	hashNum2 := cbf.hash2(s)
	hashNum3 := cbf.hash3(s)

	if cbf.hashArr[hashNum1] > 0 && cbf.hashArr[hashNum2] > 0 && cbf.hashArr[hashNum3] > 0 {
		return true
	}
	return false
}

func (cbf *CountingBloomFilter) Set(s string) {
	hashNum1 := cbf.hash1(s)
	hashNum2 := cbf.hash2(s)
	hashNum3 := cbf.hash3(s)

	cbf.mu.Lock()
	cbf.hashArr[hashNum1]++
	cbf.hashArr[hashNum2]++
	cbf.hashArr[hashNum3]++
	cbf.mu.Unlock()
}

func (cbf *CountingBloomFilter) Delete(s string) bool {
	hashNum1 := cbf.hash1(s)
	hashNum2 := cbf.hash2(s)
	hashNum3 := cbf.hash3(s)

	if cbf.hashArr[hashNum1] > 0 && cbf.hashArr[hashNum2] > 0 && cbf.hashArr[hashNum3] > 0 {
		cbf.hashArr[hashNum1]--
		cbf.hashArr[hashNum2]--
		cbf.hashArr[hashNum3]--
		return true
	}
	return false
}

func (cbf *CountingBloomFilter) hash1(s string) int32 {
	alder := adler32.New()
	alder.Write([]byte(s))
	return cbf.abs(alder.Sum32() / 10000)
}

func (cbf *CountingBloomFilter) hash2(s string) int32 {
	alder := adler32.New()
	alder.Write([]byte(s))
	return cbf.abs(alder.Sum32() / 10100)
}

func (cbf *CountingBloomFilter) hash3(s string) int32 {
	alder := adler32.New()
	alder.Write([]byte(s))
	return cbf.abs(alder.Sum32() / 10101)
}

func (cbf *CountingBloomFilter) abs(val uint32) int32 {
	return int32(val)
}
