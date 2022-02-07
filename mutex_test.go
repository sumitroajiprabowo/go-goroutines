package go_goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestUsingMutex(t *testing.T) {
	x := 0
	var mutex sync.Mutex
	for i := 0; i < 1000; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				mutex.Lock()
				x++
				mutex.Unlock()
			}
		}()
	}
	time.Sleep(5 * time.Second)
	fmt.Println("Counter: ", x)
}

type BankAccount struct {
	RWMutex sync.RWMutex
	Balance int
}

func (account *BankAccount) AddBalance(amount int) {
	account.RWMutex.Lock()
	defer account.RWMutex.Unlock()
	account.Balance = account.Balance + amount
}

func (account *BankAccount) GetBalance() int {
	account.RWMutex.RLock()
	defer account.RWMutex.RUnlock()
	return account.Balance
}

func TestReadWriteMutex(t *testing.T) {
	account := BankAccount{}
	for i := 0; i < 100; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				account.AddBalance(1)
				fmt.Println("Balance:", account.GetBalance())
			}
		}()
	}
	time.Sleep(5 * time.Second)
	fmt.Println("Final Balance", account.GetBalance())
}

type UserBalance struct {
	sync.Mutex
	Name    string
	Balance int
}

func (user *UserBalance) Lock() {
	user.Mutex.Lock()
}

func (user *UserBalance) Unlock() {
	user.Mutex.Unlock()
}

func (user *UserBalance) Change(amount int) {
	user.Balance = user.Balance + amount
}

func Transfer(user1 *UserBalance, user2 *UserBalance, ammount int) {
	user1.Lock()
	fmt.Println("Lock User 1", user1.Name)
	user1.Change(-ammount)

	time.Sleep(1 * time.Second)

	user2.Lock()
	fmt.Println("Lock User 2", user2.Name)
	user2.Change(+ammount)

	time.Sleep(1 * time.Second)

	user1.Unlock()
	fmt.Println("Unlock", user1.Name)
	user2.Unlock()
	fmt.Println("Unlock", user2.Name)
}

func TestDeadlock(t *testing.T) {
	user1 := UserBalance{Name: "Budi", Balance: 100000}
	user2 := UserBalance{Name: "Danu", Balance: 100000}

	go Transfer(&user1, &user2, 10000)
	go Transfer(&user2, &user1, 20000)

	time.Sleep(5 * time.Second)

	fmt.Println("Final Balance", user1.Name, user2.Balance)
	fmt.Println("Final Balance", user2.Name, user1.Balance)
}
