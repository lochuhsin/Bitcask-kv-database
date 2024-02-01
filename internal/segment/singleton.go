package segment

import (
	"sync"
)

var (
	SegManager     *Manager
	segManagerOnce sync.Once
)

func InitSegmentManager() {
	segManagerOnce.Do(func() {
		if SegManager == nil {
			manager := NewSegmentManager()
			SegManager = &manager
		}
	})
}

func GetSegmentManager() *Manager {
	return SegManager
}
