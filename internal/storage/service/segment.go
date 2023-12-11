package service

import (
	"rebitcask/internal/storage/dao"
	"rebitcask/internal/storage/memory"
	"rebitcask/internal/storage/segment"
)

var segManager segment.SegmentManager

func SegmentInit() {
	/**
	 * Implement segment specific data structure
	 */
	segManager = segment.InitSegmentManager()

	/**
	 *  TODO: implement reload from segment log files
	 * 1. Segment Collection
	 * 2. Segment Index
	 * 3. Implement transaction log files
	 * */
}

func SGet(k dao.NilString) (dao.Base, bool) {
	return segManager.Get(k)
}

func memoryToSegment(m memory.MemoryBase) error {
	segManager.ConvertToSegment(m)
	return nil
}
