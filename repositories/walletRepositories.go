package repositories

import (
	"github.com/Pham-Xuan-Thanh/TSS-Wallet/entities"
	"github.com/thanhxeon2470/TSS_chain/cli"
)

type WalletRepositories interface {
	CreateWallet() (*entities.Wallet, error)
	GetAddress(entities.Wallet) (*entities.Address, error)
}

type walletrepositories struct {
	blkchain cli.CLI
}

func (w *walletrepositories) CreateWallet() (*entities.Wallet, error) {
	var wallet entities.Wallet
	var addr, pub, priv = w.blkchain.CreateWallet()
	wallet.PrivKey = string(priv)
	wallet.PubKey = string(pub)
	wallet.Address = string(addr)
	return &wallet, nil
}

func (w *walletrepositories) GetAddress(getAddrwallet entities.Wallet) (*entities.Address, error) {
	res := entities.Address{}
	res.Address = string(w.blkchain.AddWallet([]byte(getAddrwallet.PrivKey)))
	return &res, nil
}

func NewWalletRepository(blkchain cli.CLI) WalletRepositories {
	return &walletrepositories{blkchain}
}
