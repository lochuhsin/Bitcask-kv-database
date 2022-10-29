package test

import (
	"math/rand"
	"time"
)

var LETTERS = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func generatePureRandomData() map[string]string {
	rand.Seed(time.Now().UnixNano())
	ans := make(map[string]string)
	for i := 0; i < 100000; i++ {
		k := make([]rune, rand.Intn(30))
		for i := range k {
			k[i] = LETTERS[rand.Intn(len(LETTERS))]
		}
		v := make([]rune, rand.Intn(30))
		for i := range v {
			v[i] = LETTERS[rand.Intn(len(LETTERS))]
		}
		ans[string(k)] = string(v)
	}
	return ans
}

func generateKeyDuplicateRandomData() map[string]string {
	rand.Seed(time.Now().UnixNano())
	ans := make(map[string]string)
	for i := 0; i < 1000000; i++ {
		k := make([]rune, 2)
		for i := range k {
			k[i] = LETTERS[rand.Intn(len(LETTERS))]
		}
		v := make([]rune, rand.Intn(30))
		for i := range v {
			v[i] = LETTERS[rand.Intn(len(LETTERS))]
		}
		ans[string(k)] = string(v)
	}
	return ans
}
