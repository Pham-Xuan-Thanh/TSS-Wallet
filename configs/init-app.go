package configs

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/thanhxeon2470/TSS_chain/cli"

	"github.com/Pham-Xuan-Thanh/TSS-Wallet/controllers"
	"github.com/Pham-Xuan-Thanh/TSS-Wallet/docs"
	"github.com/Pham-Xuan-Thanh/TSS-Wallet/repositories"
	"github.com/Pham-Xuan-Thanh/TSS-Wallet/services"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	// "gorm.io/gorm"
)

type Server struct {
	blkchain   cli.CLI
	router     *gin.Engine
	ipfsDaemon exec.Cmd
}

func InitServer(chaincore cli.CLI) *Server {
	url_docs := fmt.Sprintf("%s:%s", os.Getenv("SERVER_HOST"), os.Getenv("SERVER_PORT"))
	// programmatically set swagger info
	docs.SwaggerInfo.Title = "TSS Wallet App"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = url_docs
	docs.SwaggerInfo.Schemes = []string{"http"}

	server := &Server{blkchain: chaincore}
	router := gin.New()

	// init module

	// init logger
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "GET", "POST", "PATCH"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		// AllowOriginFunc: func(origin string) bool {
		// 	return origin == "https://github.com"
		// },
		// MaxAge: 12 * time.Hour,
	}))
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {

		// your custom format
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))

	router.Use(gin.Recovery())

	walletReposi := repositories.NewWalletRepository(chaincore)
	walletService := services.NewWalletService(walletReposi)
	walletController := controllers.NewWalletController(walletService)

	walletRouter := router.Group("/api/user/wallet")
	{
		walletRouter.POST("/", walletController.GetAddress)
		walletRouter.GET("/", walletController.CreateWallet)

	}

	balanceReposi := repositories.NewBalanceRepository(chaincore)
	balanceService := services.NewBalanceService(balanceReposi)
	balanceController := controllers.NewBalanceController(balanceService)

	balanceRouter := router.Group("/api/user/balance")
	{
		balanceRouter.POST("/", balanceController.GetBalance)
		balanceRouter.POST("/filehash", balanceController.FindIPFSHash)
	}

	txReposi := repositories.NewTxRepositories(chaincore)
	txService := services.NewTxService(txReposi)
	txController := controllers.NewTxController(txService)

	txRouter := router.Group("/api/user/transaction/create")
	{
		txRouter.POST("/", txController.CreateTX)
	}

	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	server.router = router

	ipfs := exec.Command("ipfs", "daemon")
	server.ipfsDaemon = *ipfs
	return server
}

func (server *Server) Start() error {
	address := fmt.Sprintf(":%s", os.Getenv("SERVER_PORT"))
	if _, err := exec.LookPath("ipfs"); err != nil {
		return err
	}
	// Join IPFS network
	err := server.ipfsDaemon.Start()
	if err != nil {
		return err
	}
	isquit := make(chan bool)
	// server.blkchain.ReindexUTXO()
	go func(core cli.CLI) {
		core.StartNode("")
		if <-isquit {
			return
		}
	}(server.blkchain)

	fmt.Println("IPFS daemon is running.... ")
	defer func() {
		isquit <- true
		fmt.Println("IPFS downed T.T")
		server.ipfsDaemon.Process.Kill()
	}()

	return server.router.Run(address)
}
