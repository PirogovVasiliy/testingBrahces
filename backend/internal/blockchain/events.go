package blockchain

import (
	"HhatBridge/internal/contract"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

type TransferEvent struct {
	chainID uint32
	address common.Address
	amount  *big.Int
}

func (TE *TransferEvent) GetChainID() *big.Int {
	num := new(big.Int)
	num.SetUint64(uint64(TE.chainID))
	return num
}

func (TE *TransferEvent) GetAddress() common.Address { return TE.address }

func (TE *TransferEvent) GetAmount() *big.Int { return TE.amount }

func ListenTransfer(btInstance *contract.BridgeToken, outChanal chan<- TransferEvent) {

	transferChanal := make(chan *contract.BridgeTokenOtherBlockchainTransfer)

	sub, err := btInstance.WatchOtherBlockchainTransfer(&bind.WatchOpts{}, transferChanal)
	if err != nil {
		log.Panic(err)
	}
	defer sub.Unsubscribe()

	for {
		select {
		case err := <-sub.Err():
			log.Fatalln("Ошибка чтения события OtherBlockchainTransfer!", err)
		case event := <-transferChanal:
			outChanal <- TransferEvent{event.ChainID, event.To, event.Amount}
		}
	}
}
