package test

import (
	"testing"

	"../wallet"
)

func TestInvariant(t *testing.T) {
	dai := wallet.NewToken("DAI")
	usdt := wallet.NewToken("USDT")
	daiUsdtExchange := wallet.NewExchange(dai, usdt, 10000,10000)
	if daiUsdtExchange.GetInvariant() != 10000*10000{
		t.Error("invariant error")
	}
}

func TestSwap(t *testing.T){
	dai := wallet.NewToken("DAI")
	usdt := wallet.NewToken("USDT")
	daiUsdtExchange := wallet.NewExchange(dai, usdt, 10000,10000)
	var fakeToken *wallet.Token
	_, isSuccess := daiUsdtExchange.Swap(fakeToken, 1000)
	if isSuccess{
		t.Error("token swap check fail")
	}

	amountGet, isSuccess := daiUsdtExchange.Swap(dai, 1000)
	if !isSuccess{
		t.Error("token swap check fail")
	}

	if amountGet != 909.0909090909099{
		t.Error("token swap amount fail")
	}

	if daiUsdtExchange.GetToken1Amount()!=11000{
		t.Error("token1 after amount fail")
	}

	if daiUsdtExchange.GetToken2Amount()!=10000-909.0909090909099{
		t.Error("token2 after amount fail")
	}

}