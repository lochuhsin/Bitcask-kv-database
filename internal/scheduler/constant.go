package scheduler

type tStatus string

const (
	PROCESSING tStatus = "processing"
	FINISHED   tStatus = "finished"
	FALIED     tStatus = "failed"
)
