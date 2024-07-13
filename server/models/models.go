package models

type networkType string

type User struct {
	UserName string `json:"userName" gorm:"primaryKey"`
	Password string `json:"password"`
	Address  string `json:"address" gorm:"uniqueIndex"`
	Network  string `json:"network"`
}
