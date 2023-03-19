package common

import (
	"fmt"
	"strconv"
)

type Realm struct {
	Encyclopedia Encyclopedia
	Market       Market
}

// gets the production cost of a given resource
func (r *Realm) GetProductionCost(res *Resource, admin float64, quality uint) (*ProductionInfo, error) {
	var (
		p        = &ProductionInfo{}
		q        = uint(0)
		infoChan = make(chan ResourceInput, len(res.ProducedFrom))
	)

	if quality > 0 {
		q = quality - 1
	}

	done := make(chan struct{})
	go func() {
		for range res.ProducedFrom {
			info, open := <-infoChan
			if !open {
				return
			}
			p.Inputs = append(p.Inputs, info)
		}
		done <- struct{}{}
	}()

	quit := make(chan struct{})
	isError := false
	for _, pf := range res.ProducedFrom {
		if isError {
			fmt.Println("Breaking out")
			break
		}

		go func(from ProducedFrom) {
			records, err := r.Market.GetMarketInfo(strconv.Itoa(int(from.Resource.Id)))
			if err != nil {
				isError = true
				quit <- struct{}{}
				return
			}

			var sr MarketRecord
			for _, record := range records {
				if uint(record.Quality) == q {
					sr = record
					break
				}
			}

			infoChan <- ResourceInput{
				ProducedFrom: from,
				Price:        sr.Price,
				Quality:      q,
			}
		}(pf)
	}

	p.LaborCost = res.BaseSalary / res.ProducedAnHour
	p.AdminCost = p.LaborCost * admin / 100

	select {
	case <-quit:
		close(infoChan)
		return nil, fmt.Errorf("failed to calculate production cost")
	case <-done:
		return p, nil
	}
}
