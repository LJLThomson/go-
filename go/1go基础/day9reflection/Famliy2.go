package day9reflection

type Family2 struct {
	name   string
	age    int
	person *Person
}

func (f *Family2) Person() *Person {
	return f.person
}

func (f *Family2) SetPerson(person *Person) {
	f.person = person
}

func NewFamily2(name string, age int, person *Person) *Family2 {
	return &Family2{name: name, age: age, person: person}
}

func (f *Family2) Age() int {
	return f.age
}

func (f *Family2) SetAge(age int) {
	f.age = age
}

func (f *Family2) Name() string {
	return f.name
}

func (f *Family2) SetName(name string) {
	f.name = name
}
