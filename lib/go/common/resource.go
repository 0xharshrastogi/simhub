package common

type Resource struct {
	Name                  string         `json:"name"`
	Image                 string         `json:"image"`
	Id                    int64          `json:"db_letter"`
	Transportation        int64          `json:"transportation"`
	IsRetailAble          bool           `json:"retainable"`
	Research              bool           `json:"research"`
	ExchangeTradable      bool           `json:"exchangeTradable"`
	RealmAvailable        bool           `json:"realmAvailable"`
	ProducedFrom          []ProducedFrom `json:"producedFrom"`
	SoldAt                *BuildingBase  `json:"soldAt"`
	SoldAtRestaurant      *BuildingBase  `json:"soldAtRestaurant"`
	ProducedAt            BuildingBase   `json:"producedAt"`
	NeededFor             []ResourceBase `json:"neededFor"`
	TransportNeeded       int64          `json:"transportNeeded"`
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
