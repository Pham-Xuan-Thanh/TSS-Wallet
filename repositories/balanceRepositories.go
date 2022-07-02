package repositories

import (
	"bytes"
	"encoding/gob"
	"encoding/hex"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"github.com/Pham-Xuan-Thanh/TSS-Wallet/entities"
	"github.com/thanhxeon2470/TSS_chain/rpc"
)

type balancerepositories struct {
}

type BalanceRepositories interface {
	GetBalance(entities.Address) (*entities.Balance, error)
	FindIPFSHash(string) (*entities.AllowUsers, error)
}

func (b *balancerepositories) GetBalance(balanceAddr entities.Address) (*entities.Balance, error) {
	result := entities.NewBalance()
	rpc.SendGetBlance(os.Getenv("SERVER_RPC"), string(balanceAddr.Address))
	port := os.Getenv("PORT_LSRPC")
	port = fmt.Sprintf(":%s", port)
	ln, err := net.Listen("tcp", port)
	if err != nil {
		log.Panic(err)
	}
	defer ln.Close()

	conn, err := ln.Accept()
	if err != nil {
		log.Panic(err)
	}
	buff_t := rpc.HandleRPCReceive(conn)
	var payload rpc.Balance
	buff := bytes.NewBuffer(buff_t)
	dec := gob.NewDecoder(buff)
	err = dec.Decode(&payload)
	if err != nil {
		return nil, err
	}
	var balance, fileInfo = payload.Value, payload.FTXs
	result.Balanced = balance
	for iHash, in4 := range fileInfo {
		fmt.Printf("=================== %s %t %s", iHash, in4.Author, time.Unix(in4.Exp, 0))
		result.FileOwned[iHash] = entities.IPFSInfo{Author: in4.Author, Exp: in4.Exp}
	}
	return result, nil
}

func (b *balancerepositories) FindIPFSHash(ipfsHashENC string) (*entities.AllowUsers, error) {
	result := entities.NewAllowUsers()
	encBytes, err := hex.DecodeString(ipfsHashENC)
	if err != nil {
		return nil, err
	}
	rpc.SendFindIPFS(os.Getenv("SERVER_RPC"), encBytes)
	port := os.Getenv("PORT_LSRPC")
	port = fmt.Sprintf(":%s", port)
	ln, err := net.Listen("tcp", port)
	if err != nil {
		log.Panic(err)
	}
	defer ln.Close()

	conn, err := ln.Accept()
	if err != nil {
		log.Panic(err)
	}
	buff_t := rpc.HandleRPCReceive(conn)
	var payload rpc.Ipfs
	buff := bytes.NewBuffer(buff_t)
	dec := gob.NewDecoder(buff)
	err = dec.Decode(&payload)
	if err != nil {
		return nil, err
	}

	resp := payload.User
	fmt.Println("aaaaaaaa", ipfsHashENC)
	for address, isAuthor := range resp {
		fmt.Println("CL", address, isAuthor)
		result.Users[address] = isAuthor
	}
	return result, nil
}

func NewBalanceRepository() BalanceRepositories {
	return &balancerepositories{}
}
