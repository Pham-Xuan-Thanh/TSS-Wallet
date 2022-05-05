package dto

type TransactionDTO struct {
	PrivKey      string   `json:"priv_key" binding:"required"`
	Reciever     string   `json:"reciever" binding:"required"`
	Amount       int      `json:"amount" binding:"required"`
	AllowAddress []string `json:"allow_address"`
	FilePath     string   `json:"file_path" binding:"required"`
}
