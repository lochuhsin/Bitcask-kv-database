package test

import (
	"fmt"
	"os"
	"rebitcask/internal/storage"
	"testing"
	"time"
)

func TestPureRandomWriteLoad(t *testing.T) {
	// TODO: Implement write queue and frozen memory
	// inorder to improve write performance
	// probably need coroutines
	// this load is too heavy
	storage.Init()
	dataCount := 20000000
	keys, vals := generateHugeLowDuplicateRandomData(dataCount)
	s := time.Now()
	fmt.Println(s)
	for i, k := range keys {
		err := storage.Set(k, vals[i])

		if err != nil {
			t.Error("Something went wrong while setting")
		}
	}
	fmt.Println("done set")
	timeLength := time.Since(s)
	fmt.Println("test finished")
	fmt.Printf("Cost: %v", timeLength)
	defer RemoveSegment()
}

func TestGetSetPureRandom(t *testing.T) {
	storage.Init()
	keys, vals := generateLowDuplicateRandomData()
	s := time.Now()
	fmt.Println(s)
	for i, k := range keys {
		err := storage.Set(k, vals[i])

		if err != nil {
			t.Error("Something went wrong while setting")
		}
	}
	fmt.Println("done set")
	for i, k := range keys {
		res, status := storage.Get(k)

		if res != vals[i] {
			t.Error(res, status)
			panic("")
		}
	}
	timeLength := time.Since(s)
	fmt.Println("test finished")
	fmt.Printf("Cost: %v", timeLength)

}

func TestGetSetKeyDuplicateRandom(t *testing.T) {
	keys, vals := generateHighDuplicateRandom()
	lastValMap := make(map[string]string)

	s := time.Now()
	fmt.Println(s)
	for i, k := range keys {
		err := storage.Set(k, vals[i])

		if err != nil {
			t.Error("Something went wrong while setting")
		}
		lastValMap[k] = vals[i]
	}
	fmt.Println("done set")
	for k, v := range lastValMap {
		res, _ := storage.Get(k)

		if res != v {
			t.Error("Get value error")
		}
	}
	timeLength := time.Since(s)
	fmt.Println("test finished")
	fmt.Printf("Cost: %v", timeLength)
}

func TestGetSetFixValue(t *testing.T) {
	keys, _ := generateLowDuplicateRandomData()
	s := time.Now()
	fmt.Println(s)
	for _, k := range keys {
		err := storage.Set(k, "@@@@@@@")

		if err != nil {
			t.Fatal("Something went wrong while setting")
		}

		res, _ := storage.Get(k)

		if res != "@@@@@@@" {
			t.Error("Get value error")
		}
	}
	timeLength := time.Since(s)
	fmt.Println("test finished")
	fmt.Printf("Cost: %v", timeLength)

}

func TestGetSetFixKey(t *testing.T) {
	_, vals := generateLowDuplicateRandomData()
	s := time.Now()
	fmt.Println(s)

	sameKey := "Same"
	var lastVal string
	for _, v := range vals {
		err := storage.Set(sameKey, v)
		if err != nil {
			t.Fatal("Something went wrong while setting")
		}
		lastVal = v
	}
	if res, _ := storage.Get(sameKey); res != lastVal {
		t.Fatal("final assertion error")
	}
}

func TestDelete(t *testing.T) {
	s := time.Now()
	keys, vals := generateLowDuplicateRandomData()

	for i, k := range keys {
		err := storage.Set(k, vals[i])
		if err != nil {
			t.Fatal("Something went wrong while setting")
		}

		err = storage.Delete(k)
		if err != nil {
			t.Fatal("Something went wrong while deleting")
		}
		val, status := storage.Get(k)
		if status != false {
			fmt.Println(vals[i], val)
			t.Error("Get value error")
		}
	}
	timeLength := time.Since(s)
	fmt.Println("test finished")
	fmt.Printf("Cost: %v", timeLength)
}

// for debug purpose
func TestSmallVal(t *testing.T) {

	s := time.Now()
	for _, val := range []string{"a", "b", "c", "d"} {
		err := storage.Set(val, val)
		if err != nil {
			t.Fatal("Set function failed")
		}
	}

	res, _ := storage.Get("a")
	if res != "a" {
		fmt.Println(res)
		t.Error("")
	}

	res, _ = storage.Get("b")
	if res != "b" {
		fmt.Println(res)
		t.Error("")
	}
	timeLength := time.Since(s)
	fmt.Println("test finished")
	fmt.Printf("Cost: %v", timeLength)
}

// used for removing all log files after tests are done
func TestRemoveLog(t *testing.T) {
	_ = os.RemoveAll("./log/")
	fmt.Println("success remove")
}
