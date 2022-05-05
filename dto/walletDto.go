package dto

type PrivateKeyDto struct {
	PrivKey string `json:"priv_key" binding:"required"`
}
