package repositories

import (
	"fmt"
	"time"

	"github.com/Pham-Xuan-Thanh/TSS-Wallet/entities"
	"github.com/thanhxeon2470/TSS_chain/cli"
)

type balancerepositories struct {
	blkchain cli.CLI
}

type BalanceRepositories interface {
	GetBalance(entities.Address) (*entities.Balance, error)
	FindIPFSHash(string) (*entities.AllowUsers, error)
}

func (b *balancerepositories) GetBalance(balanceAddr entities.Address) (*entities.Balance, error) {
	result := entities.NewBalance()
	var balance, fileInfo = b.blkchain.GetBalance(string(balanceAddr.Address))
	result.Balanced = balance
	for iHash, in4 := range fileInfo {
		fmt.Printf("=================== %s %t %s", iHash, in4.Author, time.Unix(in4.Exp, 0))
		result.FileOwned[iHash] = entities.IPFSInfo{Author: in4.Author, Exp: in4.Exp}
	}
	return result, nil
}

func (b *balancerepositories) FindIPFSHash(ipfsHash string) (*entities.AllowUsers, error) {
	result := entities.NewAllowUsers()

	resp := b.blkchain.FindIPFS(ipfsHash)
	fmt.Println("aaaaaaaa", ipfsHash)
	for address, isAuthor := range resp {
		fmt.Println("CL", address, isAuthor)
		result.Users[address] = isAuthor
	}
	return result, nil
}

func NewBalanceRepository(blkchain cli.CLI) BalanceRepositories {
	return &balancerepositories{blkchain}
}
