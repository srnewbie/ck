package queue

import (
	"container/list"
	"sync"
)

type (
	Queue interface {
		Len() int
		Push(v interface{}) interface{}
		Pop() interface{}
	}
	queue struct {
		list *list.List
		lock sync.RWMutex
	}
)

func New() Queue {
	return &queue{
		list: list.New(),
	}
}

func (q *queue) Len() int {
	q.lock.Lock()
	defer q.lock.Unlock()

	return q.list.Len()
}

func (q *queue) Push(v interface{}) interface{} {
	q.lock.Lock()
	defer q.lock.Unlock()

	return q.list.PushBack(v).Value
}

func (q *queue) Pop() interface{} {
	if q.list.Len() > 0 {
		e := q.list.Front()
		q.list.Remove(e)
		return e.Value
	}
	return nil
}
