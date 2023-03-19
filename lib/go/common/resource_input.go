package common

import "fmt"

type ResourceInput struct {
	Quality uint
	Price   float64
	ProducedFrom
}

// Total returns the multiplication of the given resource price and amount
func (r *ResourceInput) Total() float64 {
	return r.Price * r.Amount
}

func (r *ResourceInput) String() string {
	return fmt.Sprintf("%s [Q%d] %f*%f=$%f", r.Resource.Name, r.Quality, r.Price, r.Amount, r.Total())
}
