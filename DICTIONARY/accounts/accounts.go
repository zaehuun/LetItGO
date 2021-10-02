package accounts

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
