package pq

import (
	"container/heap"
	"sync"

	"ck/models"
)

type (
	PQ interface {
		Push(interface{})
		Pop() interface{}
		Len() int
	}
	pq struct {
		items *Items
		lock  sync.RWMutex
	}
	Items []*Item
	Item  struct {
		Value    interface{}
		Priority int
		Index    int
	}
)

func New() PQ {
	return &pq{items: &Items{}}
}

func (pq *pq) Push(x interface{}) {
	pq.lock.Lock()
	defer pq.lock.Unlock()

	heap.Push(pq.items, x)
}

func (pq *pq) Pop() interface{} {
	pq.lock.Lock()
	defer pq.lock.Unlock()

	return heap.Pop(pq.items)
}

func (pq *pq) Len() int {
	return pq.items.Len()
}

func (it Items) Len() int { return len(it) }

func (it Items) Less(i, j int) bool {
	return it[i].Priority > it[j].Priority
}

func (it Items) Swap(i, j int) {
	it[i], it[j] = it[j], it[i]
	it[i].Index = i
	it[j].Index = j
}

func (it *Items) Push(x interface{}) {
	n := len(*it)
	item := x.(*Item)
	item.Index = n
	*it = append(*it, item)
}

func (it *Items) Pop() interface{} {
	old := *it
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.Index = -1
	*it = old[0 : n-1]
	return item
}

func (it *Items) update(item *Item, value *models.Order, priority int) {
	item.Value = value
	item.Priority = priority
	heap.Fix(it, item.Index)
}
