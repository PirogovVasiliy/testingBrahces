package blockchain

import (
	"context"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/ethclient"
)

func Connect(nodeURL string) (client *ethclient.Client) {
	client, err := ethclient.Dial(nodeURL)
	if err != nil {
		log.Panic(err)
	}
	return
}

func Disconnect(client *ethclient.Client) {
	client.Close()
}

func GetChainID(client *ethclient.Client) *big.Int {
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		log.Panic(err)
	}
	return chainID
}
