package qstore

import "sync"

type Queue struct {
	m     sync.RWMutex
	Queue []string
}

type QueueStore struct {
	m     sync.RWMutex
	Store map[string]*Queue
}

var QStore *QueueStore

func CreateNewQueueStore() {
	QStore = &QueueStore{
		Store: make(map[string]*Queue),
	}
}

func (qs *QueueStore) QueueExists(name string) (*Queue, bool) {
	qs.m.Lock()
	defer qs.m.Unlock()

	queue, bool := qs.Store[name]
	return queue, bool
}

func (qs *QueueStore) CreateQueue(name string) {
	qs.m.Lock()
	defer qs.m.Unlock()

	qs.Store[name] = &Queue{
		Queue: make([]string, 0),
	}
}

func (q *Queue) Enqueue(value string) {
	q.m.Lock()
	defer q.m.Unlock()

	q.Queue = append(q.Queue, value)
}

func (q *Queue) Dequeue() string {
	q.m.Lock()
	defer q.m.Unlock()

	if len(q.Queue) == 0 {
		return "null"
	}
	val := q.Queue[0]
	q.Queue = q.Queue[1:]
	return val
}
