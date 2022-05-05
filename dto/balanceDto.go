package dto

type AddressDTO struct {
	Address string `json:"address" binding:"required"`
}
