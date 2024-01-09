package scheduler

type taskStatus string

const (
	PROCESSING taskStatus = "processing"
	FINISHED   taskStatus = "finished"
	FALIED     taskStatus = "failed"
)
