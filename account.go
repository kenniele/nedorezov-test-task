package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
)

type BankAccount interface {
	Deposit(amount float64) error
	Withdraw(amount float64) error
	GetBalance() float64
}

type Account struct {
	ID      int     `json:"ID,-"`
	Balance float64 `json:"-"`
}

type Operation struct {
	AccountID int     `json:"ID"`
	Amount    float64 `json:"Amount"`
}

type Result struct {
	ID      int
	Balance float64
	Error   error
}

func (account Account) Create() (int, error) {
	row := connection.QueryRow(`SELECT * FROM "Account" WHERE "id" = $1`, account.ID)
	if err := row.Scan(&account.ID, &account.Balance); errors.Is(err, sql.ErrNoRows) {
		row = connection.QueryRow(`INSERT INTO "Account" ("Balance") VALUES (0) RETURNING "id"`)
		if err := row.Scan(&account.ID); err != nil {
			log.Println("Error while inserting account", err)
			return 0, errors.New("ошибка во время вставки в БД")
		}
	} else if err != nil {
		log.Println("Error while selecting account", err)
		return 0, errors.New("ошибка во время запроса к БД")
	}
	log.Println("Account created")
	return account.ID, nil

}

func (account *Account) isIn() error {
	row := connection.QueryRow(`SELECT * FROM "Account" WHERE "id"=$1`, account.ID)
	if err := row.Scan(&account.ID, &account.Balance); err != nil {
		return errors.New("аккаунт не найден")
	}
	return nil
}

func (account *Account) Select(operation Operation) error {
	row := connection.QueryRow(`SELECT * FROM "Account" WHERE "id"=$1`, operation.AccountID)
	if err := row.Scan(&account.ID, &account.Balance); err != nil {
		return err
	}
	return nil
}

func (account *Account) Deposit(amount float64) error {
	operation := Operation{AccountID: account.ID, Amount: amount}
	row := connection.QueryRow(`SELECT * FROM "Account" WHERE "id"=$1`, operation.AccountID)
	if err := row.Scan(&account.ID, &account.Balance); errors.Is(err, sql.ErrNoRows) {
		log.Println("Account not found")
		return errors.New("аккаунт не найден")
	}
	account.Balance += amount
	if _, err := connection.Exec(`UPDATE "Account" SET "Balance"=$1 WHERE "id"=$2`, account.Balance, operation.AccountID); err != nil {
		log.Println("Error while updating table", err)
		return errors.New("ошибка во время обновления таблицы")
	}
	return nil
}

func (account *Account) Withdraw(amount float64) error {
	operation := Operation{AccountID: account.ID, Amount: amount}
	fmt.Println(operation)
	row := connection.QueryRow(`SELECT * FROM "Account" WHERE "id"=$1`, operation.AccountID)
	if err := row.Scan(&account.ID, &account.Balance); errors.Is(err, sql.ErrNoRows) {
		log.Println("Account not found")
		return errors.New("аккаунт не найден")
	}
	if account.Balance < operation.Amount {
		log.Println("There's no enough balance")
		return errors.New("недостаточно средств")
	}
	account.Balance -= operation.Amount
	if _, err := connection.Exec(`UPDATE "Account" SET "Balance"=$1 WHERE "id"=$2`, account.Balance, operation.AccountID); err != nil {
		log.Println("Error while updating table", err)
		return errors.New("ошибка во время обновления таблицы")
	}
	return nil
}

func (account *Account) GetBalance() Result {
	row := connection.QueryRow(`SELECT "Balance" FROM "Account" WHERE "id"=$1`, account.ID)
	err := row.Scan(&account.Balance)
	if err != nil {
		return Result{Error: errors.New("аккаунт не найден")}
	}
	return Result{ID: account.ID, Balance: account.Balance, Error: nil}
}
