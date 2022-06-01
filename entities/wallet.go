package entities

type Wallet struct {
	Address string `json:"address" binding:"required"`
	PrivKey string `json:"priv_key" binding:"required"`
	PubKey  string `json:"pub_key": binding:"required"`
}

type Address struct {
	Address string `json:"address" binding:"required"`
}
