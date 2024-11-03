package main

type Account struct {
	balance int
}

// NewAccount is a constructor function
func NewAccount() *Account {
	return &Account{balance: 0}
}

// Method has a pointer receiver for the Account type
func (a *Account) deposit(amount int) int {
	a.balance += amount
	println(a.balance)

	return a.balance
}

func main() {
	acc := Account{balance: 100}
	acc.deposit(50)
	println(acc.balance)

	newAcc := NewAccount()
	newAcc.deposit(100)
	println(newAcc.balance)
}
