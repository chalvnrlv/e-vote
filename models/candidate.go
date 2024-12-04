package models

type Candidate struct {
	ID         uint     `json:"id"`
	Name       string   `json:"name"`
	DivisionID uint     `json:"division_id"`
	Division   Division `gorm:"foreignkey:DivisionID"`
}

type Division struct {
	ID       uint   `json:"id"`
	Division string `json:"division"`
}
