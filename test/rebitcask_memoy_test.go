package test

import (
	"fmt"
	"rebitcask/internal/models"
	"testing"
)

// for debug usage
func TestBinarySearchTreeSmall(t *testing.T) {

	var bst models.BinarySearchTree

	for _, val := range []string{"a", "b", "c", "d"} {
		bst.Set(val, []byte(val))
		v, _ := bst.Get(val)
		fmt.Println(string(v))
	}
}

func TestBinarySearchTreePureRandom(t *testing.T) {
	testTimer(func(t *testing.T) {
		var bst models.BinarySearchTree
		keys, vals := generateLowDuplicateRandomData()

		for i, k := range keys {
			bst.Set(k, []byte(vals[i]))
		}
		fmt.Println("done set")
		for i, k := range keys {
			res, _ := bst.Get(k)

			if string(res) != vals[i] {
				t.Error("Get value error")
			}
		}
		fmt.Printf("tree size: %v\n", bst.GetSize())
	}, t)
}

func TestBinarySearchTreeDuplicateKey(t *testing.T) {
	testTimer(func(t *testing.T) {
		var bst models.BinarySearchTree
		keys, vals := generateHighDuplicateRandom()
		kvMap := make(map[string]string)
		for i, _ := range keys {
			bst.Set("a", []byte(vals[i]))
			kvMap["a"] = vals[i]
		}
		fmt.Println("done set")
		for k, v := range kvMap {
			res, _ := bst.Get(k)
			fmt.Println(res)
			if string(res) != v {
				t.Error("Get value error")
			}
		}
		fmt.Printf("tree size: %v\n", bst.GetSize())
	}, t)
}

func TestTemp(t *testing.T) {
	tmpArr := make(map[int]int)

	for i := 1; i <= 50; i++ {
		tmpArr[i] = i
		if len(tmpArr) != i {
			println(len(tmpArr))
		}
	}

	fmt.Println(len(tmpArr))
	fmt.Println(tmpArr)
}

func TestTemp2(t *testing.T) {
	tmpArr := make([]int, 0)

	for i := 1; i <= 50; i++ {
		tmpArr = append(tmpArr, i)
		if len(tmpArr) != i {
			fmt.Println(len(tmpArr))
		}
	}
	fmt.Println(len(tmpArr))
	fmt.Println(tmpArr)
}
