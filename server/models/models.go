package models

type networkType string

type Player struct {
	UserName string `json:"userName" gorm:"primaryKey;type:varchar(20)"`
	Password string `json:"password" gorm:"type:varchar(50)"`
	Address  string `json:"address" gorm:"uniqueIndex"`
	Network  string `json:"network"`
}
