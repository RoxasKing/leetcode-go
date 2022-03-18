package main

// Difficulty:
// Medium

type Bank struct {
	balance []int64
	n       int
}

func Constructor(balance []int64) Bank {
	return Bank{balance: balance, n: len(balance)}
}

func (this *Bank) Transfer(acc1 int, acc2 int, money int64) bool {
	if acc1 < 1 || this.n < acc1 || acc2 < 1 || this.n < acc2 || this.balance[acc1-1] < money {
		return false
	}
	this.balance[acc1-1] -= money
	this.balance[acc2-1] += money
	return true
}

func (this *Bank) Deposit(acc int, money int64) bool {
	if acc < 1 || this.n < acc {
		return false
	}
	this.balance[acc-1] += money
	return true
}

func (this *Bank) Withdraw(acc int, money int64) bool {
	if acc < 1 || this.n < acc || this.balance[acc-1] < money {
		return false
	}
	this.balance[acc-1] -= money
	return true
}

/**
 * Your Bank object will be instantiated and called as such:
 * obj := Constructor(balance);
 * param_1 := obj.Transfer(account1,account2,money);
 * param_2 := obj.Deposit(account,money);
 * param_3 := obj.Withdraw(account,money);
 */
