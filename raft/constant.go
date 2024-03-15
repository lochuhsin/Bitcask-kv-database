package raft

import "time"

/**
 * Election timeout is the amount of time a follower waits until
 * becoming a candidate.
 */

const (
	ELECTION_TIMEOUT  time.Duration = time.Second * 2
	HEARTBEAT_TIMEOUT time.Duration = time.Second
)

type State string

const (
	CANDIDATE State = "CANDIDATE"
	FOLLOWER  State = "FOLLOWER"
	LEADER    State = "LEADER"
)
