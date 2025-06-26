package main

import (
	"HhatBridge/internal/blockchain"
	"HhatBridge/internal/bridge"
	"HhatBridge/internal/contract"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalln("нет окружения аааа помогите", err)
	}

	var clientFST *ethclient.Client
	var fstInstance *contract.BridgeToken
	fstPK := preparation("URL_FST", "DEPLOYER_PRIVATE_KEY_FST", &clientFST, &fstInstance)
	defer blockchain.Disconnect(clientFST)

	var clientSCD *ethclient.Client
	var scdInstance *contract.BridgeToken
	scdPK := preparation("URL_SCD", "DEPLOYER_PRIVATE_KEY_SCD", &clientSCD, &scdInstance)
	defer blockchain.Disconnect(clientSCD)

	bridge.Bridge(fstInstance, scdInstance, fstPK, scdPK)

}

func preparation(node string, pk string, client **ethclient.Client, btInstance **contract.BridgeToken) string {
	nodeURL := os.Getenv(node)
	*client = blockchain.Connect(nodeURL)
	fmt.Println("Подключено к hardhat:", nodeURL)
	chainId := blockchain.GetChainID(*client)
	fmt.Println("Chain ID:", chainId)

	privateKey := os.Getenv(pk)
	var contractAddress common.Address
	*btInstance, contractAddress = blockchain.DeployContract(*client, privateKey)
	fmt.Println("Контракт развернут!")
	fmt.Println("Адрес контракта:", contractAddress.String())
	fmt.Println("-----------------------------------------")

	blockchain.CallBuyTokens(*btInstance, privateKey, chainId, big.NewInt(9000000000000000000))

	return privateKey
}
