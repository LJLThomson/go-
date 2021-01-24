package main

import (
	"fmt"
	"sync"
)

type person struct {
	Name  string
	Alive bool
}

func (p *person) Kill() {
	fmt.Println("杀死", p.Name)
	if p.Alive {
		p.Alive = false
	}
}

/**
Once ,不管哪个协程，方法只执行一次
*/
func main() {
	var wg sync.WaitGroup
	var once sync.Once
	p := &person{"比尔", true}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			// 不管哪个协程，只执行一次
			once.Do(func() {
				p.Kill()
			})
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("main over")
}
