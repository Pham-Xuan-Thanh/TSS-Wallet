package repositories

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"os"
	"os/exec"
	"strings"
	s "strings"

	"github.com/Pham-Xuan-Thanh/TSS-Wallet/entities"
	"github.com/thanhxeon2470/TSS_chain/blockchain"
	"github.com/thanhxeon2470/TSS_chain/cli"
	"github.com/thanhxeon2470/TSS_chain/rpc"
)

type txrepositories struct {
}

type TxRepositories interface {
	CreateTX(*blockchain.Transaction) (bool, error)
	CreateProposal(*cli.Proposal) (bool, error)
	// CreateSendTX(entities.Transaction) (string, error)
	// CreateShareTX(entities.Transaction) (string, error)
	GetTXins(entities.Address) (*entities.TransactionInputs, error)
}

type ipfsID struct {
	ID              string   `json:"ID"`
	PublicKey       string   `json:"PublicKey"`
	Addresses       []string `json:"Addresses"`
	AgentVersion    string   `json:"AgentVersion"`
	ProtocolVersion string   `json:"ProtocolVersion"`
	Protocols       []string `json:"Protocols"`
}

func fileExists(name string) (bool, error) {
	_, err := os.Stat(name)
	if err == nil {
		return true, nil
	}
	if errors.Is(err, os.ErrNotExist) {
		return false, nil
	}
	return false, err
}
func ipfsIsRunning() bool {
	idCmd := exec.Command("ipfs", "id")
	stdout, err := idCmd.Output()
	if err != nil {
		return false
	}
	idIn4 := ipfsID{}
	err = json.Unmarshal(stdout, &idIn4)
	if err != nil {
		fmt.Println("unmarshle k dc ")
		return false
	}

	if idIn4.Addresses == nil {
		return false
	}
	return true

}

func getFileHash(stout []byte) string {
	stoutstr := string(stout)
	fhphase := s.Split(stoutstr, "\n")[0]
	fh := s.Split(fhphase, " ")[1]
	return fh
}
func ipfsAdd(filepath string) (string, error) {
	if isRunning := ipfsIsRunning(); !isRunning {
		return "", fmt.Errorf("IPFS is not running @.@")
	}
	addCmd := exec.Command("ipfs", "add", filepath)

	stdout, err := addCmd.Output()
	if err != nil {
		return "", err
	}
	return getFileHash(stdout), nil
}
func (txrepo *txrepositories) CreateTX(tx *blockchain.Transaction) (bool, error) {

	// Create Transaction to Propogate on network
	// fmt.Print("What ups")
	cli.SendTx(strings.Split(os.Getenv("KNOWNNODE"), "_")[0], tx)
	return true, nil
}
func (txrepo *txrepositories) CreateProposal(proposal *cli.Proposal) (bool, error) {
	cli.SendProposal(strings.Split(os.Getenv("KNOWNNODE"), "_")[0], *proposal)
	// return true, nil
}

// func (txrepo *txrepositories) CreateSendTX(tx entities.Transaction) (string, error) {
// 	//Check file exist??
// 	if isExist, err := fileExists(tx.FilePath); err != nil || !isExist {
// 		return "", err
// 	}

// 	// Add file to ipfs
// 	fh, err := ipfsAdd(tx.FilePath)
// 	if err != nil {
// 		return "", err
// 	}
// 	tx.FileHash = fh
// 	// Create Transaction to Propogate on network
// 	// return txrepo.blkchain.SendProposal(tx.PrivKey, tx.Reciever, tx.Amount, tx.FileHash), nil
// }
// func (txrepo *txrepositories) CreateShareTX(tx entities.Transaction) (string, error) {

// 	// Create Transaction to Propogate on network
// 	// return txrepo.blkchain.Share(tx.PrivKey, tx.Reciever, tx.Amount, tx.PubKey2Share, tx.IpfsHashEnc), nil
// }

func (txrepo *txrepositories) GetTXins(addr entities.Address) (*entities.TransactionInputs, error) {
	result := new(entities.TransactionInputs)

	rpc.SendGetTxIns(os.Getenv("SERVER_RPC"), addr.Address)
	port := os.Getenv("PORT_LSRPC")
	port = fmt.Sprintf(":%s", port)
	ln, err := net.Listen("tcp", port)
	if err != nil {
		return nil, err
	}
	defer ln.Close()

	conn, err := ln.Accept()
	if err != nil {
		return nil, err
	}
	buff_t := rpc.HandleRPCReceive(conn)
	var payload rpc.Txins
	buff := bytes.NewBuffer(buff_t)
	dec := gob.NewDecoder(buff)
	err = dec.Decode(&payload)
	if err != nil {
		return nil, err
	}

	result.TXins = payload.ValidOutputs

	return result, nil
}

func NewTxRepositories() TxRepositories {
	return &txrepositories{}
}
