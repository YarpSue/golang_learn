package main

import "fmt"

type person struct {
	name string
	age int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	return &p
}

func main() {
	/*
	fmt.Println(person{"Bob", 12})
	fmt.Println(person{name: "Alick", age: 20})
	fmt.Println(person{name: "xiaoming"})
	fmt.Println(&person{name: "Ann", age: 40})
	fmt.Println(newPerson("John"))
*/
	s := person{name: "Season", age: 50}
	fmt.Println(s.name, s.age)
	sp := &s
	fmt.Println(sp.name)
	sp.age = 55
	fmt.Println(sp.age)

}