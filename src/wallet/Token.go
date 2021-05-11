package wallet

type Token struct {
	symbol  string
	balance map[int]float64
}

func NewToken(symbol string) *Token {
	token := new(Token)
	token.symbol = symbol
	token.balance = make(map[int]float64)
	return token
}

func (token *Token) GetSymbol() string {
	return token.symbol
}

func (token *Token) BalanceOf(user *User) float64 {
	return token.balance[user.id]
}

func (token *Token) Transfer(from *User, to *User, amount float64) bool {
	if token.balance[from.id] < amount {
		return false
	}

	token.balance[from.id] -= amount
	token.balance[to.id] += amount

	return true
}

func (token *Token) Mint(to *User, amount float64) bool {
	token.balance[to.id] += amount
	return true
}

func (token *Token) Burn(from *User, amount float64) bool {
	if token.balance[from.id] < amount {
		return false
	}
	token.balance[from.id] -= amount
	return true
}
