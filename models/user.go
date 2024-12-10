package models

// User represents the user structure for our application
type User struct {
	ID             int    `json:"id" gorm:"primaryKey"`
	Name           string `json:"name"`
	IdentityNumber string `json:"identity_number" gorm:"column:IdentityNumber"` // Map to database column 'IdentityNumber'
	Password       string `json:"password"`
	RoleID         int    `json:"role_id" gorm:"column:RoleID"`
}
