package deliverer

import (
	"ck/models"
	"ck/models/pq"
	"fmt"
	"math/rand"
	"time"
)

type Deliverer struct {
	pq pq.PQ
}

func NewDeliverer(pq pq.PQ) *Deliverer {
	return &Deliverer{pq}
}

func (d *Deliverer) Process() {
	if d.pq.Len() > 0 {
		rand.Seed(time.Now().UnixNano())
		time.Sleep(time.Duration(rand.Intn(10)) * time.Second)

		item := d.pq.Pop()
		if item != nil {
			order := item.(*pq.Item).Value.(*models.Order)
			fmt.Println(fmt.Sprintf("deliver arrived, picked order: %d (%d)", order.ID, item.(*pq.Item).Priority))
		}
	}
}
