package repositories

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/rpc"
	"os"
	"os/exec"
	"strings"

	"github.com/Pham-Xuan-Thanh/TSS-Wallet/entities"
	"github.com/thanhxeon2470/TSS_chain/blockchain"
	r "github.com/thanhxeon2470/TSS_chain/rpc"
)

type txrepositories struct {
}

type TxRepositories interface {
	CreateTX(*blockchain.Transaction) (bool, error)
	CreateProposal(*r.Proposal) (bool, error)
	// CreateSendTX(entities.Transaction) (string, error)
	// CreateShareTX(entities.Transaction) (string, error)
	GetTXins(addr string) (*entities.TransactionInputs, error)
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
	fhphase := strings.Split(stoutstr, "\n")[0]
	fh := strings.Split(fhphase, " ")[1]
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
	req := tx.Serialize()
	args := &r.Args{req}
	res := &r.Result{}

	serverAddress := os.Getenv(("SERVER_RPC"))
	client, err := rpc.DialHTTP("tcp", serverAddress)
	if err != nil {
		return false, err
	}

	err = client.Call("RPC.SendTx", args, res)
	if err != nil {
		return false, err
	}

	// Create Transaction to Propogate on network
	// fmt.Print("What ups")
	// tx.Sign(w.PrivateKey)

	return true, nil
}
func (txrepo *txrepositories) CreateProposal(proposal *r.Proposal) (bool, error) {

	req, err := r.GobEncode(proposal)
	if err != nil {
		return false, err
	}
	args := &r.Args{req}
	res := &r.Result{}

	serverAddress := os.Getenv(("SERVER_RPC"))
	client, err := rpc.DialHTTP("tcp", serverAddress)
	if err != nil {
		return false, err
	}

	err = client.Call("RPC.SendProposal", args, res)
	if err != nil {
		return false, err
	}
	return true, nil
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

func (txrepo *txrepositories) GetTXins(addr string) (*entities.TransactionInputs, error) {
	result := new(entities.TransactionInputs)

	req, err := r.GobEncode(r.Gettxins{addr})
	if err != nil {
		return nil, err
	}
	args := &r.Args{req}
	res := &r.Result{}
	serverAddress := os.Getenv(("SERVER_RPC"))
	client, err := rpc.DialHTTP("tcp", serverAddress)
	if err != nil {
		log.Fatal("dialing:", err)
	}

	err = client.Call("RPC.GetTxIns", args, res)
	if err != nil {
		log.Fatal("Call RPC:", err)
	}
	var payload r.Txins
	err = r.GobDecode(res.Res, &payload)
	if err != nil {
		return nil, err
	}
	input := new(entities.Txinput)
	for txid, infos := range payload.ValidOutputs {
		input.TxID = txid
		for _, info := range infos {
			input.Vout = info[0]
			input.Value = info[1]
			result.TXins = append(result.TXins, *input)
		}
	}

	return result, nil
}

func NewTxRepositories() TxRepositories {
	return &txrepositories{}
}
