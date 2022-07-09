package geometry

import (
	"image/color"
	"math"
)

type Point struct{ X,Y float64}

//func Distance(p, q Point) float64 {
//	return math.Hypot(q.X - p.X, q.Y - p.Y)
//}
//嵌入结构体来拓展类型
type ColoredPoint struct {
	Point
	Color color.RGBA
}
// 嵌入指针的用法
type ColoredPoint2 struct {
	*Point
	Color color.RGBA
}

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

//func main() {
//	//p := Point{1, 2}
//	//q := Point{4, 6}
//	//fmt.Println(p.Distance(q))
//	//fmt.Println(q.Distance(p))
//	////fmt.Println(Distance(p, q))
//	//perim := Path{ {1, 1}, {5, 1}, {5, 4}, {1, 1}, }
//	//fmt.Println(perim.Distance())
//
//	// --------------------------------
//	// 类型嵌套变量是共享的
//	var cp ColoredPoint
//	cp.X = 1
//	fmt.Println(cp.Point.X)
//	cp.Point.Y = 2
//	fmt.Println(cp.Y)
//
//	// 如果想要调用Point类型下的方法呢?
//	red := color.RGBA{255, 0, 0 ,255}
//	blue := color.RGBA{0 , 0, 255, 255}
//	var p = ColoredPoint{Point{1, 1}, red}
//	var q = ColoredPoint{Point{5, 4}, blue}
//	fmt.Println(p.Distance(q.Point))
//
//	g :=ColoredPoint2{&Point{1, 1}, red}
//	h :=ColoredPoint2{&Point{5, 4}, red}
//	fmt.Println(g.Distance(*h.Point))
//}
