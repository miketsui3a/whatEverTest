package wallet

import (
	"fmt"
)

type WalletSystem struct {
	users     []*User
	tokens    map[string]*Token
	exchanges map[string]*Exchange
}

func NewWalletSystem() *WalletSystem {
	walletSystem := new(WalletSystem)
	walletSystem.tokens = make(map[string]*Token)
	walletSystem.exchanges = make(map[string]*Exchange)
	return walletSystem
}

func (walletSystem *WalletSystem) Init(tokens []string) {

	fmt.Println("Available token: ",tokens, " to add more token, modify Init() in Main.go")

	var tmp []*Token
	for _,token:=range(tokens){
		tokenInstance := NewToken(token)
		walletSystem.tokens[token] = tokenInstance
		tmp = append(tmp, tokenInstance)
	}

	for i:=0; i<len(tmp)-1; i++{
		for j:=i+1;j<len(tmp);j++{
			walletSystem.exchanges[tmp[i].GetSymbol()+"/"+tmp[j].GetSymbol()]=NewExchange(tmp[i], tmp[j], 1000, 1000)
		}
	}
}

func (walletSystem *WalletSystem) Run() {
	for {

		fmt.Println("1 (Deposit)\n2 (Withdraw)\n3 (Transfer)\n4 (Balance)\n5 (Swap)\n6 (Register)\n7 (Transaction Record)")
		fmt.Println("=========================")
		fmt.Print("action: ")
		var action string
		fmt.Scanln(&action)
		switch action {
		case "1":
			{
				fmt.Println("\n\n\n1 Deposit")
				user := walletSystem.GetUserById()
				if user == nil {
					continue
				}

				token := walletSystem.GetTokenBySymbol()
				if token == nil {
					continue
				}

				amount := GetAmount()
				if amount < 0 {
					continue
				}

				user.Deposit(token, amount)
				break
			}
		case "2":
			{
				fmt.Println("\n\n\n3 Withdraw")
				user := walletSystem.GetUserById()
				if user == nil {
					continue
				}

				token := walletSystem.GetTokenBySymbol()
				if token == nil {
					continue
				}

				amount := GetAmount()
				if amount < 0 {
					continue
				}

				isSuccess := user.Withdraw(token, amount)
				if !isSuccess {
					fmt.Println("not enough token in account!")
					continue
				}
				break
			}
		case "3":
			{
				fmt.Println("\n\n\n4 Transfer")
				fmt.Print("Sender ")
				sender := walletSystem.GetUserById()
				if sender == nil {
					continue
				}

				fmt.Print("receiver ")
				receiver := walletSystem.GetUserById()
				if receiver == nil {
					continue
				}

				token := walletSystem.GetTokenBySymbol()
				if token == nil {
					continue
				}

				amount := GetAmount()
				if amount < 0 {
					continue
				}

				isSuccess := sender.Transfer(token, receiver, amount)
				if !isSuccess {
					fmt.Println("not enough token in account!")
					continue
				}

				fmt.Println("Transfer Success!!!")
				break
			}
		case "4":
			{
				fmt.Println("\n\n\n4 Balance")
				user := walletSystem.GetUserById()
				if user == nil {
					continue
				}

				token := walletSystem.GetTokenBySymbol()
				if token == nil {
					continue
				}
				userBalance := user.GetBalance(token)
				balanceString := fmt.Sprintf("%f", userBalance)

				fmt.Println(token.symbol+": ", balanceString)

			}
		case "5":
			{
				fmt.Println("Swap")
				user := walletSystem.GetUserById()
				if user == nil {
					continue
				}
				fmt.Print("From ")
				from := walletSystem.GetTokenBySymbol()
				if from == nil {
					continue
				}
				fmt.Print("To ")
				to := walletSystem.GetTokenBySymbol()
				if to == nil {
					continue
				}
				exchange := walletSystem.GetExchangeByToken(from, to)
				if exchange == nil {
					continue
				}

				amount := GetAmount()
				if amount < 0 {
					continue
				}

				amountGet, isSuccess := user.Swap(exchange, from, to, amount)
				if isSuccess != true {
					continue
				}

				fmt.Println("You get: ", fmt.Sprintf("%f", amountGet), to.symbol)
				break

			}
		case "6":
			{
				fmt.Println("Register")
				user := NewUser(len(walletSystem.users))
				walletSystem.users = append(walletSystem.users, user)
				fmt.Println("user count: ", len(walletSystem.users))
			}
		case "7":
			{
				fmt.Println("Transaction Record")
				user := walletSystem.GetUserById()
				if user == nil {
					continue
				}
				for _,e :=range(user.GetTransaction()){
					if e.from == nil {
						fmt.Println("Deposit ","token: ", e.GetToken().GetSymbol()," amount: ", e.GetAmount()," time: ", e.GetTime())
						continue
					}
					if e.to ==nil{
						fmt.Println("Withdraw ","token: ", e.GetToken().GetSymbol()," amount: ", e.GetAmount()," time: ", e.GetTime())
						continue
					}
					fmt.Println("Transfer ","token: ", e.GetToken().GetSymbol()," amount: ", e.GetAmount()," from: ",e.GetFrom().id," to: ",e.GetTo().id," time: ", e.GetTime())
				}
			}
			break
		}
	}
}

func (walletSystem *WalletSystem) GetUserById() *User {
	fmt.Print("User ID: ")
	var id int
	fmt.Scanln(&id)
	if id >= len(walletSystem.users) {
		fmt.Println("user not exist")
		return nil
	}
	return walletSystem.users[id]
}

func (walletSystem *WalletSystem) GetTokenBySymbol() *Token {
	fmt.Print("Token: ")
	var tokenName string
	fmt.Scanln(&tokenName)
	if walletSystem.tokens[tokenName] == nil {
		fmt.Println("token not exist")
		return nil
	}
	return walletSystem.tokens[tokenName]
}

func (walletSystem *WalletSystem) GetExchangeByToken(from *Token, to *Token) *Exchange {
	exchangeCandidate1 := walletSystem.exchanges[from.symbol+"/"+to.symbol]
	if exchangeCandidate1 != nil {
		return exchangeCandidate1
	}
	exchangeCandidate2 := walletSystem.exchanges[to.symbol+"/"+from.symbol]
	if exchangeCandidate2 != nil {
		return exchangeCandidate2
	}
	fmt.Println("exchange not found")
	return nil
}

func GetAmount() float64 {
	fmt.Print("Amount: ")
	var amount float64
	fmt.Scanln(&amount)
	if amount < 0 {
		fmt.Println("amount cannot be negative")
	}
	return amount
}
