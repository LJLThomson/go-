package main

import (
	"fmt"
	"sync"
	"time"
)

type Person struct {
	money int
	mt    sync.Mutex
}

func (p *Person) Money() int {
	return p.money
}

func (p *Person) SetMoney(money int) {
	p.money = money
}

func (p *Person) SaveMoney(num int) {
	p.mt.Lock()
	p.money -= num
	<-time.After(3 * time.Second)
	p.mt.Unlock()
}

func (p *Person) TakeOutMoney(num int) {
	p.mt.Lock()
	p.money += num
	<-time.After(3 * time.Second)
	p.mt.Unlock()
}

func NewPerson(money int, mt sync.Mutex) *Person {
	return &Person{money: money, mt: mt}
}

func main() {
	var mt sync.Mutex
	var wg sync.WaitGroup
	person := NewPerson(100, mt)
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			person.SaveMoney(100)
			wg.Done()
		}()
	}
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			person.TakeOutMoney(100)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("总额度", person.Money())
	fmt.Println("main over")
}
