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
	GetBalance(string) (*entities.Balance, error)
	FindIPFSHash(string) (*entities.AllowUsers, error)
}

func (b *balancerepositories) GetBalance(addrTSS string) (*entities.Balance, error) {
	result := new(entities.Balance)
	rpc.SendGetBlance(os.Getenv("SERVER_RPC"), addrTSS)
	port := os.Getenv("PORT_LSRPC")
	port = fmt.Sprintf(":%s", port)
	ln, err := net.Listen("tcp", port)
	if err != nil {
		log.Panic(err)
	}
	defer ln.Close()
	var buff_t []byte
	deadline := time.Now().Add(time.Second * 30)
	for {
		conn, err := ln.Accept()
		conn.SetDeadline(deadline)
		if err != nil {
			log.Panic(err)
		}
		var command string
		buff_t, command = rpc.HandleRPCReceive(conn)
		if command == "balance" {
			break
		}
		if time.Now().Unix() > deadline.Unix() {
			return nil, fmt.Errorf("RPC timeout")
		}
	}

	var payload rpc.Balance
	buff := bytes.NewBuffer(buff_t)
	dec := gob.NewDecoder(buff)
	err = dec.Decode(&payload)
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
	rpc.SendFindIPFS(os.Getenv("SERVER_RPC"), encBytes)
	port := os.Getenv("PORT_LSRPC")
	port = fmt.Sprintf(":%s", port)
	ln, err := net.Listen("tcp", port)
	if err != nil {
		log.Panic(err)
	}
	defer ln.Close()

	var buff_t []byte
	deadline := time.Now().Add(time.Second * 30)
	for {
		conn, err := ln.Accept()
		conn.SetDeadline(deadline)
		if err != nil {
			log.Panic(err)
		}
		var command string
		buff_t, command = rpc.HandleRPCReceive(conn)
		if command == "ipfs" {
			break
		}
		if time.Now().Unix() > deadline.Unix() {
			return nil, fmt.Errorf("RPC timeout")
		}
	}
	var payload rpc.Ipfs
	buff := bytes.NewBuffer(buff_t)
	dec := gob.NewDecoder(buff)
	err = dec.Decode(&payload)
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
