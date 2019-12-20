package dispatcher

import (
	"ck/models"
	"ck/models/cron"
	"ck/models/pq"
	"ck/models/queue"
	"fmt"
	"math/rand"
	"time"
)

type Dispatcher struct {
	cron     cron.Cron
	oq       queue.Queue
	preparer *Preparer
}

func New() *Dispatcher {
	oq := queue.New()
	pqs := make([]pq.PQ, 10)
	for i, _ := range pqs {
		pqs[i] = pq.New()
	}
	preparer := &Preparer{
		pqs: &pqs,
		oq:  oq,
	}
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
				Interval: 10,
				Cb:       func() { fmt.Println("I am deliver") },
			},
		}),
		preparer: preparer,
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

func computePQIndex(order *models.Order) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(10)
}

func computePriority(order *models.Order) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(10)
}
