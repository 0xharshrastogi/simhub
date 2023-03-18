package common

type Market interface {
	// GetMarketInfo returns the Magnates Market information of the specified resource
	GetMarketInfo(id string) (MarketInfo, error)
}
