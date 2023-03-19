package common

type ResourceBase struct {
	Name             string  `json:"name"`
	Image            string  `json:"image"`
	Id               int64   `json:"db_letter"`
	Transportation   float64 `json:"transportation"`
	Retailable       *bool   `json:"retailable"`
	Research         bool    `json:"research"`
	ExchangeTradable bool    `json:"exchangeTradable"`
	RealmAvailable   bool    `json:"realmAvailable"`
}
