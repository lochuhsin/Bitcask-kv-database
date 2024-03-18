package raft

import "sync/atomic"

/**
 * In memory commit log, this should be re-implemented
 * to store in disk
 */
type Operation string

const (
	HEARTBEAT Operation = "HEARTBEAT"
	WRITE     Operation = "WRITE"
)

func (o Operation) String() string {
	return string(o)
}

type CommitLogEntry struct {
	Term  int32
	Op    Operation
	Key   string
	Value string
	Index int32
}

type commitLogger struct {
	data  []CommitLogEntry
	index atomic.Int32
}

func newCommitLogger() *commitLogger {
	return &commitLogger{
		data: []CommitLogEntry{},
	}
}

func (c *commitLogger) append(entry CommitLogEntry) {
	c.data = append(c.data, entry)
}

func (c *commitLogger) length() int {
	return len(c.data)
}

func (c *commitLogger) getAll() []CommitLogEntry {
	return c.data
}

func (c *commitLogger) getLast() CommitLogEntry {
	return c.data[len(c.data)]
}
