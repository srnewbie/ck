package dispatcher

import (
	"ck/models"
	"ck/models/pq"
	"ck/models/queue"
	"fmt"
	"time"
)

type Preparer struct {
	pqs *[]pq.PQ
	oq  queue.Queue
}

func (p *Preparer) Process() {
	v := p.oq.Pop()
	if v != nil {
		order := v.(*models.Order)

		fmt.Println(fmt.Sprintf("preparing order: %d (%d seconds)", order.ID, order.PrepareTime))

		time.Sleep(time.Duration(order.PrepareTime) * time.Second)

		fmt.Println(fmt.Sprintf("finished order: %d", order.ID))

		idx := computePQIndex(order)
		h := (*p.pqs)[idx]
		priority := computePriority(order)
		h.Push(&pq.Item{
			Value:    order,
			Priority: priority,
		})

		str := "pq:"
		for _, i := range *p.pqs {
			str = fmt.Sprintf("%s %d", str, i.Len())
		}
		fmt.Println(str)
	} else {
		fmt.Println("No order to prepare")

		for i, d := range *p.pqs {
			str := fmt.Sprintf("%d:", i)
			for d.Len() > 0 {
				item := d.Pop().(*pq.Item)
				str = fmt.Sprintf("%s %d (%d)", str, item.Value.(*models.Order).ID, item.Priority)
			}
			fmt.Println(str)
		}

		time.Sleep(time.Duration(5) * time.Second)
	}
}
