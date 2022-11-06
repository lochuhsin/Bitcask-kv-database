package test

import (
	"fmt"
	"rebitcask/internal/models"
	"sort"
	"testing"
)

// for debug usage
func TestBinarySearchTreeSmall(t *testing.T) {

	var bst models.BinarySearchTree

	for _, val := range []string{"a", "b", "c", "d"} {

		bst.Set(val, models.Item{
			Val: []byte(val),
		})
		v, _ := bst.Get(val)
		fmt.Println(string(v.Val))
	}
}

func TestBinarySearchTreePureRandom(t *testing.T) {
	testTimer(func(t *testing.T) {
		var bst models.BinarySearchTree
		bst.Init()
		keys, vals := generateLowDuplicateRandomData()

		for i, k := range keys {
			bst.Set(k, models.Item{
				Val: []byte(vals[i]),
			})
		}
		fmt.Println("done set")
		for i, k := range keys {
			res, _ := bst.Get(k)

			if string(res.Val) != vals[i] {
				t.Error("Get value error")
			}
		}
		fmt.Printf("tree size: %v\n", bst.GetSize())
	}, t)
}

func TestBinarySearchTreeDuplicateKey(t *testing.T) {
	testTimer(func(t *testing.T) {
		var bst models.BinarySearchTree
		bst.Init()
		keys, vals := generateHighDuplicateRandom()
		kvMap := make(map[string]string)
		for i, _ := range keys {
			bst.Set("a", models.Item{
				Val: []byte(vals[i]),
			})
			kvMap["a"] = vals[i]
		}
		fmt.Println("done set")
		for k, v := range kvMap {
			res, _ := bst.Get(k)
			fmt.Println(res)
			if string(res.Val) != v {
				t.Error("Get value error")
			}
		}
		fmt.Printf("tree size: %v\n", bst.GetSize())
	}, t)
}

func TestBinarySearchTreeALL(t *testing.T) {
	testTimer(func(t *testing.T) {
		keys, _ := generateLowDuplicateRandomData()
		var bst models.BinarySearchTree
		bst.Init()
		for _, key := range keys {
			bst.Set(key, models.Item{
				Val: []byte(key),
			})
		}

		arr := make([]string, 0, len(keys))
		for _, v := range bst.GetAll() {
			arr = append(arr, v.Key)
		}

		sort.Strings(keys)

		for i := 0; i < len(keys); i++ {
			if arr[i] != keys[i] {
				t.Error("two value is not the same", keys[i])
			}
		}
	}, t)
}

func TestAVLTreePureRandom(t *testing.T) {
	testTimer(func(t *testing.T) {
		var bst models.AVLTree
		bst.Init()
		keys, vals := generateLowDuplicateRandomData()

		for i, k := range keys {
			bst.Set(k, models.Item{
				Val: []byte(vals[i]),
			})
		}
		fmt.Println("done set")
		for i, k := range keys {
			res, _ := bst.Get(k)

			if string(res.Val) != vals[i] {
				t.Error("Get value error")
			}
		}
		fmt.Printf("tree size: %v\n", bst.GetSize())
	}, t)
}

func TestAVLTreeDuplicateKey(t *testing.T) {
	testTimer(func(t *testing.T) {
		var bst models.AVLTree
		bst.Init()
		keys, vals := generateHighDuplicateRandom()
		kvMap := make(map[string]string)
		for i, _ := range keys {
			bst.Set("a", models.Item{
				Val: []byte(vals[i]),
			})
			kvMap["a"] = vals[i]
		}
		fmt.Println("done set")
		for k, v := range kvMap {
			res, _ := bst.Get(k)
			fmt.Println(res)
			if string(res.Val) != v {
				t.Error("Get value error")
			}
		}
		fmt.Printf("tree size: %v\n", bst.GetSize())
	}, t)
}

func TestAVLTreeALL(t *testing.T) {
	testTimer(func(t *testing.T) {
		keys, _ := generateLowDuplicateRandomData()
		var bst models.AVLTree
		bst.Init()
		for _, key := range keys {
			bst.Set(key, models.Item{
				Val: []byte(key),
			})
		}

		arr := make([]string, 0, len(keys))
		for _, v := range bst.GetAll() {
			arr = append(arr, v.Key)
		}

		sort.Strings(keys)

		for i := 0; i < len(keys); i++ {
			if arr[i] != keys[i] {
				t.Error("two value is not the same", keys[i])
			}
		}
	}, t)
}
