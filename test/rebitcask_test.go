package test

import (
	"fmt"
	"os"
	"rebitcask/src"
	"testing"
	"time"
)

func TestGetSetPureRandom(t *testing.T) {
	ans := generatePureRandomData()
	s := time.Now()
	fmt.Println(s)
	for k, v := range ans {
		err := src.Set(k, v)

		if err != nil {
			t.Error("Something went wrong while setting")
		}
	}
	fmt.Println("done set")
	for k, v := range ans {
		res, _ := src.Get(k)

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
		err := src.Set(k, v)

		if err != nil {
			t.Error("Something went wrong while setting")
		}
	}
	fmt.Println("done set")
	for k, v := range ans {
		res, _ := src.Get(k)

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
		err := src.Set(k, "@@@@@@@")

		if err != nil {
			t.Fatal("Something went wrong while setting")
		}

		res, _ := src.Get(k)

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
		err := src.Set(sameKey, v)
		if err != nil {
			t.Fatal("Something went wrong while setting")
		}
		lastVal = v
	}
	if res, _ := src.Get(sameKey); res != lastVal {
		t.Fatal("final assertion error")
	}
}

func TestDelete(t *testing.T) {
	s := time.Now()
	ans := generatePureRandomData()

	for k, v := range ans {
		err := src.Set(k, v)
		if err != nil {
			t.Fatal("Something went wrong while setting")
		}

		err = src.Delete(k)
		if err != nil {
			t.Fatal("Something went wrong while deleting")
		}
		val, status := src.Get(k)
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
		err := src.Set(val, val)
		if err != nil {
			t.Fatal("Set function failed")
		}
	}

	res, _ := src.Get("a")
	if res != "a" {
		fmt.Println(res)
		t.Error("")
	}

	res, _ = src.Get("b")
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
	os.RemoveAll("./log/")
	fmt.Println("success remove")
}
