package main

// import (
// 	"fmt"
// 	"log"
// 	"os"
// 	"time"

// 	"github.com/joho/godotenv"
// 	"github.com/thanhxeon2470/TSS_chain/cli"
// )

// func main() {

// 	err := godotenv.Load(".env")

// 	if err != nil {
// 		log.Fatalf("Error loading .env file")
// 	}
// 	var core = cli.CLI{}

// 	go func(blkchain cli.CLI) {
// 		blkchain.StartNode("")
// 		time.Sleep(time.Second * 10)
// 		return
// 	}(core)
// 	fmt.Println(os.Getenv("KNOWNNODE"))
// 	var allow []string
// 	allow = append(allow, "16Jcx92y5gv1vfVfzqFkghqXYSCdTkmgV7")
// 	core.SendProposal("7jczDSiSxs6ZNzQdTT6TydNHgNLKBNLjPWzJ2JHkiX", "16Jcx92y5gv1vfVfzqFkghqXYSCdTkmgV7", 1, allow, "QmX25bgRi5Niby6PFzNPwW67Ch8XfJHMdcZqjDMhDwWdNZ")
// 	time.Sleep(31 * time.Second)
// }
