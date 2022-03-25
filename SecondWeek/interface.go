package main

import (
	"errors"
	"fmt"
)

var ErrNoSufficicientBalance = errors.New("Yeterli bakiyeniz bulunmamaktadır")

func main() {
	b := BankAccount{"12345", 100}
	err := b.Withdraw(150)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Hesabınızdan para çekimi başarıyla tamamlandı")
}

type BankAccount struct {
	AccountNumber string
	Balance       int
}

func (b *BankAccount) Withdraw(a int) error {
	if b.Balance < a {
		return ErrNoSufficicientBalance
	}
	b.Balance -= a
	return nil
}
