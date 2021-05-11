package test

import (
	"testing"

	"../wallet"
)

func TestDepositTransactionRecord(t *testing.T) {
	user0 := wallet.NewUser(0)
	usdt := wallet.NewToken("USDT")
	user0.Deposit(usdt, 100)

	txs := user0.GetTransaction()

	if txs[0].GetToken() != usdt {
		t.Errorf("deposit transaction token error")
	}

	if txs[0].GetFrom() != nil {
		t.Errorf("deposit transaction from error")
	}

	if txs[0].GetTo() != user0 {
		t.Errorf("deposit transaction to error")
	}

	if txs[0].GetAmount() != 100 {
		t.Errorf("deposit transaction amount error")
	}

}

func TestWithdrawTransactionRecord(t *testing.T) {
	user0 := wallet.NewUser(0)
	usdt := wallet.NewToken("USDT")
	user0.Deposit(usdt, 100)
	user0.Withdraw(usdt, 50)

	txs := user0.GetTransaction()

	if txs[1].GetToken() != usdt {
		t.Errorf("withdraw transaction token error")
	}

	if txs[1].GetFrom() != user0 {
		t.Errorf("withdraw transaction from error")
	}

	if txs[1].GetTo() != nil {
		t.Errorf("withdraw transaction to error")
	}

	if txs[1].GetAmount() != 50 {
		t.Errorf("withdraw transaction amount error")
	}
}

func TestTransferTransactionRecord(t *testing.T) {
	user0 := wallet.NewUser(0)
	user1 := wallet.NewUser(1)
	usdt := wallet.NewToken("USDT")
	user0.Deposit(usdt, 100)

	user0.Transfer(usdt, user1, 50)

	user0txs := user0.GetTransaction()
	user1txs := user1.GetTransaction()

	if user0txs[1].GetToken() != usdt {
		t.Errorf("transfer transaction token error")
	}

	if user0txs[1].GetFrom() != user0 {
		t.Errorf("transfer transaction from error")
	}

	if user0txs[1].GetTo() != user1 {
		t.Errorf("transfer transaction to error")
	}

	if user0txs[1].GetAmount() != 50 {
		t.Errorf("transfer transaction amount error")
	}

	if user1txs[0].GetToken() != usdt {
		t.Errorf("transfer transaction token error")
	}

	if user1txs[0].GetFrom() != user0 {
		t.Errorf("transfer transaction from error")
	}

	if user1txs[0].GetTo() != user1 {
		t.Errorf("transfer transaction to error")
	}

	if user1txs[0].GetAmount() != 50 {
		t.Errorf("transfer transaction amount error")
	}
}

func TestSwapTransactionRecord(t *testing.T) {
	user0 := wallet.NewUser(0)
	dai := wallet.NewToken("DAI")
	usdt := wallet.NewToken("USDT")
	daiUsdtExchange := wallet.NewExchange(dai, usdt, 10000,10000)
	user0.Deposit(usdt, 100)
	user0.Swap(daiUsdtExchange, usdt, dai, 100)

	txs := user0.GetTransaction()

	if txs[1].GetToken() != usdt {
		t.Errorf("swap transaction token error")
	}

	if txs[1].GetFrom() != user0 {
		t.Errorf("swap transaction from error")
	}

	if txs[1].GetTo() != nil {
		t.Errorf("swap transaction to error")
	}

	if txs[1].GetAmount() != 100 {
		t.Errorf("swap transaction amount error")
	}

	if txs[2].GetToken() != dai {
		t.Errorf("swap transaction token error")
	}

	if txs[2].GetFrom() != nil {
		t.Errorf("swap transaction from error")
	}

	if txs[2].GetTo() != user0 {
		t.Errorf("swap transaction to error")
	}

	if txs[2].GetAmount() != 99.00990099009869 {
		t.Errorf("swap transaction amount error")
	}
}