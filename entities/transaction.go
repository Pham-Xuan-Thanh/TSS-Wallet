package entities

type Transaction struct {
	PrivKey      string   `json:"priv_key"`
	Reciever     string   `json:"reciever"`
	Amount       int      `json:"amount"`
	AllowAddress []string `json:"allow_address"`
	FilePath     string   `json:"file_path"`
	FileHash     string
}
