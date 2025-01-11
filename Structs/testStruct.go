package main

import "fmt"

/*
Create withdraw method on BankAccount
It should take a float amount
Subtract it from the balance if sufficient
*/

type BankAccount1 struct{ 
	accountId string
	balance float64
}

type Customer struct{
	name string
	accounts []BankAccount1
}

func (ba *BankAccount1) WithDraw(amount float64){ //thay đổi giá trị gốc thì dùng con trỏ, muốn chắc thì cứ dùng con trỏ cho tất cả method
	if amount < 1.00 || ba.balance < amount{
		fmt.Println("\nInvalid action. Please check again")
	}else{
		ba.balance -= amount
		fmt.Println("\nWithdraw action successfully occur")
	}
}

func (ba BankAccount1) DisplayInfo(){
	fmt.Printf("id: %v, balance: %v", ba.accountId, ba.balance)
}

func (cus *Customer) addAccount (ba BankAccount1) { //thêm account ngoài vào []accounts của Customer
	cus.accounts = append(cus.accounts, ba)
}

func (cus *Customer) displayAccouts() { //xuất tất cả Account dựa trên Customer
	for _, c := range cus.accounts {
		c.DisplayInfo()
		fmt.Println()
	}
}

func main(){
	myAccount := BankAccount1{
		accountId: "UID1",
		balance: 1000.00,
	}

	// myAccount.DisplayInfo() 
	// myAccount.WithDraw(300.00) //test withdraw
	// myAccount.DisplayInfo() 


	myCustomer := Customer{
		name: "Nguyen Minh Duc",
		accounts: []BankAccount1{
			{
				accountId : "UID2",
				balance: 2000.00,
			},

			{
				accountId : "UID3",
				balance: 3000.00,
			},
		},
	}
	myCustomer.displayAccouts() //in tất cả Account dựa trên myCustomer

	myCustomer.addAccount(myAccount)

	myCustomer.displayAccouts() //in tất cả Account sau khi thêm 1 Account

}