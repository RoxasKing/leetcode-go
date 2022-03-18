package main

import "fmt"

func ExampleBank() {
	bank := Constructor([]int64{10, 100, 20, 50, 30})
	fmt.Println(bank.Withdraw(3, 10))
	fmt.Println(bank.Transfer(5, 1, 20))
	fmt.Println(bank.Deposit(5, 20))
	fmt.Println(bank.Transfer(3, 4, 15))
	fmt.Println(bank.Withdraw(10, 50))
	// Output:
	// true
	// true
	// true
	// false
	// false
}
