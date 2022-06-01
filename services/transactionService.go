package services

import (
	"fmt"
	"path/filepath"

	"github.com/Pham-Xuan-Thanh/TSS-Wallet/dto"
	"github.com/Pham-Xuan-Thanh/TSS-Wallet/entities"
	"github.com/Pham-Xuan-Thanh/TSS-Wallet/repositories"
	"github.com/mashingan/smapping"
)

type txservice struct {
	repositories.TxRepositories
}

type TxService interface {
	CreateTX(dto.TransactionDTO) (bool, error)
	CreateSendTX(dto.TransactionSendDTO) (string, error)
	CreateShareTX(dto.TransactionShareDTO) (string, error)
}

func (txser *txservice) CreateTX(txDto dto.TransactionDTO) (bool, error) {
	var tx entities.Transaction
	if err := smapping.FillStruct(&tx, smapping.MapFields(&txDto)); err != nil {
		return false, err
	}
	fmt.Print("What ups")
	return txser.TxRepositories.CreateTX(tx)
}

func (txser *txservice) CreateSendTX(txDto dto.TransactionSendDTO) (string, error) {
	var tx entities.Transaction
	if err := smapping.FillStruct(&tx, smapping.MapFields(&txDto)); err != nil {
		return "", err
	}
	tx.FilePath = filepath.FromSlash(tx.FilePath)

	return txser.TxRepositories.CreateSendTX(tx)
}

func (txser *txservice) CreateShareTX(txDto dto.TransactionShareDTO) (string, error) {
	var tx entities.Transaction
	if err := smapping.FillStruct(&tx, smapping.MapFields(&txDto)); err != nil {
		fmt.Println("Loi o so -3")
		return "", err
	}

	return txser.TxRepositories.CreateShareTX(tx)
}
func NewTxService(tx repositories.TxRepositories) TxService {
	return &txservice{tx}
}
