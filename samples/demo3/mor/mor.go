package mor

import "fmt"

type Hero interface {
	GetAge() int
}

type MoRer struct {
	Name  string
	Age   int
	Email string
}

type Gopher struct {
	name string
	age  int
}

func NewGopher(name string, age int) Gopher {
	return Gopher{name, age}
}

func (g Gopher) GetName() string {
	return g.name
}

func (g Gopher) Talk() {
	fmt.Printf("Hello, my name is %s, and I'm this old: %v\n", g.name, g.age)
}

func (g Gopher) GetAge() int {
	return g.age
}

func (m MoRer) GetAge() int {
	return m.Age
}
