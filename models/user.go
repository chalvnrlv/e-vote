package models

type User struct {
	ID             uint   `gorm:"primaryKey"`
	Name           string `gorm:"column:Name"`
	IdentityNumber string `gorm:"column:IdentityNumber"`
	Password       string `gorm:"column:Password"`
	RoleID         int    `gorm:"column:RoleID"`
}
