package dto

type TransactionDTO struct {
	PrivKey  string `json:"priv_key" binding:"required"`
	Reciever string `json:"reciever" binding:"required"`
	Amount   int    `json:"amount" binding:"required"`
}

type TransactionSendDTO struct {
	PrivKey  string `json:"priv_key" binding:"required"`
	Reciever string `json:"reciever" binding:"required"`
	Amount   int    `json:"amount" binding:"required"`
	FilePath string `json:"file_path" binding:"required"`
}

type TransactionShareDTO struct {
	PrivKey      string `json:"priv_key" binding:"required"`
	Reciever     string `json:"reciever" binding:"required"`
	Amount       int    `json:"amount" binding:"required"`
	PubKey2Share string `json:"allow_pubkey"`
	IpfsHashEnc  string `json:"ipfshash_enc" binding:"required"`
}
