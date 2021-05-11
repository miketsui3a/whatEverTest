package wallet

import "fmt"

type User struct {
	id           int
	transactions []Transaction
}

func NewUser(id int) *User {
	user := new(User)
	user.id = id
	return user
}

func (user *User) AddTransaction(transaction Transaction) {
	user.transactions = append(user.transactions, transaction)
}

func (from *User) Transfer(token *Token, to *User, amount float64) bool {
	isSuccess := token.Transfer(from, to, amount)
	if !isSuccess {
		return false
	}
	transaction := NewTransaction(token, from, to, amount)
	from.AddTransaction(transaction)
	to.AddTransaction(transaction)
	return true
}

func (user *User) Deposit(token *Token, amount float64) bool {
	isSuccess := token.Mint(user, amount)
	if !isSuccess {
		return false
	}
	transaction := NewTransaction(token, nil, user, amount)
	user.AddTransaction(transaction)
	return true
}

func (user *User) Withdraw(token *Token, amount float64) bool {
	isSuccess := token.Burn(user, amount)
	if !isSuccess {
		return false
	}
	transaction := NewTransaction(token, user, nil, amount)
	user.AddTransaction(transaction)
	return true
}

func (user *User) GetBalance(token *Token) float64 {
	return token.BalanceOf(user)
}

func (user *User) GetTransaction() []Transaction {
	return user.transactions
}


func (user *User) Swap(exchange *Exchange, from *Token, to *Token, amount float64) (float64, bool) {
	if user.GetBalance(from) < amount {
		fmt.Println("not enough balance")
		return 0, false
	}

	amountGet, isSuccess := exchange.Swap(from, amount)
	if !isSuccess {
		fmt.Println("swap failed")
		return 0, false
	}

	user.Withdraw(from, amount)

	isSuccess = user.Deposit(to, amountGet)
	if !isSuccess {
		fmt.Println("deposit")
		return 0, false
	}

	return amountGet, true

}