package main

import (
	"fmt"
	"sync"
	"time"
)

type Account struct {
	ID      int
	Balance int
	mu      sync.Mutex
}

func Transfer(sender, receiver *Account, amount int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Initiating transfer of %d from Account %d to Account %d\n", amount, sender.ID, receiver.ID)

	sender.mu.Lock()
	fmt.Printf("Locked Account %d\n", sender.ID)
	time.Sleep(1 * time.Second)

	receiver.mu.Lock()
	fmt.Printf("Locked Account %d\n", receiver.ID)

	if sender.Balance >= amount {
		sender.Balance -= amount
		receiver.Balance += amount
		fmt.Printf("Transferred %d from Account %d to Account %d\n", amount, sender.ID, receiver.ID)
	} else {
		fmt.Println("Insufficient balance")
	}

	receiver.mu.Unlock()
	sender.mu.Unlock()
}

func SafeTransfer(sender, receiver *Account, amount int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Initiating safe transfer of %d from Account %d to Account %d\n", amount, sender.ID, receiver.ID)

	if sender.ID < receiver.ID {
		sender.mu.Lock()
		fmt.Printf("Locked Account %d\n", sender.ID)
		time.Sleep(1 * time.Second)

		receiver.mu.Lock()
		fmt.Printf("Locked Account %d\n", receiver.ID)
	} else {
		receiver.mu.Lock()
		fmt.Printf("Locked Account %d\n", receiver.ID)
		time.Sleep(1 * time.Second)

		sender.mu.Lock()
		fmt.Printf("Locked Account %d\n", sender.ID)
	}

	if sender.Balance >= amount {
		sender.Balance -= amount
		receiver.Balance += amount
		fmt.Printf("Transferred %d from Account %d to Account %d\n", amount, sender.ID, receiver.ID)
	} else {
		fmt.Println("Insufficient balance")
	}

	receiver.mu.Unlock()
	sender.mu.Unlock()
}

func main() {
	account1 := &Account{ID: 1, Balance: 1000}
	account2 := &Account{ID: 2, Balance: 1000}

	var wg sync.WaitGroup

	fmt.Println("Demonstrating deadlock scenario:")
	wg.Add(2)
	go Transfer(account1, account2, 100, &wg)
	go Transfer(account2, account1, 50, &wg)
	wg.Wait()

	//safe transfer

	// fmt.Println("\nDemonstrating no-deadlock scenario:")
	// wg.Add(2)
	// go SafeTransfer(account1, account2, 100, &wg)
	// go SafeTransfer(account2, account1, 50, &wg)
	// wg.Wait()
}
