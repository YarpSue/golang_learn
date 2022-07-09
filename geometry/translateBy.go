package geometry

// 根据一个变量来决定调用同一个类型的哪个函数时

func (p Point) Add(q Point) Point {return Point{p.X + q.X, p.Y + q.Y}}
func (p Point) Sub(q Point) Point {return Point{p.X - q.X, p.Y - q.Y}}

func (path Path) translateBy(offset Point, add bool) {
	// 核心的部分  选择调用函数
	var op func(p, q Point) Point
	if add {
		op = Point.Add
	} else {
		op = Point.Sub
	}
	for i := range path {
		path[i] = op(path[i], offset)
	}
}