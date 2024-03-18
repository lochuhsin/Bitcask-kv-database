package raft

import (
	"sync"
	"sync/atomic"
	"time"
)

type nodeState struct {
	state   State
	stateMu *sync.RWMutex
	term    atomic.Int32
}

func (n *nodeState) getState() State {
	n.stateMu.RLock()
	defer n.stateMu.RUnlock()
	return n.state
}

func (n *nodeState) setState(state State) {
	n.stateMu.Lock()
	defer n.stateMu.Unlock()
	n.state = state
}

func (n *nodeState) termIncr() {
	n.term.Add(1)
}

func (n *nodeState) getTerm() int32 {
	return n.term.Load()
}

type RaftNode struct {
	state            nodeState
	electionTimeout  time.Duration
	heartbeatTimeout time.Duration
	reqCh            chan CommitLogEntry
	respCh           chan CommitLogEntry
	/**
	 * TODO: Add exit channel to manage goroutines
	 */
}

func NewRaftNode() RaftNode {
	return RaftNode{
		state: nodeState{
			state:   FOLLOWER,
			stateMu: &sync.RWMutex{},
		},
		electionTimeout:  ELECTION_TIMEOUT,
		heartbeatTimeout: HEARTBEAT_TIMEOUT,
		reqCh:            make(chan CommitLogEntry),
		respCh:           make(chan CommitLogEntry),
	}
}

func (r *RaftNode) Run() {
	for {
		switch r.state.getState() {
		case FOLLOWER:
			r.runFollower()
		case LEADER:
			r.runLeader()
		case CANDIDATE:
			r.runCandidate()
		default:
			panic("invalid node state")
		}
	}
}

func (r *RaftNode) runFollower() {
}

func (r *RaftNode) runLeader() {
	/**
	 * AppendEntry, OperationType, Data,
	 *
	 * 1. If there is no incoming AppendEntry request, within certain intervals
	 *    then send type = HearBeat Entry to followers
	 *
	 * 2, If there is an incoming AppendEntry request
	 */
	heartbeatCycle := HEARTBEAT_TIMEOUT / 2
	heartBeatTimer := time.After(heartbeatCycle)
	for r.state.getState() == LEADER {

		select {
		case entry := <-r.reqCh:
			switch entry.Op {
			case WRITE:

			}

		case <-heartBeatTimer:
			r.respCh <- CommitLogEntry{
				Term: r.state.getTerm(),
				Op:   HEARTBEAT,
			}
			heartBeatTimer = time.After(heartbeatCycle)

		}
	}
}

func (r *RaftNode) runCandidate() {}

func (r *RaftNode) GetRespCh() <-chan CommitLogEntry {
	return r.respCh
}

func (r *RaftNode) GetReqCh() chan<- CommitLogEntry {
	return r.reqCh
}
