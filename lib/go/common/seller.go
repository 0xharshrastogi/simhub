package common

type Seller struct {
	ID           int64       `json:"id"`
	Company      string      `json:"company"`
	RealmID      int64       `json:"realmId"`
	Logo         string      `json:"logo"`
	Certificates int64       `json:"certificates"`
	ContestWINS  int64       `json:"contest_wins"`
	Npc          bool        `json:"npc"`
	CourseID     interface{} `json:"courseId"`
	IP           string      `json:"ip"`
}
