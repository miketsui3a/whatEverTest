package wallet

import (
	"time"
)

type Transaction struct {
	token  *Token
	from   *User
	to     *User
	amount float64
	time   time.Time
}

func NewTransaction(token *Token, from *User, to *User, amount float64) Transaction {
	var transaction Transaction
	transaction.token = token
	transaction.from = from
	transaction.to = to
	transaction.amount = amount
	transaction.time = time.Now()

	return transaction
}

func (transaction *Transaction) GetToken() *Token {
	return transaction.token
}

func (transaction *Transaction) GetFrom() *User {
	return transaction.from
}

func (transaction *Transaction) GetTo() *User {
	return transaction.to
}

func (transaction *Transaction) GetAmount() float64 {
	return transaction.amount
}

func (transaction *Transaction) GetTime() time.Time {
	return transaction.time
}
