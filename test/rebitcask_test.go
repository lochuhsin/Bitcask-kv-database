package rebitcask

import (
	"fmt"
	"math/rand"
	"rebitcask/rebitcask"
	"testing"
	"time"
)

var LETTERS = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func generate_data() []string {
	rand.Seed(time.Now().UnixNano())
	var ans []string
	for i := 0; i < 10; i++ {
		b := make([]rune, 5)
		for i := range b {
			b[i] = LETTERS[rand.Intn(len(LETTERS))]
		}
		ans = append(ans, string(b))
	}
	return ans
}

//func TestGetSet(t *testing.T) {
//	s := time.Now()
//	ans := generate_data()
//	for _, v := range ans {
//		err := rebitcask.Set(v, v)
//		if err != nil {
//			t.Fatal("Something went wrong while setting")
//		}
//	}
//
//	for _, v := range ans {
//		res, _ := rebitcask.Get(v)
//
//		if res != v {
//			t.Fatal("Get value error")
//		}
//	}
//	timeLength := time.Since(s)
//	fmt.Println("test finished")
//	fmt.Printf("Cost: %v", timeLength)
//}

func TestDelete(t *testing.T) {
	s := time.Now()
	ans := generate_data()
	for i, v := range ans {
		err := rebitcask.Set(v, v)
		if err != nil {
			t.Fatal("Something went wrong while setting")
		}

		err = rebitcask.Delete(v)
		if err != nil {
			t.Fatal("Something went wrong while deleting")
		}

		val, status := rebitcask.Get(v)
		if status != false {
			fmt.Println(i, val)
		}
	}

	//for _, v := range ans {
	//	val, status := rebitcask.Get(v)
	//
	//	if status != false {
	//		t.Error("Delete not clear", val)
	//		fmt.Println(rebitcask.GetAllInMemory())
	//	}
	//}
	timeLength := time.Since(s)
	fmt.Println("test finished")
	fmt.Printf("Cost: %v", timeLength)
}
