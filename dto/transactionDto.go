package dto

type TransactionInputsDTO struct {
	TxID      string `json:"tx_id" binding:"required"`
	Vout      int    `json:"vout" binding:"required"`
	Signature string `json:"signature" binding:"required"`
	PubKey    string `json:"pub_key" binding:"required"`
}

type TransactionOutputsDTO struct {
	Value      int    `json:"value" binding:"required"`
	PubKeyHash string `json:"pub_key_hash" binding:"required"` // string hex
}

type TransactionIPFSDTO struct {
	PubKeyOwner   string `json:"pub_key_owner" binding:"required"` // Pubkey to Verify
	SignatureFile string `json:"signature" binding:"required"`     // Sign IFPS ENC

	IpfsHashEnc string `json:"ipfs_enc" binding:"required"`
	PubKeyHash  string `json:"pub_key_user" binding:"required"` // Allow user if TX share

	Exp int64 `json:"exp" binding:"required"`
}

type TransactionDTO struct {
	TxID   string                  `json:"tx_id" binding:"required"`
	TxIns  []TransactionInputsDTO  `json:"tx_ins" binding:"required"`
	TxOuts []TransactionOutputsDTO `json:"tx_outs" binding:"required"`
	TxIPFS []TransactionIPFSDTO    `json:"tx_ipfs"`
}

type ProposalDTO struct {
	Tx                   TransactionDTO `json:"tx" binding:"required"`
	StorageMiningAddress string         `json:"address" binding:"required"`
	IpfsHash             string         `json:"ipfs" binding:"required"`
	Fee                  int            `json:"fee" binding:"required"`
}

type GetInsDTO struct {
	Address string `json:"address" binding:"required"`
}

// type TransactionShareIPFSDTO struct {
// 	PubKeyOwner   string `json:"pub_key_owner" binding:"required"` // Pubkey to Verify
// 	SignatureFile string `json:"signature" binding:"required"`     // Sign IFPS ENC

// 	IpfsHashEnc string `json:"ifps_enc" binding:"required"`
// 	PubKeyHash  string `json:"pub_key_user" binding:"required"`
// }

// type TransactionDTO struct {
// 	PrivKey  string `json:"priv_key" binding:"required"`
// 	Reciever string `json:"reciever" binding:"required"`
// 	Amount   int    `json:"amount" binding:"required"`
// }

// type TransactionDTO struct {
// 	PrivKey  string `json:"priv_key" binding:"required"`
// 	Reciever string `json:"reciever" binding:"required"`
// 	Amount   int    `json:"amount" binding:"required"`
// }

// type TransactionSendDTO struct {
// 	PrivKey  string `json:"priv_key" binding:"required"`
// 	Reciever string `json:"reciever" binding:"required"`
// 	Amount   int    `json:"amount" binding:"required"`
// 	FilePath string `json:"file_path" binding:"required"`
// }

// type TransactionShareDTO struct {
// 	PrivKey      string `json:"priv_key" binding:"required"`
// 	Reciever     string `json:"reciever" binding:"required"`
// 	Amount       int    `json:"amount" binding:"required"`
// 	PubKey2Share string `json:"allow_pubkey"`
// 	IpfsHashEnc  string `json:"ipfshash_enc" binding:"required"`
// }
