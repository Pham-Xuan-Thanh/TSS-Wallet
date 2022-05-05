package services

import (
	"github.com/Pham-Xuan-Thanh/TSS-Wallet/dto"
	"github.com/Pham-Xuan-Thanh/TSS-Wallet/entities"
	"github.com/Pham-Xuan-Thanh/TSS-Wallet/repositories"
)

type walletservice struct {
	WalletRepo repositories.WalletRepositories
}

type WalletService interface {
	CreateWallet() (*entities.Wallet, error)
	GetAddress(privKeyDto dto.PrivateKeyDto) (*entities.Address, error)
}

func (w *walletservice) CreateWallet() (*entities.Wallet, error) {
	return w.WalletRepo.CreateWallet()
}

func (w *walletservice) GetAddress(privKeyDto dto.PrivateKeyDto) (*entities.Address, error) {
	wallet := entities.Wallet{
		Address: "",
		PrivKey: privKeyDto.PrivKey,
	}
	return w.WalletRepo.GetAddress(wallet)
}

func NewWalletService(w repositories.WalletRepositories) WalletService {
	return &walletservice{w}
}
