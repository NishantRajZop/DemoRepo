package main

import "fmt"

type BankAccount struct {
	Owner   string
	balance float64
}

func (account *BankAccount) DisplayBalance() {
	fmt.Println(account.Owner)
	fmt.Println(account.balance)
}

func (account BankAccount) Deposit(amt float64) {
	account.balance += amt
	fmt.Println("Your Account was credited with ", amt)
	fmt.Println("Your total Balance Now is ", account.balance)
}

func canWithDraw(balance float64, amt float64) bool {
	return balance >= amt
}

func (account *BankAccount) withDraw(amt float64) {
	if canWithDraw(account.balance, amt) {
		account.balance -= amt
		fmt.Println("Your Account was debited with ", amt)
	} else {
		fmt.Println("Your account Doesnt Hold the Sufficient Balance  , cant WithDraw ", amt)
	}
	fmt.Println("Your total Balance Now is ", account.balance)
}
func main() {
	acc1 := BankAccount{
		"Nishant", 3000,
	}

	acc1.DisplayBalance()
	acc1.Deposit(500)
	acc1.withDraw(500)
	acc1.Deposit(4000)
	acc1.withDraw(8000)
}
