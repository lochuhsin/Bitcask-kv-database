package raft

import (
	"sync"
	"time"
)

type nodeState struct {
	state State
}

type RaftNode struct {
	state            nodeState
	electionTimeout  time.Duration
	electionTimeoutQ chan any
	stateMu          *sync.RWMutex
}

func NewRaftNode() RaftNode {
	return RaftNode{
		state:            nodeState{FOLLOWER},
		electionTimeout:  ELECTION_TIMEOUT,
		electionTimeoutQ: make(chan any),
		stateMu:          &sync.RWMutex{},
	}
}

func (r *RaftNode) Run() {
	for {
		switch r.getState() {
		case FOLLOWER:
		case LEADER:
		case CANDIDATE:
		default:
			panic("invalid node state")
		}
	}
}

func (r *RaftNode) getState() State {
	return r.state.state
}
