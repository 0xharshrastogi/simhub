package common

type BuildingBase struct {
	DBLetter       string           `json:"db_letter"`
	Image          string           `json:"image"`
	Images         []string         `json:"images"`
	Name           string           `json:"name"`
	Category       BuildingCategory `json:"category"`
	Cost           int64            `json:"cost"`
	RobotsNeeded   float64          `json:"robotsNeeded"`
	RealmAvailable bool             `json:"realmAvailable"`
}
