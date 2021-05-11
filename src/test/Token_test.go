package test

import (
	"testing"

	"../wallet"
)


func TestBalanceMintBurn(t *testing.T){
	dai := wallet.NewToken("DAI")
	user0 := wallet.NewUser(0)
	isSuccess:=dai.Mint(user0, 1000)
	if !isSuccess{
		t.Errorf("mint fail")
	}

	user0Balance:=dai.BalanceOf(user0)
	if user0Balance!= 1000{
		t.Errorf("get balance after mint incorret")
	}

	isSuccess=dai.Burn(user0, 100)
	if !isSuccess{
		t.Errorf("burn fail")
	}

	user0Balance=dai.BalanceOf(user0)
	if user0Balance!=900{
		t.Errorf("get balance after burn incorret")
	}

}

func TestBalanceTransfer(t *testing.T){
	dai := wallet.NewToken("DAI")
	user0 := wallet.NewUser(0)
	user1 := wallet.NewUser(1)

	isSuccess:=dai.Mint(user0, 1000)
	if !isSuccess{
		t.Errorf("mint fail")
	}

	isSuccess=dai.Transfer(user0,user1,2000)
	if isSuccess {
		t.Errorf("transfer amount should not greater than user have")
	}

	isSuccess=dai.Transfer(user0,user1,500)
	if !isSuccess {
		t.Errorf("transfer fail")
	}

	user0Balance:=dai.BalanceOf(user0)
	if user0Balance!=500{
		t.Errorf("transfer from amount fail")
	}

	user1Balance:=dai.BalanceOf(user1)
	if user1Balance!=500{
		t.Errorf("transfer to amount fail")
	}

}