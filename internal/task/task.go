package task

import (
	"errors"
	"rebitcask/internal/memory"
	"sync"
)

type TaskId string

type Task struct {
	Timestamp int64
	Id        TaskId         // uuid
	M         memory.IMemory // should be cloned memory
}

type node struct {
	t    *Task
	next *node
	prev *node
}

type taskPool struct {
	/**
	 * TODO: ----------------------------------------------------------------
	 * Note: taskMap could be optimized to fixed size array (or ring queue)
	 * by using task id as uint32 pointing to the linked list node
	 * just to make sure the integer (index) should always be unique
	 * during concurrency. i.e atomic counter ...etc
	 */
	taskMap map[TaskId]*node
	top     *node
	bottom  *node
	sync.Mutex
}

func NewTaskPool() *taskPool {
	/**
	 * I'm using sentinel node to implement ordered map
	 * as it is simpler to handle edge case (i.e empty)
	 */
	top := &node{}
	bottom := &node{}
	top.next = bottom
	bottom.prev = top
	return &taskPool{taskMap: map[TaskId]*node{}, top: top, bottom: bottom}
}

func (t *taskPool) Get(id TaskId) (*Task, bool) {
	t.Lock()
	tk, ok := t.taskMap[id]
	t.Unlock()
	return tk.t, ok
}

func (t *taskPool) Set(id TaskId, task Task) {
	t.Lock()
	newNode := node{
		&task, nil, nil,
	}
	t.bottom.prev.next = &newNode
	newNode.prev = t.bottom.prev
	newNode.next = t.bottom
	t.bottom.prev = &newNode

	t.taskMap[id] = &newNode
	t.Unlock()
}

func (t *taskPool) Delete(id TaskId) error {
	t.Lock()
	defer t.Unlock()
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
func (t *taskPool) GetByOrder() []Task {
	t.Lock()
	defer t.Unlock()

	if len(t.taskMap) == 0 {
		return []Task{}
	}

	arr := make([]Task, len(t.taskMap))

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
