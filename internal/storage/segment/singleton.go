package segment

import (
	"sync"
)

var (
	SegManager     *SegmentManager
	segManagerOnce sync.Once
)

func SegmentInit() {
	segManagerOnce.Do(func() {
		if SegManager == nil {
			SegManager = NewSegmentManager()
		}
	})
}
