package entities

type Balance struct {
	Address   string   `json:"address"`
	Balanced  int      `json:"amount"`
	FileOwned []string `json:"filehashes"`
}
