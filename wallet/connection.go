package wallet

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

func EthConnect() bool {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	url := os.Getenv("INFURA_URL")

	fmt.Println(url)

	client, err := ethclient.Dial("wss://rpc-mainnet.matic.network")

	if err != nil {
		fmt.Println(err)

		return false
	}
	_ = client

	//so now connected so we wqant tor ead somehting.
	account := common.HexToAddress("0xc94d737b36A32BbC4eaf545832C05420fa11B916")
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(balance)
	return true
}
