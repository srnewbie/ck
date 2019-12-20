package deliverer

import (
	"ck/models"
	"ck/models/pq"
	"ck/util"
	"math/rand"
	"time"
)

var qm map[int]string = map[int]string{
	0: "hot",
	1: "cold",
	2: "frozen",
}

type Deliverer struct {
	pqs map[string]pq.PQ
	idx int
}

func NewDeliverer(pqs map[string]pq.PQ) *Deliverer {
	return &Deliverer{pqs: pqs}
}

func (d *Deliverer) Process() {
	if !util.EmptyShelves(d.pqs) {
		d.WaitCourier()

		idx := util.NextShelfIdx(d.pqs, qm, d.idx)
		item := d.pqs[qm[idx]].Pop()
		if item != nil {
			order := item.(*pq.Item).Value.(*models.Order)
			for order.Value <= 0 {
				util.DisplayEvent("discarded")
				util.DisplayOrder(order)
				util.DisplayShelves(d.pqs)
				item = d.pqs[qm[idx]].Pop()
				order = item.(*pq.Item).Value.(*models.Order)
			}
			util.DisplayEvent("handedoff")
			util.DisplayOrder(order)
			util.DisplayShelves(d.pqs)

			d.idx = d.idx + 1
			if d.idx == 3 {
				d.idx = 0
			}
		}
	}
}

func (d *Deliverer) WaitCourier() {
	time.Sleep(time.Duration(rand.Intn(7)+2) * time.Second)
}
