package test

import (
	"math/rand"
	"time"
)

var LETTERS = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890!@#$%^&*()+-[]~.,><?")

func GenerateLowDuplicateRandom(dataCount int) ([]string, []string) {
	keyLen := 50
	valLen := 50
	rand.New(rand.NewSource(time.Now().UnixNano()))
	keys := make([]string, dataCount)
	vals := make([]string, dataCount)
	count := 0
	for i := 0; i < dataCount; i++ {
		k := make([]byte, keyLen)
		for j := range k {
			k[j] = LETTERS[rand.Intn(len(LETTERS))]
		}
		v := make([]byte, 1+rand.Intn(valLen))
		for m := range v {
			v[m] = LETTERS[rand.Intn(len(LETTERS))]
		}
		keys[i] = string(k)
		vals[i] = string(v)
		count += 1
	}
	return keys, vals
}

func GenerateHighDuplicateRandom(dataCount int) ([]string, []string) {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	keys := make([]string, 0, dataCount)
	vals := make([]string, 0, dataCount)
	for i := 0; i < dataCount; i++ {
		k := make([]byte, 2)
		for i := range k {
			k[i] = LETTERS[rand.Intn(len(LETTERS))]
		}
		v := make([]byte, 1+rand.Intn(30))
		for i := range v {
			v[i] = LETTERS[rand.Intn(len(LETTERS))]
		}

		keys = append(keys, string(k))
		vals = append(vals, string(v))
	}
	return keys, vals
}
