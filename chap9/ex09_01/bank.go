package main

import "fmt"

type withdraw struct {
	amount int
	success chan bool
}

var balances = make(chan int)
var deposits = make(chan int)
var withdraws = make(chan withdraw)

func Deposit(amount int) {
	deposits <- amount
}

func Balance() int {
	return <-balances
}

func Withdraw(amount int) bool {
	ch := make(chan bool)
	withdraws <- withdraw{amount, ch}
	return <- ch
}

func teller() {
	var balance = 0
	for {
		select {
		case balances <- balance:
		case amount := <- deposits:
			balance += amount
		case wd := <- withdraws:
			if wd.amount > balance {
				wd.success <- false
			} else {
				balance -= wd.amount
				wd.success <- true
			}
		}
	}
}

func main() {
	done := make(chan struct{})
	go teller()
	go func() {
		Deposit(100)
		Withdraw(150)
		fmt.Println("=", Balance())
		done <- struct{}{}
	}()
	go func() {
		Deposit(40)
		fmt.Println("=", Balance())
		done <- struct{}{}
	}()
	<-done
	<-done
	bal := Balance()
	expected := 140
	fmt.Printf("bal %d = expected %d: %v", bal, expected, bal == expected)
}