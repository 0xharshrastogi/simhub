package common

type ProductionInfo struct {
	Inputs    []ResourceInput
	LaborCost float64
	AdminCost float64
}

func (info ProductionInfo) TotalInputCost() (sum float64) {
	for _, input := range info.Inputs {
		sum += input.Total()
	}
	return
}

func (info ProductionInfo) TotalCost() (sum float64) {
	return info.TotalInputCost() + info.LaborCost + info.AdminCost
}
