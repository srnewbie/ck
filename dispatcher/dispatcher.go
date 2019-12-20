package dispatcher

import (
	"math/rand"
	"time"

	"ck/dispatcher/deliverer"
	"ck/dispatcher/preparer"
	"ck/dispatcher/promoter"
	"ck/models"
	"ck/models/cron"
	"ck/models/pq"
	"ck/models/queue"
	"ck/util"
)

type Dispatcher struct {
	cron      cron.Cron
	oq        queue.Queue
	pqs       map[string]pq.PQ
	preparer  *preparer.Preparer
	deliverer *deliverer.Deliverer
	promoter  *promoter.Promoter
}

func New() *Dispatcher {
	rand.Seed(time.Now().UnixNano())
	pqs := map[string]pq.PQ{
		"hot":      pq.New(15),
		"cold":     pq.New(15),
		"frozen":   pq.New(15),
		"overflow": pq.New(20),
	}
	oq := queue.New()
	preparer := preparer.NewPreparer(pqs, oq)
	deliverer := deliverer.NewDeliverer(pqs)
	promoter := promoter.NewPromoter(pqs)
	return &Dispatcher{
		oq:  oq,
		pqs: pqs,
		cron: cron.New([]*cron.Config{
			&cron.Config{
				Name:     "preparer",
				Interval: 0,
				Cb:       preparer.Process,
			},
			&cron.Config{
				Name:     "deliver",
				Interval: 0,
				Cb:       deliverer.Process,
			},
			&cron.Config{
				Name:     "promoter",
				Interval: 1,
				Cb:       promoter.Process,
			},
		}),
		preparer:  preparer,
		deliverer: deliverer,
		promoter:  promoter,
	}
}

func (d *Dispatcher) Start() {
	d.cron.Start()
}

func (d *Dispatcher) Push(order *models.Order) {
	d.oq.Push(order)
	util.DisplayEvent("accepted")
	util.DisplayOrder(order)
	util.DisplayShelves(d.pqs)
}
