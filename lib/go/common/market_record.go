package common

type MarketRecord struct {
	Id         int64   `json:"id"`
	ResourceId int64   `json:"kind"`
	Quantity   int64   `json:"quantity"`
	Quality    int64   `json:"quality"`
	Price      float64 `json:"price"`
	Seller     Seller  `json:"seller"`
	Posted     string  `json:"posted"`
	Fees       int64   `json:"fees"`
}
