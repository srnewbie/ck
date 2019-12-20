package promoter

import (
	"ck/models"
	"ck/models/pq"
	"ck/util"
)

type Promoter struct {
	pqs map[string]pq.PQ
}

func NewPromoter(pqs map[string]pq.PQ) *Promoter {
	return &Promoter{pqs: pqs}
}

func (p *Promoter) Process() {
	multiplier := 1
	for k, pq := range p.pqs {
		if k == "overflow" {
			multiplier = 2
		}
		pq.Adjust(multiplier)
	}

	if !util.FullShelves(p.pqs) {
		if q, ok := p.pqs["overflow"]; ok {
			if !q.Empty() {
				item := q.Pop()
				order := item.(*pq.Item).Value.(*models.Order)
				target := p.pqs[order.Temp]
				if !target.Full() {
					target.Push(item)
				}
			}
		}
	}
}
