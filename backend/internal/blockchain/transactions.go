package blockchain

import (
	"HhatBridge/internal/contract"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

func CallBuyTokens(btInstance *contract.BridgeToken, privateKey string, chainID *big.Int, amount *big.Int) {
	pk := getPrivateKey(privateKey)

	auth, err := bind.NewKeyedTransactorWithChainID(pk, chainID)
	if err != nil {
		log.Println(err)
		return
	}

	auth.Value = amount

	_, err = btInstance.BuyTokens(auth)
	if err != nil {
		log.Println(err)
		return
	}
}

func CallRecieveFromBlockchain(
	btInstance *contract.BridgeToken,
	privateKey string,
	chainID *big.Int,
	to common.Address,
	amount *big.Int,
) {
	pk := getPrivateKey(privateKey)

	auth, err := bind.NewKeyedTransactorWithChainID(pk, chainID)
	if err != nil {
		log.Println(err)
		return
	}

	_, err = btInstance.RecieveFromBlockchain(auth, to, amount)
	if err != nil {
		log.Println(err)
		return
	}
}
