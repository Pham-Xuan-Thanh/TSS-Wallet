package main

import (
	"log"

	"github.com/Pham-Xuan-Thanh/TSS-Wallet/configs"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/thanhxeon2470/TSS_chain/cli"
)

const (
	privPayed = "7jczDSiSxs6ZNzQdTT6TydNHgNLKBNLjPWzJ2JHkiX"
	pubkPayed = "1Jg2aTGmJQpPMTNyKr7CMRspxV1pUAe5x4"
	privNone  = "2brnL7iBJjQ8S1rZiNPrgz18J3dR9aZDssi7NkDGQfXg"
	pubkNone  = "16Jcx92y5gv1vfVfzqFkghqXYSCdTkmgV7"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	//Connect Mysql Database
	// db, err := utils.ConnectToDatabaseMysql()

	// if err != nil {
	// 	defer utils.CloseConnectToDatabaseMysql(db)
	// }

	gin.ForceConsoleColor()
	blkchaincore := cli.CLI{}
	server := configs.InitServer(blkchaincore)
	// Miragte db
	// db.AutoMigrate(entities.Device{})
	err = server.Start()
	if err != nil {
		log.Fatal("Cannot start server:", err)
	}
}
