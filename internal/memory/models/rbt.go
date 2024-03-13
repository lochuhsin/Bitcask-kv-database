package models

import "rebitcask/internal/dao"

type RedBlackTree struct {
}

func InitRedBlackTree() *RedBlackTree {
	panic("Not implemented")
}

func (r *RedBlackTree) Get([]byte) (dao.Base, bool) {
	panic("Not implemented")
}
func (r *RedBlackTree) Set(dao.Entry) {
	panic("Not implemented")
}
func (r *RedBlackTree) GetSize() int {
	panic("Not implemented")
}
func (r *RedBlackTree) GetAll() []dao.Entry {
	panic("Not implemented")
}
