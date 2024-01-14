package segment

import (
	"sync"
)

var (
	SegManager     *Manager
	segManagerOnce sync.Once
)

func SegmentInit() {
	segManagerOnce.Do(func() {
		if SegManager == nil {
			SegManager = NewSegmentManager()
		}
	})
}
