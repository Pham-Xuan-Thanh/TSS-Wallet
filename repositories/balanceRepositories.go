package repositories

import (
	"github.com/Pham-Xuan-Thanh/TSS-Wallet/entities"
	"github.com/thanhxeon2470/TSS_chain/cli"
)

type balancerepositories struct {
	blkchain cli.CLI
}

type BalanceRepositories interface {
	GetBalance(entities.Address) (*entities.Balance, error)
}

func (b *balancerepositories) GetBalance(balanceAddr entities.Address) (*entities.Balance, error) {
	var result entities.Balance
	result.Address, result.Balanced, result.FileOwned = b.blkchain.GetBalance(string(balanceAddr.Address))

	return &result, nil
}

func NewBalanceRepository(blkchain cli.CLI) BalanceRepositories {
	return &balancerepositories{blkchain}
}
