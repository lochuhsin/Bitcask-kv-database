package main

import (
	"net"
	"sync"
)

type NullableConn interface {
	net.Conn
}
type ConnectionQueue struct {
	mu sync.Mutex
	q  []net.Conn
}

// Insert inserts the item into the queue
func (q *ConnectionQueue) Enqueue(conn net.Conn) {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.q = append(q.q, conn)

}

// Remove removes the oldest element from the queue
func (q *ConnectionQueue) Dequeue() (conn NullableConn, status bool) {
	q.mu.Lock()
	defer q.mu.Unlock()
	if len(q.q) > 0 {
		conn := q.q[0]
		q.q = q.q[1:]
		return conn, false
	}
	return nil, true
}

func (q *ConnectionQueue) GetSize() (size int) {
	return len(q.q)
}

// CreateQueue creates an empty queue with desired capacity
func CreateQueue() ConnectionQueue {
	return ConnectionQueue{
		q: []net.Conn{},
	}
}
