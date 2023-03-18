package common

type ProducedFrom struct {
	Resource ResourceBase `json:"resource"`
	Amount   float64      `json:"amount"`
}
