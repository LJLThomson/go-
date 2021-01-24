package main

import "fmt"

type Person struct {
	Name string
}

func (p *Person) Eat() {
	fmt.Println("吃东西")
}

func (p *Person) Drink() {
	fmt.Println("喝东西")
}

type Coder struct {
	Person
	Languages []string
}

func (c *Coder) UndenstandLanges() {
	c.Name = "不好"
	fmt.Println(c.Languages)
}

/**
复写方法
*/
func (c *Coder) Drink() {
	fmt.Println("喝水")
}
func main() {
	c := new(Coder)
	c.Languages = []string{"go", "java"}
	c.Eat()
	c.Drink()
	c.UndenstandLanges()
	fmt.Println(c.Name)
}
