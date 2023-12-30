package service

import (
	"rebitcask/internal/storage/dao"
	"rebitcask/internal/storage/segment"
)

func SGet(k dao.NilString) (val dao.Base, status bool) {
	/**
	 * Get function always return two values
	 * 1. data
	 * 2. status which indicates whether the key exists or not
	 */
	return segment.SegManager.Get(k)
}
