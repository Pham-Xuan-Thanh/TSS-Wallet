package services

import (
	"encoding/hex"

	"fmt"

	"github.com/Pham-Xuan-Thanh/TSS-Wallet/dto"
	"github.com/Pham-Xuan-Thanh/TSS-Wallet/entities"
	"github.com/Pham-Xuan-Thanh/TSS-Wallet/repositories"
	"github.com/thanhxeon2470/TSS_chain/blockchain"
	"github.com/thanhxeon2470/TSS_chain/rpc"
)

type txservice struct {
	repositories.TxRepositories
}

type TxService interface {
	CreateTX(*dto.TransactionDTO) (bool, error)
	CreateTXipfs(*dto.ProposalDTO) (bool, error)
	// CreateSendTX(dto.TransactionSendDTO) (string, error)
	// CreateShareTX(dto.TransactionShareDTO) (string, error)
	GetTXins(*dto.GetInsDTO) (*entities.TransactionInputs, error)
}

func (txser *txservice) CreateTX(txDto *dto.TransactionDTO) (bool, error) {
	var inputs []blockchain.TXInput
	var outputs []blockchain.TXOutput

	// Import Tx inputs
	dtoTxIns := txDto.TxIns
	for _, txin := range dtoTxIns {
		txID, err := hex.DecodeString(txin.TxID)
		if err != nil {
			return false, err
		}
		signature, err := hex.DecodeString(txin.Signature)
		if err != nil {
			return false, err
		}
		pubKey, err := hex.DecodeString(txin.PubKey)
		if err != nil {
			return false, err
		}
		tmp := blockchain.TXInput{txID, txin.Vout, signature, pubKey}
		inputs = append(inputs, tmp)
	}

	// Import Tx outputs
	dtoTxOuts := txDto.TxOuts
	for _, txout := range dtoTxOuts {
		pubKey, err := hex.DecodeString(txout.PubKeyHash)
		if err != nil {
			return false, err
		}
		tmp := blockchain.TXOutput{txout.Value, pubKey}
		outputs = append(outputs, tmp)
	}
	txID, err := hex.DecodeString(txDto.TxID)
	if err != nil {
		return false, err
	}
	tx := blockchain.Transaction{txID, nil, inputs, outputs}

	return txser.TxRepositories.CreateTX(&tx)
}

func (txser *txservice) CreateTXipfs(propoDto *dto.ProposalDTO) (bool, error) {
	var inputs []blockchain.TXInput
	var outputs []blockchain.TXOutput
	var ipfsList []blockchain.TXIpfs

	// Import Tx inputs
	dtoTxIns := propoDto.Tx.TxIns
	for _, txin := range dtoTxIns {
		txID, err := hex.DecodeString(txin.TxID)
		fmt.Println("cc 1")
		if err != nil {
			return false, err
		}
		signature, err := hex.DecodeString(txin.Signature)
		fmt.Println("cc 2")

		if err != nil {
			return false, err
		}
		pubKey, err := hex.DecodeString(txin.PubKey)
		if err != nil {
			return false, err
		}
		tmp := blockchain.TXInput{txID, txin.Vout, signature, pubKey}
		inputs = append(inputs, tmp)
	}

	// Import Tx outputs
	dtoTxOuts := propoDto.Tx.TxOuts
	for _, txout := range dtoTxOuts {
		pubKey, err := hex.DecodeString(txout.PubKeyHash)
		fmt.Println("cc 3")

		if err != nil {
			return false, err
		}
		tmp := blockchain.TXOutput{Value: txout.Value, PubKeyHash: pubKey}
		outputs = append(outputs, tmp)
	}

	// Import Tx ipfs
	dtoTxIPFS := propoDto.Tx.TxIPFS
	for _, txipfs := range dtoTxIPFS {
		pubKeyOwner, err := hex.DecodeString(txipfs.PubKeyOwner)
		if err != nil {
			return false, err
		}
		signatureFile, err := hex.DecodeString(txipfs.SignatureFile)
		fmt.Println("cc 4")

		if err != nil {
			return false, err
		}
		ipfsHashEnc, err := hex.DecodeString(txipfs.IpfsHashEnc)
		if err != nil {
			return false, err
		}
		// var allowUsers []byte
		fmt.Println("")
		// allowUsers = append(allowUsers, wallet.HashPubKey(pubKeyOwner))
		allowUser, err := hex.DecodeString(txipfs.PubKeyHash)
		if err != nil {
			return false, err
		}
		// if txipfs.PubKeyHash != "" {
		// 	if err != nil {
		// 		return false, err
		// 	}
		// 	allowUsers = append(allowUsers, allowUser...)
		// }
		tmp := blockchain.NewTXIpfs(pubKeyOwner, signatureFile, ipfsHashEnc, allowUser)
		ipfsList = append(ipfsList, *tmp)
		fmt.Println("xu vcl ady ne ", ipfsHashEnc, allowUser)
	}

	txID, err := hex.DecodeString(propoDto.Tx.TxID)
	fmt.Println("cc 5")

	if err != nil {
		return false, err
	}
	tx := blockchain.Transaction{ID: txID, Ipfs: ipfsList, Vin: inputs, Vout: outputs}

	address := []byte(propoDto.StorageMiningAddress)
	// if err != nil {
	// 	return false, err
	// }
	proposal := rpc.Proposal{TxHash: txID, StorageMiningAddress: address, FileHash: []byte(propoDto.IpfsHash), Amount: propoDto.Fee}
	fmt.Println("Cli proposa thu 2 ", proposal)
	res, err := txser.TxRepositories.CreateProposal(&proposal)
	// if err != nil {
	// 	return false, err
	// }
	if res != true {
		return false, err
	}
	fmt.Println(("toi dc return "))
	return txser.TxRepositories.CreateTX(&tx)
}

func (txser *txservice) GetTXins(getTXins *dto.GetInsDTO) (*entities.TransactionInputs, error) {
	address := getTXins.Address
	return txser.TxRepositories.GetTXins(address)
}

// func (txser *txservice) CreateSendTX(txDto dto.TransactionSendDTO) (string, error) {
// 	var tx entities.Transaction
// 	if err := smapping.FillStruct(&tx, smapping.MapFields(&txDto)); err != nil {
// 		return "", err
// 	}
// 	tx.FilePath = filepath.FromSlash(tx.FilePath)

// 	return txser.TxRepositories.CreateSendTX(tx)
// }

// func (txser *txservice) CreateShareTX(txDto dto.TransactionShareDTO) (string, error) {
// 	var tx entities.Transaction
// 	if err := smapping.FillStruct(&tx, smapping.MapFields(&txDto)); err != nil {
// 		fmt.Println("Loi o so -3")
// 		return "", err
// 	}

// 	return txser.TxRepositories.CreateShareTX(tx)
// }
func NewTxService(tx repositories.TxRepositories) TxService {
	return &txservice{tx}
}
