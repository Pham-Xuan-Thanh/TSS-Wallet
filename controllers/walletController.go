package controllers

import (
	// "github.com/Pham-Xuan-Thanh/TSS-Wallet/dto"
	"net/http"

	"github.com/Pham-Xuan-Thanh/TSS-Wallet/dto"
	"github.com/Pham-Xuan-Thanh/TSS-Wallet/helpers"
	"github.com/Pham-Xuan-Thanh/TSS-Wallet/services"
	"github.com/gin-gonic/gin"
)

type walletcontroller struct {
	WalletService services.WalletService
}
type WalletController interface {
	CreateWallet(ctx *gin.Context)
	GetAddress(ctx *gin.Context)
}

// @BasePath /api/user/wallet

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /example/helloworld [get]
func (w *walletcontroller) CreateWallet(ctx *gin.Context) {
	result, err := w.WalletService.CreateWallet()
	if err != nil {
		res := helpers.BuildErrorResponse("Failed to create wallet", err.Error(), helpers.EmptyObject{})
		ctx.JSON(http.StatusBadGateway, res)
	} else {
		response := helpers.BuildResponse(true, "Successfully", result)
		ctx.JSON(http.StatusOK, response)
	}
}

//
// @Summary Add a new pet to the store
// @Description get string by ID
// @Accept  json
// @Produce  json
// @Param   some_id     path    int     true        "Some ID"
// @Success 200 {string} string	"ok"
// @Failure 400 {object} web.APIError "We need ID!!"
// @Failure 404 {object} web.APIError "Can not find ID"
// @Router /testapi/get-string-by-int/{some_id} [get]
func (w *walletcontroller) GetAddress(ctx *gin.Context) {
	var privKeyDto dto.PrivateKeyDto
	ctx.ShouldBind(&privKeyDto)
	result, err := w.WalletService.GetAddress(privKeyDto)
	if err != nil {
		res := helpers.BuildErrorResponse("Invalid Key", err.Error(), helpers.EmptyObject{})
		ctx.JSON(http.StatusBadGateway, res)
	} else {
		response := helpers.BuildResponse(true, "Successfully", result)
		ctx.JSON(http.StatusOK, response)
	}
}

func NewWalletController(w services.WalletService) WalletController {
	return &walletcontroller{w}
}
