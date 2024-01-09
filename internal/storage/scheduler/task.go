package scheduler

import (
	"errors"
	"rebitcask/internal/storage/memory"
	"sync"
)

type taskId string

type task struct {
	timestamp int64
	id        taskId         // uuid
	m         memory.IMemory // should be cloned memory
}

func (t *task) GetMemory() memory.IMemory {
	return t.m
}

type taskPool struct {
	queue []task
	mu    sync.Mutex // Guards the Add/Pop operation
}

func (t *taskPool) Add(task task) {
	t.mu.Lock()
	t.queue = append(t.queue, task)
	t.mu.Unlock()
}

func (t *taskPool) Pop() (task, bool) {
	t.mu.Lock()
	defer t.mu.Unlock()
	if len(t.queue) == 0 {
		return task{}, false
	}
	val := t.queue[0]
	t.queue = t.queue[1:]
	return val, true
}

// Note: This function returns in inverse order of tasks
// from lastest task to oldest task
// TODO: Might have performance drop ... figure out best practice
func (t *taskPool) GetWaitingTasks() []task {
	t.mu.Lock()

	tasks := make([]task, len(t.queue))

	for i, tk := range t.queue {
		tasks[len(t.queue)-1-i] = task{
			timestamp: tk.timestamp,
			id:        tk.id,
			m:         tk.m,
		}
	}
	t.mu.Unlock()
	return tasks
}

type linkedlist struct {
	t    *task
	next *linkedlist
	prev *linkedlist
}

type TaskOrderedMap struct {
	taskMap map[taskId]*linkedlist
	top     *linkedlist // use sentinel node to implement this
	bottom  *linkedlist // use sentinel node to implement this
	mu      sync.Mutex
}

func InitTaskOrderedMap() *TaskOrderedMap {
	top := &linkedlist{}
	bottom := &linkedlist{}
	top.next = bottom
	bottom.prev = top
	return &TaskOrderedMap{taskMap: map[taskId]*linkedlist{}, top: top, bottom: bottom}
}

func (t *TaskOrderedMap) Set(id taskId, task task) {
	t.mu.Lock()
	newNode := linkedlist{
		&task, nil, nil,
	}
	t.bottom.prev.next = &newNode
	newNode.prev = t.bottom.prev
	newNode.next = t.bottom
	t.bottom.prev = &newNode

	t.taskMap[id] = &newNode
	t.mu.Unlock()
}

func (t *TaskOrderedMap) Delete(id taskId) error {
	t.mu.Lock()
	defer t.mu.Unlock()
	if len(t.taskMap) == 0 {
		return errors.New("deleting empty taskOrderedMap is not allowed")
	}
	if _, ok := t.taskMap[id]; !ok {
		panic("task id not found in ordered map, data is missing")
	}
	node := t.taskMap[id]
	node.prev.next = node.next
	node.next.prev = node.prev

	delete(t.taskMap, id)

	return nil
}

// Note: This function returns in inverse order of tasks
// from lastest task to oldest task
func (t *TaskOrderedMap) GetByOrder() []task {
	/**
	 * doesn't allow dirty read in this area
	 * since i'm doing sort of snapshot mechanism here.
	 * and also, I'm assuming the number of tasks should be reasonable
	 * small
	 */
	t.mu.Lock()
	defer t.mu.Unlock()

	if len(t.taskMap) == 0 {
		return []task{}
	}

	arr := make([]task, len(t.taskMap))

	// loop backwards, from latest to oldest task
	// and skip the first sentinel node
	rPtr := t.bottom.prev
	i := 0

	// the second condition stops when it reaches the
	// top node, which is also a sentinel node
	for rPtr != nil && rPtr.t != nil {
		arr[i] = *rPtr.t
		i++
		rPtr = rPtr.prev
	}
	return arr
}
