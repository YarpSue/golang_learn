package main

import "fmt"

type Inter interface {
	Ping()
	Pang()
}

type Anter interface {
	Inter
	String()
}

type St struct {
	Name string
}

func (St) Ping() {
	println("ping")
}

func (*St) Pang() {
	println("pang")
}

func main() {
	st := &St{"andes"}
	var i interface{} = st
	// 判断i绑定的实例是否实现了接口类型Inter
	if o, ok := i.(Inter); ok {
		o.Ping()
		o.Pang()
	}

	if p, ok := i.(Anter); ok {
		p.String()
	}

	// 对实例 i 直接类型的判断
	if s, ok := i.(St); ok {
		fmt.Printf("%s", s.Name)
	}

	if x, ok := i.(*St); ok {
		fmt.Printf("*St %s ", x.Name)
	}

}