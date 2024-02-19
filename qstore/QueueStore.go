package qstore

import (
	"context"
	"log"
	"sync"
	"time"
)

type Queue struct {
	m         sync.Mutex
	Queue     []string
	condition *sync.Cond
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
		Queue:     make([]string, 0),
		condition: sync.NewCond(&sync.Mutex{}),
	}
}

func (q *Queue) Enqueue(value string) {
	q.m.Lock()
	defer q.m.Unlock()

	q.Queue = append(q.Queue, value)
	q.condition.Signal()
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

func (q *Queue) BQPop(qtime int) (string, error) {
	waitOnCond := func(ctx context.Context, cond *sync.Cond, conditionMet func() bool) (string, error) {
		stopf := context.AfterFunc(ctx, func() {
			cond.L.Lock()
			defer cond.L.Unlock()
			cond.Broadcast()
		})
		defer stopf()
		for !conditionMet() {
			cond.Wait()
			if ctx.Err() != nil {
				return "", ctx.Err()
			}
		}

		item := q.Queue[0]
		q.Queue = q.Queue[1:]
		return item, nil
	}

	var wg sync.WaitGroup
	wg.Add(1)
	defer wg.Wait()

	var result string
	var popErr error

	go func() {
		defer wg.Done()

		ctx, cancel := context.WithTimeout(context.Background(), time.Duration(qtime)*time.Second)
		defer cancel()

		q.condition.L.Lock()
		defer q.condition.L.Unlock()

		var err error
		result, err = waitOnCond(ctx, q.condition, func() bool { return len(q.Queue) > 0 })
		popErr = err
		log.Print(result)
	}()

	wg.Wait()
	return result, popErr
}

// stopf := context.AfterFunc(ctx, func() {
// 	q.condition.L.Lock()
// 	defer q.condition.L.Unlock()
// 	q.condition.Broadcast()
// })

// defer stopf()
// for len(q.Queue) == 0 {
// 	q.condition.Wait()
// }
// item := q.Queue[0]
// q.Queue = q.Queue[1:]

// return item, nil
