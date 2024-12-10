package models

type Candidate struct {
	ID         uint     `json:"id"`
	Name       string   `json:"name"`
	DivisionID uint     `json:"division_id"`
	Division   Division `json:"division" gorm:"foreignKey:DivisionID;references:ID"`
	Image      []byte   `json:"image,omitempty"`
}

type Division struct {
	ID       uint   `json:"id"`
	Division string `json:"division"`
}
