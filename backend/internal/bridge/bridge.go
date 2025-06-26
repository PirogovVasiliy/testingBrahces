package bridge

import (
	"HhatBridge/internal/blockchain"
	"HhatBridge/internal/contract"
	"fmt"
)

func Bridge(fstInstance *contract.BridgeToken, scdInstance *contract.BridgeToken, fstPK string, scdPK string) {
	fmt.Println("Начинаем слушать события 1")
	fmt.Println("-----------------------------------------")
	fstChan := make(chan blockchain.TransferEvent)
	go blockchain.ListenTransfer(fstInstance, fstChan)

	fmt.Println("Начинаем слушать события 2")
	fmt.Println("-----------------------------------------")
	scdChan := make(chan blockchain.TransferEvent)
	go blockchain.ListenTransfer(scdInstance, scdChan)

	for {
		select {
		case transferEvent := <-fstChan:
			fmt.Println("Получено событие OtherBlockchainTransfer 1")
			fmt.Println("ChainID:", transferEvent.GetChainID().String())
			fmt.Println("Address:", transferEvent.GetAddress().String())
			fmt.Println("Amount:", transferEvent.GetAmount().String())
			fmt.Println("-----------------------------------------")

			blockchain.CallRecieveFromBlockchain(
				scdInstance,
				scdPK,
				transferEvent.GetChainID(),
				transferEvent.GetAddress(),
				transferEvent.GetAmount(),
			)
		case transferEvent := <-scdChan:
			fmt.Println("Получено событие OtherBlockchainTransfer 2")
			fmt.Println("ChainID:", transferEvent.GetChainID().String())
			fmt.Println("Address:", transferEvent.GetAddress().String())
			fmt.Println("Amount:", transferEvent.GetAmount().String())
			fmt.Println("-----------------------------------------")

			blockchain.CallRecieveFromBlockchain(
				fstInstance,
				fstPK,
				transferEvent.GetChainID(),
				transferEvent.GetAddress(),
				transferEvent.GetAmount(),
			)
		}
	}
}
