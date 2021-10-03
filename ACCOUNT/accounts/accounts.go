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

// Change Owner of the account
func (a *Account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}

// Owner of the account
func (a Account) Owner() string {
	return a.owner
}

// Balance of your account
func (a Account) Balance() int {
	return a.balance
}

// struct를 출력하면 내부적으로 이걸 자동 호출해서 출력함
// fmt.Println(account) 하면 hello 출력 됨
func (a Account) String() string {
	return "hello"
}
