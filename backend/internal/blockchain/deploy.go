package blockchain

import (
	"HhatBridge/internal/contract"
	"context"
	"crypto/ecdsa"
	"log"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func DeployContract(client *ethclient.Client, privateKey string) (btInstance *contract.BridgeToken, contractAddress common.Address) {
	pk := getPrivateKey(privateKey)

	auth, err := bind.NewKeyedTransactorWithChainID(pk, GetChainID(client))
	if err != nil {
		log.Panic(err)
	}

	contractAddress, tx, btInstance, err := contract.DeployBridgeToken(auth, client)
	if err != nil {
		log.Panic(err)
	}

	_, err = bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		log.Panic(err)
	}
	return
}

func getPrivateKey(privateKeyStr string) *ecdsa.PrivateKey {
	pk, err := crypto.HexToECDSA(privateKeyStr)
	if err != nil {
		log.Panic(err)
	}
	return pk
}
