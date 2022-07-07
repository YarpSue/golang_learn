package main

import (
	"fmt"
	"math"
)

type Point struct{ X,Y float64}

//func Distance(p, q Point) float64 {
//	return math.Hypot(q.X - p.X, q.Y - p.Y)
//}

// 定义一个方法
func (p Point) Distance(q Point) float64 {
	// 勾股定理
	return math.Hypot(q.X - p.X, q.Y - p.Y)
}

type Path []Point

func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}

func main() {
	p := Point{1, 2}
	q := Point{4, 6}
	fmt.Println(p.Distance(q))
	fmt.Println(q.Distance(p))
	//fmt.Println(Distance(p, q))
	perim := Path{ {1, 1}, {5, 1}, {5, 4}, {1, 1}, }
	fmt.Println(perim.Distance())
}
