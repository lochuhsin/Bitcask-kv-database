package raft

import "sync"

var (
	raftNode RaftNode
	raftOnce sync.Once
)

func InitializeRaftNode() {
	raftOnce.Do(func() {
		node := NewRaftNode()
		go node.ElectionTimeoutListener()
		raftNode = node
	})

}

func GetRaftNode() RaftNode {
	return raftNode
}
