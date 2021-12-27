package main

import (
	"easypay/api"
	"easypay/dao"
)

func main() {
	dao.InitDB()
	api.InitRouter()
}
