package accounts

import "errors"

// Acoount struct
type Account struct {
	owner   string
	balance int
}

// NewAccont creates Account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

//리시버
// Deposit x amount on your account
func (a *Account) Deposit(amount int) {
	a.balance += amount
}

// Withdraw x amount from your account
// go에는 exception 없음
var errNoMoney = errors.New("Can't withdraw you are poor")

func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		// return errors.New("Can't withdraw you are poor")
		return errNoMoney
	}
	a.balance -= amount
	return nil
}

// Balance of your account
func (a Account) Balance() int {
	return a.balance
}
