package common

type Building struct {
	Id             string         `json:"db_letter"`
	Image          string         `json:"image"`
	Images         []string       `json:"images"`
	Name           string         `json:"name"`
	Category       string         `json:"category"`
	Cost           int64          `json:"cost"`
	RobotsNeeded   float64        `json:"robotsNeeded"`
	RealmAvailable bool           `json:"realmAvailable"`
	DoesProduce    []ResourceBase `json:"doesProduce"`
	DoesSell       []ResourceBase `json:"doesSell"`
	CostUnits      float64        `json:"costUnits"`
	Hours          float64        `json:"hours"`
	Wages          float64        `json:"wages"`
}
