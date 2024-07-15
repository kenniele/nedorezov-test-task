package main

import (
	"database/sql"
	"errors"
	"log"
)

type BankAccount interface {
	Deposit(amount float64) error
	Withdraw(amount float64) error
	GetBalance() float64
}

type Account struct {
	ID      int     `json:"ID,-"`
	Balance float64 `json:"Balance"`
}

func (account Account) Create() (int, error) {
	row := connection.QueryRow(`SELECT * FROM "Account" WHERE "id" = $1`, account.ID)

	if err := row.Scan(&account.ID, &account.Balance); errors.Is(err, sql.ErrNoRows) {
		row = connection.QueryRow(`INSERT INTO "Account" ("Balance") VALUES (0)`)
		if err := row.Scan(&account.ID, &account.Balance); err != nil {
			return 0, err
		}
	} else if err != nil {
		log.Println(err)
		log.Println("Account already exists")
		return account.ID, nil
	}
	log.Println("Account created")
	return account.ID, nil

}

func (account *Account) Deposit(amount float64) error {
	return nil
}

func (account *Account) Withdraw(amount float64) error {
	return nil
}

func (account *Account) GetBalance() float64 {
	return 0
}
