package preparer

import (
	"fmt"
	"math/rand"
	"time"

	"ck/models"
	"ck/models/pq"
	"ck/models/queue"
)

type Preparer struct {
	pq pq.PQ
	oq queue.Queue
}

func NewPreparer(pq pq.PQ, oq queue.Queue) *Preparer {
	return &Preparer{pq, oq}
}

func (p *Preparer) Process() {
	v := p.oq.Pop()
	if v != nil {
		order := v.(*models.Order)

		fmt.Println(fmt.Sprintf("preparing order: %d (%d seconds)", order.ID, order.PrepareTime))

		time.Sleep(time.Duration(order.PrepareTime) * time.Second)

		fmt.Println(fmt.Sprintf("finished order: %d", order.ID))

		priority := computePriority(order)
		p.pq.Push(&pq.Item{
			Value:    order,
			Priority: priority,
		})
		fmt.Println(fmt.Sprintf("-- pq size: %d", p.pq.Len()))
	} else {
		fmt.Println("No order to prepare")
		time.Sleep(time.Duration(5) * time.Second)
	}
}

func computePriority(order *models.Order) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(10)
}
