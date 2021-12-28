package model

type User struct {
	Id       int    `json:"userId"`
	Username string `json:"nikename"`
	Password string
	Money    float64 `json:"money"`
	Avatar   string
}
