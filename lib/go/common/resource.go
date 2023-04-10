package common

type Resource struct {
	Name                  string         `json:"name"`
	Image                 string         `json:"image"`
	Id                    int64          `json:"db_letter"`
	Transportation        float64        `json:"transportation"`
	IsRetailAble          bool           `json:"retainable"`
	Research              bool           `json:"research"`
	ExchangeTradable      bool           `json:"exchangeTradable"`
	RealmAvailable        bool           `json:"realmAvailable"`
	ProducedFrom          []ProducedFrom `json:"producedFrom"`
	SoldAt                string         `json:"soldAt"`
	SoldAtRestaurant      string         `json:"soldAtRestaurant"`
	ProducedAt            string         `json:"producedAt"`
	NeededFor             []ResourceBase `json:"neededFor"`
	TransportNeeded       float64        `json:"transportNeeded"`
	ProducedAnHour        float64        `json:"producedAnHour"`
	BaseSalary            float64        `json:"baseSalary"`
	AverageRetailPrice    *float64       `json:"averageRetailPrice"`
	MarketSaturation      *float64       `json:"marketSaturation"`
	MarketSaturationLabel *string        `json:"marketSaturationLabel"`
	RetailModeling        *string        `json:"retailModeling"`
	StoreBaseSalary       *float64       `json:"storeBaseSalary"`
	RetailData            []RetailRecord `json:"retailData"`
	ImprovesQualityOf     []ResourceBase `json:"improvesQualityOf"`
}
