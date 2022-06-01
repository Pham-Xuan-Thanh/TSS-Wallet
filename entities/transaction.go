package entities

type Transaction struct {
	PrivKey      string `json:"priv_key"`
	Reciever     string `json:"reciever"`
	Amount       int    `json:"amount"`
	PubKey2Share string `json:"allow_pubkey"`
	FilePath     string `json:"file_path"`
	IpfsHashEnc  string
	FileHash     string
}
