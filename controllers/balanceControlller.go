package controllers

import (
	"fmt"
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
	FindIPFSHash(ctx *gin.Context)
}

func (b *balancecontroller) GetBalance(ctx *gin.Context) {
	var addr dto.AddressDTO
	err := ctx.ShouldBind(&addr)
	if err != nil {
		res := helpers.BuildErrorResponse("Invalid Parameter", err.Error(), helpers.EmptyObject{})
		ctx.JSON(http.StatusBadGateway, res)
		return
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

func (b *balancecontroller) FindIPFSHash(ctx *gin.Context) {
	var ipfsHash dto.IPFSHASH
	err := ctx.ShouldBind(&ipfsHash)
	fmt.Print("whattt... ", ipfsHash.IPFSHash)
	if err != nil {
		fmt.Print("whattt... ", ipfsHash.IPFSHash)
		res := helpers.BuildErrorResponse("Invalid Parameter", err.Error(), helpers.EmptyObject{})
		ctx.JSON(http.StatusBadGateway, res)
		return
	}
	fmt.Print("whattt... ", ipfsHash.IPFSHash)
	result, err := b.BalanceService.FindIPFSHash(ipfsHash.IPFSHash)
	if err != nil {
		res := helpers.BuildErrorResponse("Invalid IPFS hash", err.Error(), helpers.EmptyObject{})
		ctx.JSON(http.StatusBadGateway, res)
	} else {
		response := helpers.BuildResponse(true, "Successfully", result)
		ctx.JSON(http.StatusOK, response)
	}
}

func NewBalanceController(b services.BalanceService) BalanceController {
	return &balancecontroller{b}
}
