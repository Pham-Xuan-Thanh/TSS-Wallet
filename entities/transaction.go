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
type Txinput struct {
	TxID  string `json:"tx_id"`
	Vout  int    `json:"vout"`
	Value int    `json:"value"`
}

type TransactionInputs struct {
	TXins []Txinput `json:"tx_ins"`
}
