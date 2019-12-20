package dispatcher

import (
	"ck/models"
	"ck/models/pq"
	"fmt"
)

type Deliverer struct {
	pqs *[]pq.PQ
}

func (d *Deliverer) Process() {
	for _, pq := range d.pqs {
		if pq.Len() > 0 {
			item := pq.Pop().(*pq.Item).Value.(*models.Order)
			fmt.Println(fmt.Sprintf("preparing order: %d (%d seconds)", order.ID, order.PrepareTime))
			continue
		}
	}
}
