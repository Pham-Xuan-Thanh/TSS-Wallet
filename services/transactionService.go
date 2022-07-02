package services

import (
	"encoding/hex"

	"github.com/Pham-Xuan-Thanh/TSS-Wallet/dto"
	"github.com/Pham-Xuan-Thanh/TSS-Wallet/repositories"
	"github.com/thanhxeon2470/TSS_chain/blockchain"
	"github.com/thanhxeon2470/TSS_chain/cli"
	"github.com/thanhxeon2470/TSS_chain/utils"
	"github.com/thanhxeon2470/TSS_chain/wallet"
)

type txservice struct {
	repositories.TxRepositories
}

type TxService interface {
	CreateTX(*dto.TransactionDTO) (bool, error)
	CreateTXipfs(*dto.ProposalDTO) (bool, error)
	// CreateSendTX(dto.TransactionSendDTO) (string, error)
	// CreateShareTX(dto.TransactionShareDTO) (string, error)
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
		pubKey := utils.Base58Decode([]byte(txin.PubKey))

		tmp := blockchain.TXInput{txID, txin.Vout, signature, pubKey}
		inputs = append(inputs, tmp)
	}

	// Import Tx inputs
	dtoTxOuts := txDto.TxOuts
	for _, txout := range dtoTxOuts {

		tmp := blockchain.NewTXOutput(txout.Value, txout.Address)
		outputs = append(outputs, *tmp)
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
		if err != nil {
			return false, err
		}
		signature, err := hex.DecodeString(txin.Signature)
		if err != nil {
			return false, err
		}
		pubKey := utils.Base58Decode([]byte(txin.PubKey))

		tmp := blockchain.TXInput{txID, txin.Vout, signature, pubKey}
		inputs = append(inputs, tmp)
	}

	// Import Tx inputs
	dtoTxOuts := propoDto.Tx.TxOuts
	for _, txout := range dtoTxOuts {

		tmp := blockchain.NewTXOutput(txout.Value, txout.Address)
		outputs = append(outputs, *tmp)
	}

	// Import Tx ipfs
	dtoTxIPFS := propoDto.Tx.TxIPFS
	for _, txipfs := range dtoTxIPFS {
		pubKeyOwner := utils.Base58Decode([]byte(txipfs.PubKeyOwner))
		signatureFile, err := hex.DecodeString(txipfs.SignatureFile)
		if err != nil {
			return false, err
		}
		ipfsHashEnc, err := hex.DecodeString(txipfs.IpfsHashEnc)

		var allowUsers [][]byte
		allowUsers = append(allowUsers, wallet.HashPubKey(pubKeyOwner))

		if txipfs.PubKeyHash != "" {
			allowUser, err := hex.DecodeString(txipfs.PubKeyHash)
			if err != nil {
				return false, err
			}
			allowUsers = append(allowUsers, allowUser)
		}
		tmp := blockchain.NewTXIpfs(pubKeyOwner, signatureFile, ipfsHashEnc, allowUsers)
		ipfsList = append(ipfsList, *tmp)
	}

	txID, err := hex.DecodeString(propoDto.Tx.TxID)
	if err != nil {
		return false, err
	}
	tx := blockchain.Transaction{txID, ipfsList, inputs, outputs}

	address, err := hex.DecodeString(propoDto.StorageMiningAddress)
	if err != nil {
		return false, err
	}
	proposal := cli.Proposal{txID, address, []byte(propoDto.IpfsHash), propoDto.Fee}
	res, err := txser.TxRepositories.CreateProposal(&proposal)
	if err != nil {
		return false, err
	}
	if res != false {
		return false, nil
	}
	return txser.TxRepositories.CreateTX(&tx)
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
