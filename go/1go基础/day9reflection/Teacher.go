package day9reflection

type Teacher struct {
	Name   string
	Age    int
	Animal Animal
}

func (t Teacher) GetName() string {
	return t.Name
}

func (t Teacher) SetName(name string) {
	t.Name = name
}
