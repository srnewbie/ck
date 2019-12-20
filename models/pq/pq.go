package pq

import (
	"container/heap"
	"fmt"
	"sync"
	"time"

	"ck/models"
)

type (
	PQ interface {
		Push(interface{})
		Pop() interface{}
		Len() int
		Full() bool
		Empty() bool
		Adjust(int)
		ToPrint() string
	}
	pq struct {
		items    *Items
		capacity int
		lock     sync.RWMutex
	}
	Items []*Item
	Item  struct {
		Value    interface{}
		Priority float32
		Index    int
	}
)

func New(c int) PQ {
	return &pq{items: &Items{}, capacity: c}
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

func (pq *pq) Full() bool {
	return pq.Len() == pq.capacity
}

func (pq *pq) Empty() bool {
	return pq.Len() == 0
}

func (pq *pq) ToPrint() string {
	pq.lock.Lock()
	defer pq.lock.Unlock()

	str := ""
	for _, i := range *(pq.items) {
		order := i.Value.(*models.Order)
		str = fmt.Sprintf("%s %s (%.2f)", str, order.Name, order.Value)
	}
	return str
}

func (pq *pq) Adjust(m int) {
	m = m * 5
	pq.lock.Lock()
	defer pq.lock.Unlock()

	items := *(pq.items)
	for i, _ := range items {
		order := items[i].Value.(*models.Order)
		elapsed := float32(time.Since(order.OnShelfTS)) / 1000000000
		order.OnShelfTS = time.Now()
		order.CurrentShelfLife = int((float32(order.CurrentShelfLife) - float32(m)*order.DecayRate*elapsed))
		value := float32(order.CurrentShelfLife) / float32(order.ShelfLife)
		if value < 0 {
			value = 0
		}
		order.Value = value
		pq.items.update(items[i], order, 1.0-value)
	}
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

func (it *Items) update(item *Item, value *models.Order, priority float32) {
	item.Value = value
	item.Priority = priority
	heap.Fix(it, item.Index)
}
