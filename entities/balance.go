package entities

type IPFSInfo struct {
	Author bool  `json:"isauthor"`
	Exp    int64 `json:"expiration"`
}
type Balance struct {
	Balanced  int                 `json:"amount"`
	FileOwned map[string]IPFSInfo `json:"filehashes"`
}

type AllowUsers struct {
	Users map[string]bool `json:"allowedusers"`
}

func NewAllowUsers() *AllowUsers {
	na := new(AllowUsers)
	na.Users = make(map[string]bool)
	return na
}
func NewBalance() *Balance {
	bl := new(Balance)
	bl.FileOwned = make(map[string]IPFSInfo)
	return bl
}
