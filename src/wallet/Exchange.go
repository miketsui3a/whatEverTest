package wallet

type Exchange struct {
	token1       *Token
	token2       *Token
	token1Amount float64
	token2Amount float64
	invariant    float64
}

func NewExchange(token1 *Token, token2 *Token, token1Amount float64, token2Amount float64) *Exchange {
	exchange := new(Exchange)
	exchange.token1 = token1
	exchange.token2 = token2
	exchange.token1Amount = token1Amount
	exchange.token2Amount = token2Amount
	exchange.invariant = token1Amount * token2Amount
	return exchange
}

func (exchange *Exchange) GetToken1() *Token {
	return exchange.token1
}

func (exchange *Exchange) GetToken2() *Token {
	return exchange.token2
}

func (exchange *Exchange) GetToken1Amount() float64 {
	return exchange.token1Amount
}

func (exchange *Exchange) GetToken2Amount() float64 {
	return exchange.token2Amount
}

func (exchange *Exchange) GetInvariant() float64 {
	return exchange.invariant
}

func (exchange *Exchange) Swap(from *Token, amount float64) (float64, bool) {

	if from == exchange.token1 {
		token1AfterAmount := exchange.token1Amount + amount
		token2AfterAmount := exchange.invariant / token1AfterAmount
		swapAmount := exchange.token2Amount - token2AfterAmount

		exchange.token1Amount = token1AfterAmount
		exchange.token2Amount = token2AfterAmount

		return swapAmount, true
	} else if from == exchange.token2 {
		token2AfterAmount := exchange.token2Amount + amount
		token1AfterAmount := exchange.invariant / token2AfterAmount
		swapAmount := exchange.token1Amount - token1AfterAmount

		exchange.token1Amount = token1AfterAmount
		exchange.token2Amount = token2AfterAmount

		return swapAmount, true
	}

	return 0, false

}
