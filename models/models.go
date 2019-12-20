package models

import "time"

type (
	Orders []*Order
	Order  struct {
		Name             string `json:"name"`
		Temp             string `json:"temp"`
		ShelfLife        int    `json:"shelfLife"`
		CurrentShelfLife int
		DecayRate        float32 `json:"decayRate"`
		Value            float32
		OnShelfTS        time.Time
	}
)
