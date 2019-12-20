package preparer

import (
	"math/rand"
	"time"

	"ck/models"
	"ck/models/pq"
	"ck/models/queue"
	"ck/util"
)

type Preparer struct {
	pqs map[string]pq.PQ
	oq  queue.Queue
}

func NewPreparer(pqs map[string]pq.PQ, oq queue.Queue) *Preparer {
	return &Preparer{pqs, oq}
}

func (p *Preparer) Process() {
	v := p.oq.Pop()
	if v == nil {
		time.Sleep(time.Duration(5) * time.Second)
		return
	}
	go p.Preparing(v.(*models.Order))
}

func (p *Preparer) Preparing(order *models.Order) {
	time.Sleep(time.Duration(rand.Intn(10)+1) * time.Second)

	order.OnShelfTS = time.Now()
	order.CurrentShelfLife = order.ShelfLife
	order.Value = 1.0
	target := p.pqs[order.Temp]
	if target.Full() {
		target = p.pqs["overflow"]
		if target.Full() {
			util.DisplayEvent("discarded")
			util.DisplayOrder(order)
			util.DisplayShelves(p.pqs)
			return
		}
	}
	target.Push(&pq.Item{Value: order, Priority: float32(1.0 - order.Value)})
	util.DisplayEvent("prepared")
	util.DisplayOrder(order)
	util.DisplayShelves(p.pqs)
}
