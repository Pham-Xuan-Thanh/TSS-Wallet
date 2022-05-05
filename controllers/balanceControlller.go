package controllers

import (
	"net/http"

	"github.com/Pham-Xuan-Thanh/TSS-Wallet/dto"
	"github.com/Pham-Xuan-Thanh/TSS-Wallet/helpers"

	"github.com/Pham-Xuan-Thanh/TSS-Wallet/services"
	"github.com/gin-gonic/gin"
)

type balancecontroller struct {
	BalanceService services.BalanceService
}

type BalanceController interface {
	GetBalance(ctx *gin.Context)
}

func (b *balancecontroller) GetBalance(ctx *gin.Context) {
	var addr dto.AddressDTO
	err := ctx.ShouldBind(&addr)
	if err != nil {
		res := helpers.BuildErrorResponse("Invalid Parameter", err.Error(), helpers.EmptyObject{})
		ctx.JSON(http.StatusBadGateway, res)
	} // Bind(&addr)
	result, err := b.BalanceService.GetBalance(addr)
	if err != nil {
		res := helpers.BuildErrorResponse("Invalid Address", err.Error(), helpers.EmptyObject{})
		ctx.JSON(http.StatusBadGateway, res)
	} else {
		response := helpers.BuildResponse(true, "Successfully", result)
		ctx.JSON(http.StatusOK, response)
	}
}

func NewBalanceController(b services.BalanceService) BalanceController {
	return &balancecontroller{b}
}
