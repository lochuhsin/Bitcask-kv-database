package transaction

import "sync"

var cLogger *CommitLogger
var cOnce sync.Once

func InitCommitLogger() {
	cOnce.Do(
		func() {
			if cLogger == nil {
				logger := NewCommitLogger()
				cLogger = &logger
			}
		})
}

func GetCommitLogger() *CommitLogger {
	return cLogger
}
