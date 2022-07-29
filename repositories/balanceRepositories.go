package repositories

import (
	"encoding/hex"
	"fmt"
	"log"
	"net/rpc"
	"os"
	"time"

	"github.com/Pham-Xuan-Thanh/TSS-Wallet/entities"
	r "github.com/thanhxeon2470/TSS_chain/rpc"
)

type balancerepositories struct {
}

type BalanceRepositories interface {
	GetBalance(string) (*entities.Balance, error)
	FindIPFSHash(string) (*entities.AllowUsers, error)
}

func (b *balancerepositories) GetBalance(addrTSS string) (*entities.Balance, error) {
	result := new(entities.Balance)
	req, err := r.GobEncode(r.Getbalance{addrTSS})
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

	err = client.Call("RPC.GetBlance", args, res)
	if err != nil {
		log.Fatal("Call RPC:", err)
	}

	var payload r.Balance
	err = r.GobDecode(res.Res, &payload)
	if err != nil {
		return nil, err
	}
	var balance, fileInfo = payload.Value, payload.FTXs
	result.Balanced = balance
	in4ipfs := new(entities.IPFSInfo)
	for iHash, in4 := range fileInfo {
		fmt.Printf("=================== %s %t %s", iHash, in4.Author, time.Unix(in4.Exp, 0))
		in4ipfs.Author = in4.Author
		in4ipfs.IpfsEnc = iHash
		in4ipfs.Exp = in4.Exp
		result.FileOwned = append(result.FileOwned, *in4ipfs)
	}
	return result, nil
}

func (b *balancerepositories) FindIPFSHash(ipfsHashENC string) (*entities.AllowUsers, error) {
	result := new(entities.AllowUsers)
	encBytes, err := hex.DecodeString(ipfsHashENC)
	if err != nil {
		return nil, err
	}

	req, err := r.GobEncode(r.Findipfs{encBytes})
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

	err = client.Call("RPC.FindIPFS", args, res)
	if err != nil {
		log.Fatal("Call RPC:", err)
	}

	var payload r.Ipfs
	err = r.GobDecode(res.Res, &payload)
	if err != nil {
		return nil, err
	}

	resp := payload.User
	user := new(entities.User)
	fmt.Println("aaaaaaaa", ipfsHashENC)
	for address, isAuthor := range resp {
		fmt.Println("CL", address, isAuthor)
		user.Address = address
		user.Author = isAuthor
		result.Users = append(result.Users, *user)
	}
	return result, nil
}

func NewBalanceRepository() BalanceRepositories {
	return &balancerepositories{}
}
