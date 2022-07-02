package controllers

import (
	"fmt"
	"net/http"

	"github.com/Pham-Xuan-Thanh/TSS-Wallet/dto"
	"github.com/Pham-Xuan-Thanh/TSS-Wallet/helpers"
	"github.com/Pham-Xuan-Thanh/TSS-Wallet/services"
	"github.com/gin-gonic/gin"
)

type txcontroller struct {
	services.TxService
}

type TxController interface {
	CreateTX(ctx *gin.Context)
	CreateTXipfs(ctx *gin.Context) // Create tx send file
	// CreateShareTX(ctx *gin.Context) // Tx to share file
}

func (tx *txcontroller) CreateTX(ctx *gin.Context) {
	var txDto *dto.TransactionDTO
	if err := ctx.ShouldBind(txDto); err != nil {
		res := helpers.BuildErrorResponse("Invalid Parameter", err.Error(), helpers.EmptyObject{})
		ctx.JSON(http.StatusBadGateway, res)
		return
	}
	fmt.Print("What ups")
	result, err := tx.TxService.CreateTX(txDto)
	if err != nil {
		res := helpers.BuildErrorResponse("Invalid paramerter", err.Error(), helpers.EmptyObject{})
		ctx.JSON(http.StatusBadGateway, res)
	} else {
		var response helpers.Response
		if result {
			response = helpers.BuildResponse(true, "Successfully", result)
		} else {
			response = helpers.BuildResponse(false, "Unsuccessfully", result)
		}
		ctx.JSON(http.StatusOK, response)
	}
}

func (tx *txcontroller) CreateTXipfs(ctx *gin.Context) {
	var propoDto *dto.ProposalDTO

	if err := ctx.ShouldBind(propoDto); err != nil {
		res := helpers.BuildErrorResponse("Invalid Parameter", err.Error(), helpers.EmptyObject{})
		ctx.JSON(http.StatusBadGateway, res)
		return
	}
	fmt.Print("What ups")
	result, err := tx.TxService.CreateTXipfs(propoDto)
	if err != nil {
		res := helpers.BuildErrorResponse("Invalid paramerter", err.Error(), helpers.EmptyObject{})
		ctx.JSON(http.StatusBadGateway, res)
	} else {
		var response helpers.Response
		if result {
			response = helpers.BuildResponse(true, "Successfully", result)
		} else {
			response = helpers.BuildResponse(false, "Unsuccessfully", result)
		}
		ctx.JSON(http.StatusOK, response)
	}
}

// func (tx *txcontroller) CreateSendTX(ctx *gin.Context) {
// 	var txDto dto.TransactionSendDTO
// 	if err := ctx.ShouldBind(&txDto); err != nil {
// 		res := helpers.BuildErrorResponse("Invalid Parameter", err.Error(), helpers.EmptyObject{})
// 		ctx.JSON(http.StatusBadGateway, res)
// 		return
// 	}
// 	result, err := tx.TxService.CreateSendTX(txDto)
// 	if err != nil {
// 		res := helpers.BuildErrorResponse("Invalid paramerter", err.Error(), helpers.EmptyObject{})
// 		ctx.JSON(http.StatusBadGateway, res)
// 	} else {
// 		var response helpers.Response
// 		if result != "" {
// 			response = helpers.BuildResponse(true, "Successfully", result)
// 		} else {
// 			response = helpers.BuildResponse(false, "Unsuccessfully", result)
// 		}
// 		ctx.JSON(http.StatusOK, response)
// 	}
// }

// func (tx *txcontroller) CreateShareTX(ctx *gin.Context) {
// 	var txDto dto.TransactionShareDTO
// 	if err := ctx.ShouldBind(&txDto); err != nil {
// 		res := helpers.BuildErrorResponse("Invalid Parameter", err.Error(), helpers.EmptyObject{})
// 		ctx.JSON(http.StatusBadGateway, res)
// 		return
// 	}
// 	result, err := tx.TxService.CreateShareTX(txDto)
// 	if err != nil {
// 		res := helpers.BuildErrorResponse("Invalid paramerter", err.Error(), helpers.EmptyObject{})
// 		ctx.JSON(http.StatusBadGateway, res)
// 	} else {
// 		var response helpers.Response
// 		if result != "" {
// 			response = helpers.BuildResponse(true, "Successfully", result)
// 		} else {
// 			response = helpers.BuildResponse(false, "Unsuccessfully", result)
// 		}
// 		ctx.JSON(http.StatusOK, response)
// 	}
// }

func NewTxController(tx services.TxService) TxController {
	return &txcontroller{tx}
}
