package util

import (
	"ck/models"
	"ck/models/pq"
	"fmt"
)

var Events map[string]string = map[string]string{
	"accepted":  "Order is accepted",
	"prepared":  "Order is palced on shelf",
	"handedoff": "Order is handed-off to courier",
	"discarded": "Order is discarded as waste",
}

func DisplayEvent(k string) {
	fmt.Println(fmt.Sprintf("\n%s", Events[k]))
}

func DisplayOrder(order *models.Order) {
	fmt.Println(" - order info:")
	fmt.Println(
		fmt.Sprintf(
			"   Name: %s, Temp: %s, ShelfLife: %d, DecayRate: %.2f",
			order.Name,
			order.Temp,
			order.ShelfLife,
			order.DecayRate,
		),
	)
}

func DisplayShelves(pqs map[string]pq.PQ) {
	fmt.Println(" - shelves info:")
	for k, v := range pqs {
		str := fmt.Sprintf("%s: %s", k, v.ToPrint())
		fmt.Println(fmt.Sprintf("   %s", str))
	}
}

func EmptyShelves(pqs map[string]pq.PQ) bool {
	for k, v := range pqs {
		if k == "overflow" {
			continue
		}
		if !v.Empty() {
			return false
		}
	}
	return true
}

func FullShelves(pqs map[string]pq.PQ) bool {
	for k, v := range pqs {
		if k == "overflow" {
			continue
		}
		if !v.Full() {
			return false
		}
	}
	return true
}

func NextShelfIdx(pqs map[string]pq.PQ, qm map[int]string, idx int) int {
	for i := 0; i < 3; i++ {
		if pqs[qm[idx%3]].Empty() {
			idx++
			if idx == 3 {
				idx = 0
			}
			continue
		}
		return idx
	}
	panic("no shelf found")
}
