package models

type networkType string

const (
	ETHEREUM networkType = "ETHEREUM"
	STELLAR  networkType = "STELLAR"
)

// type User struct {
// 	UserName string `json:"userName" gorm:"primaryKey"`
// 	Password string `json:"password"`
// 	Address  string `json:"address" gorm:"uniqueIndex"`
// 	Network  string `json:"network"`
// }

type User struct {
	UserName string `json:"userName" gorm:"primaryKey"`
	Password string `json:"password"`
	Address  string `json:"address" gorm:"uniqueIndex"`
	Network  string `json:"network" gorm:"type:networkType"`
}
