package dispatcher

import (
	"fmt"

	"ck/dispatcher/deliverer"
	"ck/dispatcher/preparer"
	"ck/models"
	"ck/models/cron"
	"ck/models/pq"
	"ck/models/queue"
)

type Dispatcher struct {
	cron      cron.Cron
	oq        queue.Queue
	preparer  *preparer.Preparer
	deliverer *deliverer.Deliverer
}

func New() *Dispatcher {
	pq := pq.New()
	oq := queue.New()
	preparer := preparer.NewPreparer(pq, oq)
	deliverer := deliverer.NewDeliverer(pq)
	return &Dispatcher{
		oq: oq,
		cron: cron.New([]*cron.Config{
			&cron.Config{
				Name:     "preparer",
				Interval: 0,
				Cb:       preparer.Process,
			},
			&cron.Config{
				Name:     "deliver",
				Interval: -1,
				Cb:       deliverer.Process,
			},
		}),
		preparer:  preparer,
		deliverer: deliverer,
	}
}

func (d *Dispatcher) Start() {
	d.cron.Start()
}

func (d *Dispatcher) Push(order *models.Order) {
	fmt.Println(fmt.Sprintf("order received: %d", order.ID))
	d.oq.Push(order)
	fmt.Println(fmt.Sprintf("-- order queue size: %d", d.oq.Len()))
}
