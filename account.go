package main

type BankAccount interface {
	Deposit(amount float64) error
	Withdraw(amount float64) error
	GetBalance() float64
}

type Account struct {
	balance float64
}

func (account Account) Create() error {
	return nil
}

func (account *Account) Deposit(amount float64) error {
	return nil
}

func (account *Account) Withdraw(amount float64) error {
	return nil
}

func (account *Account) GetBalance() float64 {
	return account.balance
}
