package main

import (
	"./wallet"
)

func main() {
	walletSystem := wallet.NewWalletSystem()
	walletSystem.Init([]string{"USDT","DAI","ETH"})
	walletSystem.Run()
}
