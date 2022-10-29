package test

import (
	"math/rand"
	"time"
)

var LETTERS = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func generateLowDuplicateRandomData() ([]string, []string) {
	rand.Seed(time.Now().UnixNano())
	const size = 100000
	keys := make([]string, size)
	vals := make([]string, size)
	count := 0
	for i := 0; i < size; i++ {
		k := make([]rune, 30)
		for j, _ := range k {
			k[j] = LETTERS[rand.Intn(len(LETTERS))]
		}
		v := make([]rune, rand.Intn(30))
		for m, _ := range v {
			v[m] = LETTERS[rand.Intn(len(LETTERS))]
		}
		keys = append(keys, string(k))
		vals = append(vals, string(v))
		count += 1
	}
	return keys, vals
}

func generateHighDuplicateRandom() ([]string, []string) {
	rand.Seed(time.Now().UnixNano())
	const size = 100000
	keys := make([]string, size)
	vals := make([]string, size)
	for i := 0; i < size; i++ {
		k := make([]rune, 2)
		for i := range k {
			k[i] = LETTERS[rand.Intn(len(LETTERS))]
		}
		v := make([]rune, rand.Intn(30))
		for i := range v {
			v[i] = LETTERS[rand.Intn(len(LETTERS))]
		}

		keys = append(keys, string(k))
		vals = append(vals, string(v))
	}
	return keys, vals
}
