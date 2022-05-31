package repositories

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/exec"
	s "strings"

	"github.com/Pham-Xuan-Thanh/TSS-Wallet/entities"
	"github.com/thanhxeon2470/TSS_chain/cli"
)

type txrepositories struct {
	blkchain cli.CLI
}

type TxRepositories interface {
	CreateTX(entities.Transaction) (bool, error)
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
func (txrepo *txrepositories) CreateTX(tx entities.Transaction) (bool, error) {
	//Check file exist??
	if isExist, err := fileExists(tx.FilePath); err != nil || !isExist {
		return isExist, err
	}

	// Add file to ipfs
	fh, err := ipfsAdd(tx.FilePath)
	if err != nil {
		return false, err
	}
	tx.FileHash = fh
	// Create Transaction to Propogate on network

	return txrepo.blkchain.SendProposal(tx.PrivKey, tx.Reciever, tx.Amount, tx.AllowAddress, tx.FileHash), nil

}
func NewTxRepositories(blkchain cli.CLI) TxRepositories {
	return &txrepositories{blkchain: blkchain}
}
