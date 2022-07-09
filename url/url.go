package main

import "fmt"

type Values map[string][]string

type IntList struct {
	Value int
	Tail *IntList
}

func (v Values) Get(key string) string {
	if vs := v[key]; len(vs) > 0 {
		return vs[0]
	}
	return ""
}

func (v Values) Add(key, value string) {
	v[key] = append(v[key], value)
}

func (list *IntList) Sum() int {
	if list == nil {
		return 0
	}
	return list.Value + list.Tail.Sum()
}

func main() {
	//m := Values{"lang": {"en"}}
	//m.Add("item", "1")
	//m.Add("item", "2")
	////
	////fmt.Println(m.Get("lang"))
	////fmt.Println(m.Get("q"))
	////fmt.Println(m.Get("item"))
	////fmt.Println(m["item"])
	//
	//m = nil
	//fmt.Println(m)
	//fmt.Println(m.Get("lang"))
	//m.Add("a", "1")

	l := IntList{1, &IntList{1, nil}}
	fmt.Println(l.Sum())
}