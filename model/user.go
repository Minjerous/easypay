package model

type User struct {
	Id       int    `json:"userId"`
	Username string `json:"nikename"`
	Password string
	Money    int `json:"money"`
	Avatar   string
}
