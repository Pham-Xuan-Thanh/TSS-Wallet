package entities

type IPFSInfo struct {
	IpfsEnc string `json:"ipfs_enc"`
	Author  bool   `json:"isauthor"`
	Exp     int64  `json:"expiration"`
}
type Balance struct {
	Balanced  int        `json:"amount"`
	FileOwned []IPFSInfo `json:"filehashes"`
}

type User struct {
	Address string `json:"adrress"`
	Author  bool   `json:"isauthor"`
}

type AllowUsers struct {
	Users []User `json:"allowedusers"`
}

// func NewAllowUsers() *AllowUsers {
// 	na := new(AllowUsers)
// 	na.Users = make(map[string]bool)
// 	return na
// }
// func NewBalance() *Balance {
// 	bl := new(Balance)
// 	bl.FileOwned = make(map[string]IPFSInfo)
// 	return bl
// }
