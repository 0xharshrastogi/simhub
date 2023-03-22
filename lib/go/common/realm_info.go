package common

type RealmInfo struct {
	ID                     int     `json:"id"`
	Name                   string  `json:"name"`
	Logo                   string  `json:"logo"`
	Phase                  int     `json:"phase"`
	ResearchLimit          int     `json:"researchLimit"`
	Bonds                  bool    `json:"bonds"`
	GovernmentOrders       bool    `json:"governmentOrders"`
	Executives             bool    `json:"executives"`
	RecreationalBuildings  bool    `json:"recreationalBuildings"`
	Collectibles           bool    `json:"collectibles"`
	Robots                 bool    `json:"robots"`
	Purchases              bool    `json:"purchases"`
	SimboostsExchangeLimit int     `json:"simboostsExchangeLimit"`
	BondsMinInterest       float64 `json:"bondsMinInterest"`
	BondsMaxInterest       int     `json:"bondsMaxInterest"`
}
