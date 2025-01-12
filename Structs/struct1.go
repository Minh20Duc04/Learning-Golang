package main

import (
	"fmt"
	"time"
)

type AuditInfo struct{
	CreateAt time.Time
	LastModified time.Time
}

type BankAccount struct { //giống như class
	AccountNumber string
	Balance       float64
	AuditInfo 
}

func (ba BankAccount) DisplayBalance(){ //value receiver: định nghĩa phương thức để làm việc với dữ liệu, như kiểu get/set, khai báo trong notepad
	fmt.Printf("Account number: %s, Balance: %v\n", ba.AccountNumber, ba.Balance)
}

func (ba *BankAccount) Deposit(amount float64){ //void               
	ba.Balance += amount
}

type Balance1 = float64

func main() {
	account := BankAccount{AccountNumber: "123456", Balance: 1000.00}
	fmt.Println(account.Balance)

	// account2 := &BankAccount{AccountNumber : "987654", Balance : 6000.00}
	// fmt.Println(account2.AccountNumber)

	// account3 := BankAccount{
	// 	AccountNumber: "334455",
	// 	Balance: 1900.00,
	// 	AuditInfo: AuditInfo{CreateAt: time.Now(), LastModified: time.Now()},
	// }
	// fmt.Println(account3.CreateAt)

	//--------------------------------------------------//

	// account.DisplayBalance() //thao tác với dữ liệu dùng value receiver
	// account.Deposit(2000.00)
	// account.DisplayBalance()

	//--------------------------------------------------//

	var balance Balance1
	fmt.Printf("balane is %T\n", balance)
	
}