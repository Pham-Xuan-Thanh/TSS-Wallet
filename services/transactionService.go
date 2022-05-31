package services

import (
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
}

func (txser *txservice) CreateTX(txDto dto.TransactionDTO) (bool, error) {
	var tx entities.Transaction
	if err := smapping.FillStruct(&tx, smapping.MapFields(&txDto)); err != nil {
		return false, err
	}
	tx.FilePath = filepath.FromSlash(tx.FilePath)

	return txser.TxRepositories.CreateTX(tx)
}
func NewTxService(tx repositories.TxRepositories) TxService {
	return &txservice{tx}
}
