package common

type RetailRecord struct {
	AveragePrice         *float64 `json:"averagePrice,omitempty"`
	AmountSoldRestaurant *int64   `json:"amountSoldRestaurant,omitempty"`
	Demand               float64  `json:"demand"`
	Date                 string   `json:"date"`
	Label                string   `json:"label"`
	AmountSold           *int64   `json:"amountSold,omitempty"`
}
