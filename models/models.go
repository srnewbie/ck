package models

type (
	Orders []*Order
	Order  struct {
		ID          int    `json:"id"`
		Name        string `json:"name"`
		Type        string `json:"type"`
		OrderTime   string `json:"order_time"`
		Distance    int    `json:"distance"`
		CodeRate    int    `json:"code_rate"`
		DecayRate   int    `json:"Decay_rate"`
		PrepareTime int    `json:"prepare_time"`
	}
)
