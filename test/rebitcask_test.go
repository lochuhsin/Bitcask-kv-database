package rebitcask

import (
	"fmt"
	"math/rand"
	"os"
	"rebitcask/rebitcask"
	"testing"
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
	for i := 0; i < 100000; i++ {
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

func TestGetSetPureRandom(t *testing.T) {
	ans := generatePureRandomData()
	s := time.Now()
	fmt.Println(s)
	for k, v := range ans {
		err := rebitcask.Set(k, v)

		if err != nil {
			t.Error("Something went wrong while setting")
		}
	}
	fmt.Println("done set")
	for k, v := range ans {
		res, _ := rebitcask.Get(k)

		if res != v {
			t.Error("Get value error")
		}
	}
	timeLength := time.Since(s)
	fmt.Println("test finished")
	fmt.Printf("Cost: %v", timeLength)

}

func TestGetSetKeyDuplicateRandom(t *testing.T) {
	ans := generateKeyDuplicateRandomData()
	s := time.Now()
	fmt.Println(s)
	for k, v := range ans {
		err := rebitcask.Set(k, v)

		if err != nil {
			t.Error("Something went wrong while setting")
		}
	}
	fmt.Println("done set")
	for k, v := range ans {
		res, _ := rebitcask.Get(k)

		if res != v {
			t.Error("Get value error")
		}
	}
	timeLength := time.Since(s)
	fmt.Println("test finished")
	fmt.Printf("Cost: %v", timeLength)
}

func TestGetSetFixValue(t *testing.T) {
	ans := generatePureRandomData()
	s := time.Now()
	fmt.Println(s)
	for k, _ := range ans {
		err := rebitcask.Set(k, "@@@@@@@")

		if err != nil {
			t.Fatal("Something went wrong while setting")
		}

		res, _ := rebitcask.Get(k)

		if res != "@@@@@@@" {
			t.Error("Get value error")
		}
	}
	timeLength := time.Since(s)
	fmt.Println("test finished")
	fmt.Printf("Cost: %v", timeLength)

}

func TestGetSetFixKey(t *testing.T) {
	ans := generatePureRandomData()
	s := time.Now()
	fmt.Println(s)

	sameKey := "Same"
	var lastVal string
	for _, v := range ans {
		err := rebitcask.Set(sameKey, v)
		if err != nil {
			t.Fatal("Something went wrong while setting")
		}
		lastVal = v
	}
	if res, _ := rebitcask.Get(sameKey); res != lastVal {
		t.Fatal("final assertion error")
	}
}

func TestDelete(t *testing.T) {
	s := time.Now()
	ans := generatePureRandomData()

	for k, v := range ans {
		err := rebitcask.Set(k, v)
		if err != nil {
			t.Fatal("Something went wrong while setting")
		}

		err = rebitcask.Delete(k)
		if err != nil {
			t.Fatal("Something went wrong while deleting")
		}
		val, status := rebitcask.Get(k)
		if status != false {
			fmt.Println(v, val)
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
		err := rebitcask.Set(val, val)
		if err != nil {
			t.Fatal("Set function failed")
		}
	}

	res, _ := rebitcask.Get("a")
	if res != "a" {
		fmt.Println(res)
		t.Error("")
	}

	res, _ = rebitcask.Get("b")
	if res != "b" {
		fmt.Println(res)
		t.Error("")
	}
	timeLength := time.Since(s)
	fmt.Println("test finished")
	fmt.Printf("Cost: %v", timeLength)
}

func TestRemoveLog(t *testing.T) {
	os.RemoveAll("./log/")
	fmt.Println("success remove")
}
