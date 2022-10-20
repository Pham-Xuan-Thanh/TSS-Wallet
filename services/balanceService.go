package services

import (
	"github.com/Pham-Xuan-Thanh/TSS-Wallet/dto"
	"github.com/Pham-Xuan-Thanh/TSS-Wallet/entities"
	"github.com/Pham-Xuan-Thanh/TSS-Wallet/repositories"
)

type balanceservice struct {
	BalanceRepo repositories.BalanceRepositories
}

type BalanceService interface {
	GetBalance(dto.AddressDTO) (*entities.Balance, error)
	FindIPFSHash(string) (*entities.AllowUsers, error)
}

func (b *balanceservice) GetBalance(balanceGet dto.AddressDTO) (*entities.Balance, error) {

	// err := smapping.FillStruct(&balanceAddr, smapping.MapFields(&balanceGet))
	// if err != nil {
	// 	return nil, err
	// }

	return b.BalanceRepo.GetBalance(balanceGet.Address)
}

func (b *balanceservice) FindIPFSHash(ipfsHash string) (*entities.AllowUsers, error) {
	return b.BalanceRepo.FindIPFSHash(ipfsHash)
}
func NewBalanceService(b repositories.BalanceRepositories) BalanceService {
	return &balanceservice{b}
}
