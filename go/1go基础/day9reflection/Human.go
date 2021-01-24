package day9reflection

type Human struct {
	name   string
	age    int
	animal *Animal
}

func (h *Human) Animal() *Animal {
	return h.animal
}

func (h *Human) SetAnimal(animal *Animal) {
	h.animal = animal
}

func NewHuman(name string, age int, animal *Animal) *Human {
	return &Human{name: name, age: age, animal: animal}
}

func (h *Human) Age() int {
	return h.age
}

func (h *Human) SetAge(age int) {
	h.age = age
}

func (h *Human) Name() string {
	return h.name
}

func (h *Human) SetName(name string) {
	h.name = name
}
