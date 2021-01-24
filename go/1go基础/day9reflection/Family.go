package day9reflection

type Family struct {
	Name   string
	Age    int
	Person *Person
}

func (f *Family) GetName() string {
	return f.Name
}

func (f Family) ToString() string {
	return "Name:" + f.Name + " Age:" + string(f.Age)
}
func (f *Family) SetName(name string) {
	f.Name = name
}

