package dto

type AddressDTO struct {
	Address string `json:"address" binding:"required"`
}

type IPFSHASH struct {
	IPFSHash string `json:"ipfshash" binding:"required"`
}
