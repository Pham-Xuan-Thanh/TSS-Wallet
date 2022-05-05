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
}

func (b *balanceservice) GetBalance(balanceGet dto.AddressDTO) (*entities.Balance, error) {
	balanceAddr := entities.Address{balanceGet.Address}

	// err := smapping.FillStruct(&balanceAddr, smapping.MapFields(&balanceGet))
	// if err != nil {
	// 	return nil, err
	// }

	return b.BalanceRepo.GetBalance(balanceAddr)
}
func NewBalanceService(b repositories.BalanceRepositories) BalanceService {
	return &balanceservice{b}
}
