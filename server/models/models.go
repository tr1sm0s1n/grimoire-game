package models

type User struct {
	Id              int    `json:"id" gorm:"primaryKey"`
	UserName        string `json:"userName"`
	Password        string `json:"password"`
	EthereumAddress string `json:"ethereumAddress"`
	StellarAddress  string `json:"stellarAddress"`
}
