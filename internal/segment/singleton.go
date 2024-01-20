package segment

import (
	"sync"
)

var (
	SegManager     *Manager
	segManagerOnce sync.Once
)

func InitSegment() {
	segManagerOnce.Do(func() {
		if SegManager == nil {
			SegManager = NewSegmentManager()
		}
	})
}

func GetSegmentManager() *Manager {
	return SegManager
}
