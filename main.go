package main

import (
	"github.com/Pham-Xuan-Thanh/utils"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	//Connect Mysql Database
	db, err := utils.ConnectToDatabaseMysql()
}
